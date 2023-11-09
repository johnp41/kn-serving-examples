package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("On my way to unikernels!!!")

	// Get the hostname of the local machine
	hostname, err := os.Hostname()
	if err != nil {
		fmt.Printf("Error getting hostname: %v\n", err)
		os.Exit(1)
	}

	// Resolve the hostname to an IP address
	addrs, err := net.LookupHost(hostname)
	if err != nil {
		fmt.Printf("Error resolving hostname to IP: %v\n", err)
		os.Exit(1)
	}

	// Print the first IP address found (assuming it's the primary one)
	fmt.Printf("Node IP: %s\n", addrs[0])

	fmt.Fprintf(w, "Hello ! The  IP  is : "+addrs[0]+"\n")
}

func main() {
	log.Print("helloworld: starting server...")

	http.HandleFunc("/", handler)

	port := os.Getenv("C_PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("helloworld: listening on port %s", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}
