package main

import (
	"fmt"
	"time"
)

func server1(ch chan string) {
	time.Sleep(2 * time.Second)
	ch <- "data from server1"
}

func server2(ch chan string) {
	time.Sleep(1 * time.Second)
	ch <- "data from server2"
}

func process(ch chan string) {
	time.Sleep(10500 * time.Millisecond)
	ch <- "Process completed!"
}

func main() {
	// usage
	output1 := make(chan string)
	output2 := make(chan string)
	go server1(output1)
	go server2(output2)
	select {
	case s1 := <-output1:
		fmt.Println(s1)
	case s2 := <-output2:
		fmt.Println(s2)
	}

	// default case
	ch := make(chan string)
	go process(ch)

check:
	for {
		time.Sleep(time.Second)
		select {
		case s := <-ch:
			fmt.Println("Received data:", s)
			break check
		default:
			fmt.Println("No received data yet.")
		}
	}

	// default case prevents deadlock
	ch1 := make(chan string)
	select {
	case <-ch1:
	default: // deadlock if without default case
		fmt.Println("default case executed.")
	}

	// random selection
	go server1(output1)
	go server2(output2)
	time.Sleep(3 * time.Second)
	select {
	case s1 := <-output1:
		fmt.Println(s1)
	case s2 := <-output2:
		fmt.Println(s2)
	}
}
