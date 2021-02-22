package main

import (
	"fmt"
)

func find(num int, nums ...int) {
	fmt.Printf("type of nums is %T\n", nums)
	found := false

	for i, v := range nums {
		if v == num {
			fmt.Println(num, "found at index", i, "in", nums)
			found = true
		}
	}

	if !found {
		fmt.Println(num, "not found in", nums)
	}
	fmt.Printf("\n")
}

func modify(s ...string) {
	s[0] = "Go"
}

func change(s ...string) {
	s[0] = "beginner"
	s = append(s, "go")
	fmt.Println(s)
	fmt.Println("len:", len(s), ",capacity:", cap(s))
}

func main() {
	find(34, 56, 34, 65)
	find(56, 12, 45, 56, 56)
	find(11, 22, 33, 44)
	find(78)

	// can not passed a slice to the variadic argument
	// we need ... operator to pass a slice
	nums := []int{45, 87, 23, 78}
	find(23, nums...)

	// modify a slice in variadic function
	welcome := []string{"hello", "world"}
	modify(welcome...)
	fmt.Println(welcome)

	//
	change(welcome...)
	fmt.Println(welcome)
	fmt.Println("len:", len(welcome), ",capacity:", cap(welcome))
	welcome = append(welcome, "vuong")
	fmt.Println(welcome)
	fmt.Println("len:", len(welcome), ",capacity:", cap(welcome))
}
