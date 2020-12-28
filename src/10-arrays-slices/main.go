package main

import "fmt"

func main() {
	var a [3]int
	fmt.Println(a)

	a1 := [3]int{23, 34, 54}
	fmt.Println(a1)

	a2 := [...]int{13, 22}
	fmt.Println(a2)

	// arrays are value types
	countries := [...]string{"USA", "China", "India"}
	newCountries := countries
	newCountries[0] = "Vietnam"
	fmt.Println("old:", countries)
	fmt.Println("new:", newCountries)

	// iterating
	for i := 0; i < len(a1); i++ {
		fmt.Printf("%d th element of a1 is %d\n", i, a1[i])
	}
	fmt.Println("-----------------------")
	for i, v := range a1 {
		fmt.Printf("%d th element of a1 is %d\n", i, v)
	}

	// multi-dimensional array

}
