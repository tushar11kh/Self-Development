#!/bin/bash

# Function to get CPU times
get_cpu_times() {
    awk '/cpu / {print $2, $3, $4, $5, $6, $7, $8}' /proc/stat
}

read -r user1 nice1 system1 idle1 iowait1 irq1 softirq1 < <(get_cpu_times)

total_usage1=$((user1+nice1+system1+idle1+iowait1+irq1+softirq1))
idle_time1=$((idle1+iowait1))

sleep 1

read -r user2 nice2 system2 idle2 iowait2 irq2 softirq2 < <(get_cpu_times)

total_usage2=$((user2+nice2+system2+idle2+iowait2+irq2+softirq2))
idle_time2=$((idle2+iowait2))

total_diff=$((total_usage2-total_usage1))
idle_diff=$((idle_time2-idle_time1))

cpu_usage=$((((total_diff-idle_diff)*100)/total_diff))

echo "CPU usage:$cpu_usage%"