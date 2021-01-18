package main

import (
	"fmt"
	"sync"
)

func finished() {
	fmt.Println("Finished finding largest!")
}

func findLargest(nums []int) int {
	defer finished()
	fmt.Println("Started finding largest:")

	max := nums[0]
	for _, v := range nums {
		if v > max {
			max = v
		}
	}

	fmt.Println("Largest number:", max)
	return max
}

type person struct {
	firstName string
	lastName  string
}

func (p person) fullName() string {
	fmt.Printf("%s %s\n", p.firstName, p.lastName)
	return p.firstName + p.lastName
}

func methodApply() {
	p := person{
		"Vuong",
		"Le",
	}
	defer p.fullName()
	fmt.Printf("Welcome ")
}

func printNum(a int) {
	fmt.Println("Value of num in defered function:", a)
}

func argEvaluation() {
	a := 5
	printNum(a)
	a = 10
	fmt.Println("Value of a before deferred function call:", a)
}

func stackOfDefers() {
	name := "Vuong"
	fmt.Printf("Original name: %s\n", name)
	fmt.Printf("Reversed string: ")
	defer fmt.Printf("\n")
	for _, v := range []rune(name) {
		defer fmt.Printf("%c", v)
	}
}

type rect struct {
	length int
	width  int
}

func (r rect) area(wg *sync.WaitGroup) {
	defer wg.Done()
	if r.length < 0 {
		fmt.Printf("rect %v's length should be greater than zero\n", r)
		// wg.Done()
		return
	}

	if r.width < 0 {
		fmt.Printf("rect %v's width should be greater than zero\n", r)
		// wg.Done()
		return
	}

	area := r.length * r.width
	fmt.Printf("rect %v's area %d\n", r, area)
	// wg.Done()
}

func wgUse() {
	var wg sync.WaitGroup
	r1 := rect{-34, 5}
	r2 := rect{55, -45}
	r3 := rect{5, 8}
	rects := []rect{r1, r2, r3}
	for _, v := range rects {
		wg.Add(1)
		go v.area(&wg) // the area method is called as concurrent gorountine
	}

	wg.Wait()
	fmt.Println("All go routines finished executing")
}

func main() {
	// Defer sample
	nums := []int{78, 109, 2, 563, 300}
	max := findLargest(nums)
	fmt.Println("Max:", max)

	// Apply to methods
	methodApply()

	// Argument evaluation
	argEvaluation()

	// stack of defers
	stackOfDefers()

	// Practical use of defer
	wgUse()
}
