package main

import (
	"bitbucket.org/bertimus9/systemstat"
	"fmt"
	"time"
	"flag"
	"os"
	"strconv"
)

var last_cpu systemstat.CPUSample
var current_cpu systemstat.CPUSample
var cpu_stat systemstat.SimpleCPUAverage
var cpu_average systemstat.CPUAverage

func main() {
	var w = flag.String("warning", "80", "Warning Level(%)")
	var c = flag.String("critical", "100", "Critical Level(%)")
	//var s = flag.Int("Sleep", 1000, "Set SleepTime")
	flag.Parse()

	warning, _ := strconv.ParseFloat(*w, 64)
	critical, _ := strconv.ParseFloat(*c, 64)
	//sleep, _ := strconv.ParseFloat(*s, 64)

	last_cpu = systemstat.GetCPUSample()
	//fmt.Println(last_cpu)
	time.Sleep(1000 * time.Millisecond) 
	current_cpu = systemstat.GetCPUSample()
	//fmt.Println(current_cpu)

	cpu_stat = systemstat.GetSimpleCPUAverage(last_cpu, current_cpu)
	cpu_average = systemstat.GetCPUAverage(last_cpu, current_cpu)

	fmt.Printf("CPUAverage=")
	fmt.Println(cpu_average)

	if cpu_stat.BusyPct > critical {
		fmt.Println("critical")
		os.Exit(2)
	} else if cpu_stat.BusyPct > warning {
		fmt.Println("warning")
		os.Exit(1)
	} else {
		fmt.Println("ok")
		os.Exit(0)
	}
}
