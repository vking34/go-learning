package main

import "fmt"

func printBytes(s string) {
	fmt.Println("Bytes: ")
	for i := 0; i < len(s); i++ {
		fmt.Printf("%x ", s[i])
	}
}

func main() {
	s := "hello world"
	fmt.Println("String:", s)
	printBytes(s)
}
