package main

import (
    "bitbucket.org/bertimus9/systemstat"
    "fmt"
    "time"
)

var last_cpu systemstat.CPUSample
var current_cpu systemstat.CPUSample
var cpu systemstat.CPUAverage

func main() {
    last_cpu = systemstat.GetCPUSample()
    time.Sleep(1000 * time.Millisecond) 
    current_cpu = systemstat.GetCPUSample()
    cpu = systemstat.GetCPUAverage(last_cpu, current_cpu) 
    fmt.Println(cpu)
}
