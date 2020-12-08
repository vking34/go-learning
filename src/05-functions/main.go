package main

import "fmt"

func sum(a, b int) int {
	var sum = a + b
	return sum
}

func getRectPros(length, width float64) (area, perimeter float64) {
	area = length * width
	perimeter = 2 * (length + width)
	return area, perimeter
}

func main() {
	x, y := 2, 5
	total := sum(x, y)
	fmt.Println("Sum:", total)

	area, _ := getRectPros(4.5, 2.5)
	fmt.Printf("Area %f\n", area)
}
