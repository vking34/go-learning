package main

import "fmt"

func sum(a, b int) int {
	var sum = a + b
	return sum
}

func rectProps(length, width float64) (float64, float64) {
	var area = length * width
	var perimeter = 2 * (length + width)
	return area, perimeter
}

func main() {
	x, y := 2, 5
	total := sum(x, y)
	fmt.Println("Sum:", total)

	area, perimeter := rectProps(4.5, 2.5)
	fmt.Println("Area:", area, ", Perimeter:", perimeter)
	fmt.Printf("Area: %f, Perimeter: %f\n", area, perimeter)
}
