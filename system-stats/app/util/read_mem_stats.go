package util

import (
	"bufio"
	"os"

	"golang.org/x/exp/slices"
)

var STAT_KEYS = []string{"MemTotal", "MemFree", "MemAvailable", "SwapTotal", "SwapCached", "SwapFree"}

func ReadMemoryStats() (map[string]int, error) {
	// /proc/meminfo contains memory info. Read the file and parse the fields we need
	file, err := os.Open("/proc/meminfo")
	if err != nil {
		return nil, err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	stats := map[string]int{}
	for scanner.Scan() {
		key, value := ParseMemInfoLine(scanner.Text())
		if slices.Contains(STAT_KEYS, key) {
			stats[key] = value
		}
	}

	// http.sendStatsToProducer(stats)
	return stats, nil
}
