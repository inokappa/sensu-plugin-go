// hoge project main.go
package main

import (
	"flag"
	"fmt"
	"net"
	"os"
)

func main() {
	//(name, default, help)
	var proto = flag.String("prot", "tcp", "Set Protocol")
	var hostname = flag.String("hostname", "www.example.com", "Set Hostname")
	var port = flag.String("port", "80", "Set Port Number")
	flag.Parse()

	fmt.Print("Checking..." + *hostname + ":" + *port + "\n")
	conn, err := net.Dial(*proto, *hostname+":"+*port)

	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	} else {
		fmt.Print("OK" + "\n")
		os.Exit(0)
	}

	defer func() {
		conn.Close()
	}()
}
