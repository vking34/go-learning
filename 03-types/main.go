package main

import (
	"fmt"
	"unsafe"
)

func main() {
	// int
	var a int = 89
	b := 95
	fmt.Println("value of a is", a, "and b is", b)
	fmt.Printf("type of a is %T, size of a is %d", a, unsafe.Sizeof(a))
	fmt.Printf("\ntype of b is %T, size of b is %d\n", b, unsafe.Sizeof(b))

	// float
	x, y := 5.67, 8.97
	fmt.Printf("type of x %T y %T\n", x, y)
	sum := x + y
	diff := x - y
	fmt.Println("sum", sum, "diff", diff)

	no1, no2 := 56, 89
	fmt.Println("sum", no1+no2, "diff", no1-no2)

	// string
	firstName := "Vuong"
	lastName := "Le"
	fullName := firstName + " " + lastName
	fmt.Println("My name is", fullName)

	// type conversion

	i := 55   // int
	j := 67.8 // float64
	// ! error: sum := i + j
	k := i + int(j) // take only integer part
	fmt.Println("Sum:", k)
}
