package main

import (
	"fmt"
	"unicode/utf8"
)

func printBytes(s string) {
	fmt.Print("Bytes: ")
	for i := 0; i < len(s); i++ {
		fmt.Printf("%x ", s[i])
	}
	fmt.Print("\n")
}

func printChars(s string) {
	fmt.Print("Chars: ")
	runes := []rune(s)
	for i := 0; i < len(runes); i++ {
		fmt.Printf("%c ", runes[i])
	}
}

func printCharAndPosition(s string) {
	for i, rune := range s {
		fmt.Printf("%c starts at byte %d\n", rune, i)
	}
}

func compareStrings(s1 string, s2 string) bool {
	if s1 == s2 {
		fmt.Printf("%s and %s are equal\n", s1, s2)
		return true
	}
	fmt.Printf("%s and %s are not equal\n", s1, s2)
	return false
}

func main() {
	s := "hello world"
	fmt.Println("String:", s)
	printBytes(s)
	fmt.Printf("\n")

	//
	unicodeStr := "Señor"
	fmt.Println("String:", unicodeStr)
	fmt.Println("Length:", len(unicodeStr))
	printChars(unicodeStr)
	fmt.Printf("\n")
	printBytes(unicodeStr)
	printCharAndPosition(unicodeStr)

	// convert a slice of bytes to string
	byteSlice := []byte{0x43, 0x61, 0x66, 0xC3, 0xA9}
	str := string(byteSlice)
	fmt.Println("string:", str)
	byteSlice1 := []byte{67, 97, 102, 195, 169}
	str = string(byteSlice1)
	fmt.Println("string:", str)

	// convert a slice of runes to string
	runeSlice := []rune{0x0053, 0x0065, 0x00f1, 0x006f, 0x0072}
	str = string(runeSlice)
	fmt.Println("string:", str)

	// length of string
	fmt.Println("String:", str)
	fmt.Println("Length:", utf8.RuneCountInString(str))
	fmt.Println("Number of bytes:", len(str))

	// compare
	str1 := "hello"
	str2 := "world"
	str3 := "hello"

	compareStrings(str1, str2)
	compareStrings(str1, str3)

	// concat
	str4 := str1 + " " + str2
	fmt.Println("string:", str4)

	// immutable
	strRune := []rune(str)
	strRune[0] = 'D'
	newStr := string(strRune)
	fmt.Println("New string:", newStr)
}
