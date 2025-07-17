package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {

	cpuFlag := flag.Bool("cpu", false, "print cpu usage")
	memFlag := flag.Bool("mem", false, "print memory usage")
	allFlag := flag.Bool("all", false, "print hardware usage")

	flag.Parse()

	if *cpuFlag {
		showCpu()
	} else if *memFlag {
		showMem()
	} else if *allFlag {
		showCpu()
		showMem()
	} else {
		fmt.Println("use --help")
	}

}

func showCpu() {
	cpu_list1, err := cpuInfo()
	if err != nil {
		fmt.Println("Cpu info cannot be found")
		return
	}

	total1 := 0
	for _, num := range cpu_list1 {
		total1 += num
	}

	idle1 := cpu_list1[3] + cpu_list1[4]

	time.Sleep(1 * time.Second)

	cpu_list2, err := cpuInfo()
	if err != nil {
		fmt.Println("Cpu info cannot be found")
		return
	}

	total2 := 0
	for _, num := range cpu_list2 {
		total2 += num
	}

	idle2 := cpu_list2[3] + cpu_list2[4]

	total_diff := total2 - total1

	idle_diff := idle2 - idle1

	cpu_usage := float64(total_diff-idle_diff) / float64(total_diff) * 100

	fmt.Printf("Cpu Usage:%.2f%%\n", cpu_usage)

}

func cpuInfo() ([]int, error) {
	file, err := os.Open("/proc/stat")
	if err != nil {
		return make([]int, 0), err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var requireString []string

	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "cpu ") {
			requireString = strings.Fields(line)
			break
		}
	}

	var i int
	requireInt := make([]int, len(requireString)-4)
	for i = 1; i < len(requireString)-3; i++ {
		requireInt[i-1], err = strconv.Atoi(requireString[i])
	}

	return requireInt, err

}

func showMem() {
	memList, err := memInfo()
	if err != nil {
		fmt.Println("Memory info cannot be found")
		return
	}

	totalMem := memList[0]
	availMem := memList[1]

	mem_usage := (float64(totalMem-availMem) * 100) / float64(totalMem)

	fmt.Printf("Memory Usage:%.2f%%\n", mem_usage)

}

func memInfo() ([]int, error) {
	file, err := os.Open("/proc/meminfo")
	if err != nil {
		return make([]int, 0), err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	requiredNum := make([]int, 2)

	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "MemTotal:") {
			requiredNum[0], err = strconv.Atoi(strings.Fields(line)[1])
		}
		if strings.HasPrefix(line, "MemAvailable:") {
			requiredNum[1], err = strconv.Atoi(strings.Fields(line)[1])
		}

		if requiredNum[0] != 0 && requiredNum[1] != 0 {
			break
		}
	}

	return requiredNum, err

}
