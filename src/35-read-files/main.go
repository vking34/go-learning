package main

import (
	"bufio"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	// parse flag
	fptr := flag.String("f", "test.txt", "file path to read from")
	fptr1 := flag.String("fs", "test.txt", "file path to read from")
	flag.Parse()

	// read an entire file into memory
	data, err := ioutil.ReadFile(*fptr)
	if err != nil {
		fmt.Println("File reading error", err)
		return
	}

	fmt.Println("Contents of test.txt file:", string(data))

	// read a file in small chunks
	fmt.Println("---------------------------")
	fmt.Println("Contents of test1.txt file in chunks of 3-bytes:")
	f, err := os.Open(*fptr1)
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		if err = f.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	stream := bufio.NewReader(f)
	chunk := make([]byte, 3)
	for {
		n, err := stream.Read(chunk)
		if err != nil {
			fmt.Println("Error reading file:", err)
			break
		}
		fmt.Println(string(chunk[0:n]))
	}

	//
	fmt.Println("---------------------------")
	fmt.Println("Contents of test1.txt file in line by line:")
	f1, err := os.Open(*fptr)
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		if err = f1.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	scanner := bufio.NewScanner(f1)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
	err = scanner.Err()
	if err != nil {
		log.Fatal(err)
	}
}
