package main

import (
	"fmt"
	"math/rand"
	"os"
	"sync"
)

func produce(data chan int, wg *sync.WaitGroup) {
	n := rand.Intn(999)
	data <- n
	wg.Done()
}

func consume(data chan int, done chan bool) {
	f, err := os.Create("concurrent")
	if err != nil {
		fmt.Println(err)
		return
	}
	for d := range data {
		_, err = fmt.Fprintln(f, d)
		if err != nil {
			fmt.Println(err)
			f.Close()
			done <- false
			return
		}
	}
	err = f.Close()
	if err != nil {
		fmt.Println(err)
		done <- false
		return
	}
	done <- true
}

func main() {
	// create file
	f1, err := os.Create("test.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	// write string to file
	numOfWrittenBytes, err := f1.WriteString("Hello World\n")
	if err != nil {
		fmt.Println(err)
		f1.Close()
		return
	}

	fmt.Println(numOfWrittenBytes, "bytes written successfully!")
	err = f1.Close()

	if err != nil {
		fmt.Println(err)
		return
	}

	// append to a file
	f2, err := os.OpenFile("test.txt", os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println(err)
		return
	}

	newLine := "File handling is easy."
	_, err = fmt.Fprintln(f2, newLine)
	if err != nil {
		fmt.Println(err)
		f2.Close()
		return
	}
	err = f2.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("file appended successfully")

	// write to file concurrently
	data := make(chan int)
	done := make(chan bool)
	wg := sync.WaitGroup{}
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go produce(data, &wg)
	}
	go consume(data, done)
	go func() {
		wg.Wait()
		close(data)
	}()
	d := <-done
	if d == true {
		fmt.Println("File written successfully")
	} else {
		fmt.Println("File writing failed")
	}
}
