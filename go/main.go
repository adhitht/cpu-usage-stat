package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func getCPUTimes(procPath string) (int, int, error) {
	file, err := os.Open(procPath + "/stat")
	if err != nil {
		return 0, 0, err
	}
	defer file.Close()

	stats := make([]int, 10)

	scanner := bufio.NewScanner(file)
	if scanner.Scan() {
		line := scanner.Text()
		fields := strings.Fields(line)

		for i := 1; i < len(fields); i++ {
			num, err := strconv.Atoi(fields[i])
			if err != nil {
				return 0, 0, fmt.Errorf("failed to parse field %d (%s): %w", i, fields[i], err)
			}
			stats[i-1] = num
		}
	}

	if err := scanner.Err(); err != nil {
		return 0, 0, err
	}
	
	idleTime := stats[3]
	totalTime := 0 

	for _, value := range stats {
		totalTime += value
	}
	
	return idleTime, totalTime, nil
}

func getCPUUsage(procPath string) (float64, error) {
	idleTime1, totalTime1, err := getCPUTimes(procPath)
	if err != nil {
		return 0, err
	}
	
	time.Sleep(time.Second)
	
	idleTime2, totalTime2, err := getCPUTimes(procPath)
	if err != nil {
		return 0, err
	}
	
	deltaIdleTime := idleTime2 - idleTime1
	deltaTotalTime := totalTime2 - totalTime1
	
	return (1 - float64(deltaIdleTime) / float64(deltaTotalTime)) * 100, nil
}

func main() {
	procPath := "/proc"
	cpuUsage, err := getCPUUsage(procPath)

	if(err != nil){
		fmt.Println(err)
		return
	}
	
	fmt.Println("CPU Usage: ", cpuUsage)
}
