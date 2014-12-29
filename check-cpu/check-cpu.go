package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
	"time"
	"strconv"
)

func main() {
	r_before := strings.Split(acquire_cpu_stats("/proc/stat"), " ")
	time.Sleep(1000 * time.Millisecond)
	r_after := strings.Split(acquire_cpu_stats("/proc/stat"), " ")
        // cpu user   nice system idle      iowait irq softirq steal guest ?
	// cpu 787111 0    740915 132510750 6695   138 108643  0     0     0

	fmt.Println(r_before)
	fmt.Println(r_after)

	var (
		us_before, us_after, ni_before, ni_after, sy_before, sy_after, id_before, id_after float64
		us_sub, ni_sub, sy_sub, id_sub, total float64
		us, ni, sy, id float64
	)
	us_before, _ = strconv.ParseFloat(r_before[2], 64)
	us_after, _ = strconv.ParseFloat(r_after[2], 64)
	ni_before, _ = strconv.ParseFloat(r_before[3], 64)
	ni_after, _ = strconv.ParseFloat(r_after[3], 64)
	sy_before, _ = strconv.ParseFloat(r_before[4], 64)
	sy_after, _ = strconv.ParseFloat(r_after[4], 64)
	id_before, _ = strconv.ParseFloat(r_before[5], 64)
	id_after, _ = strconv.ParseFloat(r_after[5], 64)

	us_sub = us_after - us_before 
	ni_sub = ni_after - ni_before 
	sy_sub = sy_after - sy_before 
	id_sub = id_after - id_before 

	total = us_sub + ni_sub + sy_sub + id_sub

	us = (us_after - us_before) / total * 100.0
	ni = (ni_after - ni_before) / total * 100.0
	sy = (sy_after - sy_before) / total * 100.0
	id = (id_after - id_before) / total * 100.0

	fmt.Printf("user = %3.2f%% ", us)
	fmt.Printf("nice = %3.2f%% ", ni)
	fmt.Printf("system = %3.2f%% ", sy)
	fmt.Printf("idle = %3.2f%%\n", id)
}

func acquire_cpu_stats(p string)(s string) {
	cpu_stats, err := os.OpenFile(p, os.O_RDONLY, 0)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	defer func() {
		cpu_stats.Close()
	}()


	scanner := bufio.NewScanner(cpu_stats)
	for scanner.Scan() {
		matched, _ := regexp.MatchString("^cpu ", scanner.Text())
		if matched {
			return scanner.Text()
		}	
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
	return s
}
