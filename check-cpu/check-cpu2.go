package main

import (
    "bitbucket.org/bertimus9/systemstat"
    "fmt"
    "time"
)

var mem systemstat.MemSample
var last_cpu systemstat.ProcCPUSample
var current_cpu systemstat.ProcCPUSample
var cpu systemstat.ProcCPUAverage
var startTime time.Time
var lc systemstat.CPUSample
var cc systemstat.CPUSample
var c systemstat.CPUAverage

// This example shows how easy it is to get memory information
func main() {
    mem = systemstat.GetMemSample()
    fmt.Println("Total available RAM in kb:", mem.MemTotal, "k total")
    fmt.Println("Used RAM in kb:", mem.MemUsed, "k used")
    fmt.Println("Free RAM in kb:", mem.MemFree, "k free")
    //fmt.Printf("The output is similar to, but somewhat different than:\n\ttop -n1 | grep Mem:\n")
    last_cpu = systemstat.GetProcCPUSample()
    fmt.Println(last_cpu)
    //time.Sleep(1000 * time.Millisecond) 
    current_cpu = systemstat.GetProcCPUSample()
    fmt.Println(current_cpu)
    cpu = systemstat.GetProcCPUAverage(last_cpu, current_cpu, time.Since(startTime).Seconds())
    fmt.Printf("%8.5f \n", cpu.UserPct)
    fmt.Printf("%8.5f \n", cpu.Seconds)
    fmt.Printf("%8.5f \n", cpu.TotalPct)
    fmt.Printf("%8.5f \n", cpu.CumulativeTotalPct)
    fmt.Printf("%8.5f \n", cpu.SystemPct)
    lc = systemstat.GetCPUSample()
    time.Sleep(1000 * time.Millisecond) 
    cc = systemstat.GetCPUSample()
    fmt.Println(lc)
    fmt.Println(cc)
    c = systemstat.GetCPUAverage(lc, cc) 
    fmt.Println(c)
}
