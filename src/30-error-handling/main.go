package main

import (
	"fmt"
	"net"
	"os"
	"path/filepath"
)

func main() {
	f, err := os.Open("./test.txt")

	// check error
	if err != nil {
		// extract error using type assertion
		if pErr, ok := err.(*os.PathError); ok {
			fmt.Println("Failed to open file at path", pErr.Path)
		} else {
			fmt.Println("Generic erorr:", err)
		}
	} else {
		fmt.Println(f.Name(), "opened successfully")
	}

	// extract error using methods
	addr, err := net.LookupHost("golangbot123.com")
	if err != nil {
		if dnsErr, ok := err.(*net.DNSError); ok {
			if dnsErr.Timeout() {
				fmt.Println("operation timed out")
			}
			if dnsErr.Temporary() {
				fmt.Println("Temporary error")
			} else {
				fmt.Println("Generic DNS error:", err)
			}
		} else {
			fmt.Println("Generic error:", err)
		}
	} else {
		fmt.Println("Address:", addr)
	}

	files, err := filepath.Glob("[")
	if err != nil {
		if err == filepath.ErrBadPattern {
			fmt.Println("Bad pattern error:", err)
		} else {
			fmt.Println("Generic error:", err)
		}
	} else {
		fmt.Println("matched files:", files)
	}
}
