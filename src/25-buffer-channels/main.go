package main

import "fmt"

func sendData(ch chan string) {
	ch <- "Hello" // write 2 strings to channel
	ch <- "World" // without blocking
}

func main() {

	// declaration and usage
	ch := make(chan string, 2)
	go sendData(ch)
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}
