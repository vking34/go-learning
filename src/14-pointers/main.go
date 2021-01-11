package main

import "fmt"

func main() {

	// declaration
	b := 255
	var a *int = &b
	fmt.Printf("Type of a: %T\n", a)
	fmt.Println("Address of b:", a)
	fmt.Println("Value that pointer a point to:", *a)

	// new function
	size := new(int)
	fmt.Printf("Size value is %d, type is %T, address is %v\n", *size, size, size)
	*size = 100
	fmt.Println("New value:", *size)

	// dereferencing
	*a++
	fmt.Println("New value of b:", b)
}
