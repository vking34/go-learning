package main

import "fmt"

type add func(a int, b int) int

func passFunc(a func(a, b int) int) {
	fmt.Println(a(60, 7))
}

func getFunc() func(a, b int) int {
	f := func(a, b int) int {
		return a + b
	}
	return f
}

func appendStr() func(string) string {
	t := "Hello"
	c := func(b string) string {
		t = t + " " + b
		return t
	}
	return c
}

type student struct {
	name  string
	grade string
}

func filter(students []student, matcher func(student) bool) []student {
	var result []student
	for _, student := range students {
		if matcher(student) {
			result = append(result, student)
		}
	}

	return result
}

func main() {
	// anonymous function
	anonymousFunc := func() {
		fmt.Println("hello worl first class function")
	}

	anonymousFunc()
	fmt.Printf("%T\n", anonymousFunc)

	// function type
	var sum add = func(a int, b int) int {
		return a + b
	}
	s0 := sum(3, 4)
	fmt.Println("sum:", s0)

	// higher-order functions
	// passing functions as arguments to other functions
	f := func(a, b int) int {
		return a + b
	}
	passFunc(f)

	// Returning functions from other functions
	s := getFunc()
	fmt.Println(s(60, 7))

	// closures
	a := 5
	func() {
		fmt.Println("a = ", a)
	}()

	// every closure is bound to its own surrounding variable.
	x := appendStr()
	y := appendStr()
	fmt.Println(x("World"))
	fmt.Println(y("Everyone"))

	fmt.Println(x("Vuong"))
	fmt.Println(y("!"))

	// practice
	s1 := student{
		"Vuong",
		"B",
	}
	s2 := student{
		"Quy",
		"A",
	}
	students := []student{s1, s2}
	r := filter(students, func(s student) bool {
		if s.grade == "B" {
			return true
		}
		return false
	})

	fmt.Println("Filter result:", r)
}
