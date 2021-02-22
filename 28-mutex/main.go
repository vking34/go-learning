package main

import (
	"fmt"
	"sync"
)

var x int = 0
var y int = 0

func add(w *sync.WaitGroup, m *sync.Mutex) {
	m.Lock()
	x = x + 1
	m.Unlock()
	w.Done()
}

func increase(w *sync.WaitGroup, ch chan bool) {
	ch <- true
	y = y + 1
	<-ch
	w.Done()
}

func main() {
	// mutex
	var w sync.WaitGroup
	var m sync.Mutex
	for i := 0; i < 1000; i++ {
		w.Add(1)
		go add(&w, &m)
	}

	w.Wait()
	fmt.Println("x:", x)

	// solving the race condition using channel
	ch := make(chan bool, 1)
	for i := 0; i < 1000; i++ {
		w.Add(1)
		go increase(&w, ch)
	}

	w.Wait()
	fmt.Println("y:", y)
}
