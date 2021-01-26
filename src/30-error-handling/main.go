package main

import (
	"fmt"
	"os"
)

func main() {
	// error
	f, err := os.Open("./test.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(f.Name(), "opened successfully!")
}
