package main

import (
	"06-packages/simpleInterest"
	"fmt"
)

func main() {
	fmt.Println("Sinple interest calucation")
	p := 5000.0
	r := 10.0
	t := 1.0
	si := simpleInterest.Calculate(p, r, t)
	fmt.Println("Simple interest is", si)
}
