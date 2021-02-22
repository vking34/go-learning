package main

import (
	"fmt"
	"runtime/debug"
)

func recovery() {
	if r := recover(); r != nil {
		fmt.Println("recovered:", r)
	}
}

func sum(a int, b int) {
	defer recovery() // can not recover the panic from the divide goroutine
	fmt.Printf("%d + %d = %d\n", a, b, a+b)
	done := make(chan bool)
	go divide(a, b, done)
	<-done // wait the divide goroutine
}

func divide(a int, b int, done chan bool) {
	fmt.Printf("%d / %d = %d", a, b, a/b) // panic: divide by zero
	done <- true
}

func recoverFullName() {
	if r := recover(); r != nil {
		fmt.Println("recovered from", r)
		debug.PrintStack()
	}
}

func fullName(firstName *string, lastName *string) {
	defer recoverFullName()
	if firstName == nil {
		panic("runtime error: first name cannot be nil")
	}
	if lastName == nil {
		panic("runtime error: last name cannot be nil")
	}
	fmt.Printf("%s %s\n", *firstName, *lastName)
	fmt.Println("returned normally from fullName")
}

func main() {
	defer fmt.Println("deffered call in main")
	firstName := "Vuong"
	fullName(&firstName, nil)
	fmt.Println("returned normally from main")

	// sum(5, 0)
	// fmt.Println("normally returned from main")
}
