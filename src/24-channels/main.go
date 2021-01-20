package main

import (
	"fmt"
	"time"
)

func hello(done chan bool) {
	fmt.Println("Hello Goroutine is going to sleep")
	time.Sleep(2 * time.Second)
	fmt.Println("Hello Goroutine awake to write data to the channel")
	done <- true
}

func sumSquares(num int, sqrChannel chan int) {
	sum := 0
	for num != 0 {
		digit := num % 10
		sum += digit * digit
		num = (num - digit) / 10
	}
	sqrChannel <- sum
}

func sumCubes(num int, cubeChannel chan int) {
	sum := 0
	for num != 0 {
		digit := num % 10
		sum += digit * digit * digit
		num = (num - digit) / 10
	}
	cubeChannel <- sum
}

func sendData(unidirectionalChannel chan<- int) {
	unidirectionalChannel <- 10
}

func producer(ch chan int) {
	for i := 0; i < 5; i++ {
		ch <- i
	}
	close(ch)
}

func producer1(ch chan int) {
	for i := 0; i < 10; i++ {
		ch <- i
	}
	close(ch)
}

func main() {

	// declaration
	var a chan int
	if a == nil {
		fmt.Println("Channel is nil, going to initialize it")
		a = make(chan int)
		fmt.Printf("Type of a is %T\n", a)
	}

	// example using channel
	done := make(chan bool)
	fmt.Println("Main going to call hello goroutine")
	go hello(done)

	// the control blocks in receive statement
	// until some Goroutine writes data to the channel
	<-done

	fmt.Println("Main received data")

	// another example: A program will print the sum of the squares and cubes of the individual digits of a number.
	// the squares are calculated in a separate Goroutine, cubes in another Goroutine and the final summation happens in the main Goroutine.
	num := 235
	sqrChannel := make(chan int)
	cubeChannel := make(chan int)
	go sumCubes(num, cubeChannel)
	go sumSquares(num, sqrChannel)
	sqrSum, cubeSum := <-cubeChannel, <-sqrChannel
	sum := sqrSum + cubeSum
	fmt.Println("Sum:", sum)

	// deadlock
	// deadChannel := make(chan string)
	// deadChannel <- "DEAD!"

	// unidirectional channel
	bidirectionalChannel := make(chan int)
	go sendData(bidirectionalChannel)
	fmt.Println("Data from unidirectional channel:", <-bidirectionalChannel)

	// close channel
	ch := make(chan int)
	go producer(ch)
	for {
		v, isOpen := <-ch
		if isOpen == false {
			fmt.Println("Received: ", v)
			break
		}
		fmt.Println("Received: ", v)
	}
	fmt.Println("-------------")

	// for range channel
	ch1 := make(chan int)
	go producer1(ch1)
	for v := range ch1 {
		fmt.Println("Received: ", v)
	}
}
