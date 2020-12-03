package main

import "fmt"

func sum(a, b int) int {
	var sum = a + b
	return sum
}

func main() {
	x, y := 2, 5
	total := sum(x, y)
	fmt.Println("Sum:", total)
}
