package main

import "fmt"

func main() {
	var age int // init value
	fmt.Println("My age is", age)

	age = 23 // assignment
	fmt.Println("My age is", age)

	var newAge = 24 // type inference
	fmt.Println("My new age is", newAge)

	var (
		name   = "vuong"
		old    = 23
		height int
	)
	fmt.Println("my name is", name, ", age is", old, "and height is", height)

	// short hand declaration
	count := 10
	fmt.Println("Count =", count)

	title, pages := "Psycho", 23
	fmt.Println("Book:", title, "pages:", pages)

	// Short hand syntax can only be used when at least one of the variables
	// on the left side of := is newly declared.
	a, b := 20, 30 // declare variables a and b
	fmt.Println("a is", a, "b is", b)
	b, c := 40, 50 // b is already declared but c is new
	fmt.Println("b is", b, "c is", c)
	b, c = 80, 90 // assign new values to already declared variables b and c
	fmt.Println("changed b is", b, "c is", c)

	
}
