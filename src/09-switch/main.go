package main

import (
	"fmt"
)

func number() int {
	num := 5 * 5
	return num
}

func main() {
	switch finger := 2; finger {
	case 1:
		fmt.Println("Thumb")
	case 2:
		fmt.Println("Index")
	case 3:
		fmt.Println("Middle")
	case 4:
		fmt.Println("Ring")
	case 5:
		fmt.Println("Pinky")
	case 6, 7:
		fmt.Println("Strange DNA")
	default: //default case
		fmt.Println("incorrect finger number")
	}

	switch num := number(); { //num is not a constant
	case num < 50:
		fmt.Printf("%d is lesser than 50\n", num)
		fallthrough
	case num < 100:
		fmt.Printf("%d is lesser than 100\n", num)

	case num < 200:
		fmt.Printf("%d is lesser than 200\n", num)
	}

}
