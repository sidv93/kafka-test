package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"

	"systemstats/app/config"

	"github.com/jasonlvhit/gocron"
	"golang.org/x/exp/slices"
)

type Memory struct {
	MemTotal     int
	MemFree      int
	MemAvailable int
	SwapTotal    int
	SwapCached   int
	SwapFree     int
}

var appConfig config.AppConfig

var STAT_KEYS = []string{"MemTotal", "MemFree", "MemAvailable", "SwapTotal", "SwapCached", "SwapFree"}

func ReadMemoryStats() {
	fmt.Println("in read memory stats")
	file, err := os.Open("/proc/meminfo")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	stats := map[string]int{}
	for scanner.Scan() {
		key, value := parseLine(scanner.Text())
		if slices.Contains(STAT_KEYS, key) {
			stats[key] = value
		}
	}
	fmt.Println("stats", stats)
	sendStatsToProducer(stats)
	// return res
}

func parseLine(raw string) (key string, value int) {
	text := strings.ReplaceAll(raw[:len(raw)-2], " ", "")
	keyValue := strings.Split(text, ":")
	return keyValue[0], toInt(keyValue[1])
}

func toInt(raw string) int {
	if raw == "" {
		return 0
	}
	res, err := strconv.Atoi(raw)
	if err != nil {
		panic(err)
	}
	return res
}

func runCronJobs() {
	s := gocron.NewScheduler()
	s.Every(10).Second().Do(ReadMemoryStats)
	<-s.Start()
}

func sendStatsToProducer(stats map[string]int) (map[string]int, error) {
	producerUrl := appConfig.ProducerUrl
	fmt.Println("in post", producerUrl, stats)

	requestBodyBytes, err := json.Marshal(stats)
	if err != nil {
		return nil, err
	}
	requestBody := bytes.NewReader(requestBodyBytes)
	request, err := http.NewRequest("POST", producerUrl, requestBody)
	fmt.Println(requestBody)
	if err != nil {
		fmt.Println("post error", err)
		return nil, err
	}

	request.Header.Set("Content-Type", "application/json")
	fmt.Println("making call")
	resp, err := http.DefaultClient.Do(request)
	if err != nil {
		fmt.Println("Post error", err)
		return nil, err
	}

	respBody, _ := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	var respData map[string]int
	err = json.Unmarshal(respBody, &respData)

	if resp.StatusCode != 200 {
		fmt.Println("post failed")
		return nil, errors.New("Unable to get flight path from planner")
	}
	if err != nil {
		return nil, err
	}
	return respData, nil
}

func main() {
	fmt.Println("System stats service started")
	appConfig = *config.NewAppConfig()
	fmt.Println("producer url", appConfig.ProducerUrl)
	runCronJobs()
	fmt.Scanln()
}
