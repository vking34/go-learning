package main

import (
	"fmt"
	"time"
)

func sendData(ch chan string) {
	ch <- "Hello"
	ch <- "World"
}

func write(ch chan int) {
	for i := 0; i < 5; i++ {
		ch <- i
		fmt.Println("wrote", i, "to the channel")
	}
	close(ch)
}

func main() {
	// declaration and usage
	ch := make(chan string, 2)
	go sendData(ch)
	fmt.Println(<-ch)
	fmt.Println(<-ch)

	// understand how it works
	ch1 := make(chan int, 2)
	go write(ch1)
	time.Sleep(2 * time.Second)
	for v := range ch1 {
		fmt.Println("read value", v, "from the channel")
		time.Sleep(2 * time.Second)
	}

	// deadlock
	ch2 := make(chan string, 2)
	ch2 <- "Le"
	ch2 <- "Van"
	// deadlock because of no goroutine consume the channel
	// ch2 <- "Vuong"

	// length vs capacity
	fmt.Println("capacity of ch2:", cap(ch2))
	fmt.Println("length of ch2:", len(ch2))
	fmt.Println("read from ch2:", <-ch2)
	fmt.Println("new length of ch2:", len(ch2))

}
