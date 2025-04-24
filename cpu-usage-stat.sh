#!/bin/bash

get_cpu_times() {
  # Read the first line of /proc/stat
  read cpu user nice system idle iowait irq softirq steal guest guest_nice < /proc/stat
  total=$((user + nice + system + idle + iowait + irq + softirq + steal))
  echo "$idle $total"
}

get_cpu_usage() {
  read idle1 total1 < <(get_cpu_times)
  sleep 1
  read idle2 total2 < <(get_cpu_times)

  delta_idle=$((idle2 - idle1))
  delta_total=$((total2 - total1))

  usage=$(awk "BEGIN {printf \"%.2f\", (1 - $delta_idle / $delta_total) * 100}")
  echo "$usage"
}

usage=$(get_cpu_usage)
echo "CPU Usage: $usage%"
