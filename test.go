package main

import (
	"fmt"
	"net"
	"time"
)

func isPortOpen(host string, port int, timeout time.Duration) bool {
	target := fmt.Sprintf("%s:%d", host, port)
	conn, err := net.DialTimeout("tcp", target, timeout)
	if err != nil {
		return false
	}
	conn.Close()
	return true
}

func main() {
	// Test if port is open
	if isPortOpen("localhost", 5252, 2*time.Second) {
		fmt.Println("Port 5252 is open and listening")
	} else {
		fmt.Println("Port 5252 is not responding")
	}
}
