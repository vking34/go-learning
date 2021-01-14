package main

import "fmt"

type VowelsFinder interface {
	FindVowels() []rune
}

type MyString string

func (str MyString) FindVowels() []rune {
	var vowels []rune

	for _, rune := range str {
		if rune == 'a' || rune == 'e' || rune == 'o' || rune == 'i' || rune == 'u' {
			vowels = append(vowels, rune)
		}
	}

	return vowels
}

func printVowels(vowels []rune) {
	fmt.Print("Vowels: ")
	for _, rune := range vowels {
		fmt.Printf("%s ", string(rune))
	}
}

func main() {
	name := MyString("Vuong")
	var v VowelsFinder
	v = name
	vowels := v.FindVowels()
	fmt.Printf("Vowels: %c\n", vowels)
	printVowels(vowels)
}
