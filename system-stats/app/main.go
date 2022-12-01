package main

import (
	"fmt"

	"systemstats/app/http"
	"systemstats/app/util"

	"github.com/jasonlvhit/gocron"
)

func main() {
	fmt.Println("System stats service started")
	runCronJob()
	fmt.Scanln()
}

func runCronJob() {
	// cron job which runs every second
	s := gocron.NewScheduler()
	s.Every(5).Second().Do(collectStatsAndPushToProducer)
	<-s.Start()
}

func collectStatsAndPushToProducer() {
	// get system stats
	stats, err := util.ReadMemoryStats()
	fmt.Println("stats", stats)
	if err != nil {
		fmt.Println("Error when reading system stats", err)
		return
	}

	// push stats to producer
	_, err = http.SendStatsToProducer(stats)
	if err != nil {
		fmt.Println("Error when sending stats to producer", err)
	}

	fmt.Println("Sent system stats to producer")
}
