package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"strings"
	"time"
)

var panicBuffer = 0

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if panicBuffer == 5 {
			os.Exit(1)
		}
		panicBuffer++

		fmt.Fprintf(w, "Hello go! It's %s. Your host has IPs %s.", time.Now().Format("01-02-2006 15:04:05"), getLocalAddresses())
	})

	fmt.Printf("Starting server at port 4522\n")
	if err := http.ListenAndServe(":4522", nil); err != nil {
		log.Fatal(err)
	}
}

// Local address of the running system
func getLocalAddresses() string {
	// Get the IP address of the server
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		panic(err)
	}

	var ipAddresses []string

	for _, addr := range addrs {
		if ipnet, ok := addr.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				ipAddresses = append(ipAddresses, ipnet.IP.String())
			}
		}
	}

	return strings.Join(ipAddresses, ", ")
}
