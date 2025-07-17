package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func getMemData() (total, available int, err error) {
	file, err := os.Open("/proc/meminfo")
	if err != nil {
		return
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "MemTotal:") {
			fields := strings.Fields(line)
			total, _ = strconv.Atoi(fields[1])
		} else if strings.HasPrefix(line, "MemAvailable") {
			fields := strings.Fields(line)
			available, _ = strconv.Atoi(fields[1])
		}

		if total != 0 && available != 0 {
			break
		}

	}

	return

}

func main() {

	logfile, err := os.Create("memlogfile.txt")

	if err != nil {
		fmt.Println("Not able to create file")
		return
	}

	defer logfile.Close()

	writer := bufio.NewWriter(logfile)

	for {
		total, available, err := getMemData()

		if err != nil {
			fmt.Println("Error in getting Memory data")
			break
		}

		memUsage := total - available

		memUsagePercentage := (100 * float64(memUsage)) / float64(total)

		currentTime := time.Now().Format("02-01-2006 15:04:05")
		logline := fmt.Sprintf("%s\nMemory used:%.2f%%\n\n", currentTime, memUsagePercentage)
		fmt.Println(logline)
		writer.WriteString(logline)
		writer.Flush()
		time.Sleep(1 * time.Second)
	}

}
