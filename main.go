package main

import (
	"flag"
	"fmt"
	"net"
	"time"
)

func main() {
	// Define a command-line flag for the target IP address
	ipPtr := flag.String("i", "", "IP address to scan for open ports")
	flag.Parse()

	if *ipPtr == "" {
		fmt.Println("Please provide an IP address to scan using the -i flag.")
		return
	}

	// Scan ports from 1 to 1024 on the provided IP address
	for port := 1; port <= 1024; port++ {
		address := fmt.Sprintf("%s:%d", *ipPtr, port)

		// Attempt to connect to the port
		conn, err := net.DialTimeout("tcp", address, 500*time.Millisecond)
		if err != nil {
			// Port is closed or unreachable
			continue
		}
		defer conn.Close()

		// If we get here, the port is open
		fmt.Printf("Port %d: Open\n", port)
	}
}
