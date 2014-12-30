package main

import (
	"bitbucket.org/bertimus9/systemstat"
	"fmt"
	"flag"
	"os"
	"strconv"
)

var mem systemstat.MemSample

func main() {
	var w = flag.String("warning", "10", "Warning Level(%)")
	var c = flag.String("critical", "5", "Critical Level(%)")
	flag.Parse()

	warning, _ := strconv.ParseFloat(*w, 64)
	critical, _ := strconv.ParseFloat(*c, 64)

	mem = systemstat.GetMemSample()

	// Buffers Cached MemTotal MemUsed MemFree SwapTotal SwapUsed SwapFree time.Time
        //fmt.Println("Buffers Cached MemTotal MemUsed MemFree SwapTotal SwapUsed SwapFree time.Time")
	//fmt.Println(mem)

	mem_free_uint := (mem.Buffers + mem.Cached + mem.MemFree)
	mem_free, _ := strconv.ParseFloat(strconv.FormatUint(mem_free_uint, 10), 64)
	fmt.Printf("Free Memory: ")
	fmt.Printf("%3.2f", mem_free)
	fmt.Println(" KB")
	
	mem_total, _ := strconv.ParseFloat(strconv.FormatUint(mem.MemTotal, 10), 64)
	fmt.Printf("Total Memory: ")
	fmt.Printf("%3.2f", mem_total)
	fmt.Println(" KB")

	mem_free_pct := mem_free / mem_total * 100.00
	fmt.Printf("Free Percent: ")
	fmt.Printf("%3.2f", mem_free_pct)
	fmt.Println(" %")

	mem_used_pct := 100.00 - mem_free_pct
	fmt.Printf("Used Percent: ")
	fmt.Printf("%3.2f", mem_used_pct)
	fmt.Println(" %")


	if mem_free_pct < critical {
		fmt.Println("critical")
		os.Exit(2)
	} else if mem_free_pct < warning {
		fmt.Println("warning")
		os.Exit(1)
	} else {
		fmt.Println("ok")
		os.Exit(0)
	}
}
