package main

import "fmt"

func subtactOne(numbers []int) {
	for i := range numbers {
		numbers[i] -= 2
	}

}

func getCountries() []string {
	countries := []string{"USA", "Singapore", "Germany", "India", "Australia"}
	neededCountries := countries[:len(countries)-2]
	countriesCpy := make([]string, len(neededCountries))
	copy(countriesCpy, neededCountries) //copies neededCountries to countriesCpy
	return countriesCpy
}

func main() {
	var a [3]int
	fmt.Println(a)

	a1 := [3]int{23, 34, 54}
	fmt.Println(a1)

	a2 := [...]int{13, 22}
	fmt.Println(a2)

	fmt.Println("-----------------------")
	// arrays are value types
	countries := [...]string{"USA", "China", "India"}
	newCountries := countries
	newCountries[0] = "Vietnam"
	fmt.Println("old:", countries)
	fmt.Println("new:", newCountries)

	fmt.Println("-----------------------")
	// iterating
	for i := 0; i < len(a1); i++ {
		fmt.Printf("%d th element of a1 is %d\n", i, a1[i])
	}
	fmt.Println("-----------------------")
	for i, v := range a1 {
		fmt.Printf("%d th element of a1 is %d\n", i, v)
	}

	// multi-dimensional array

	fmt.Println("-----------------------")
	// slices
	a3 := [5]int{34, 23, 21, 56, 88}
	var s1 []int = a3[1:4]
	fmt.Println(s1)

	// modify

	fmt.Println("-----------------------")
	// append
	s1 = append(s1, 11)
	fmt.Println(s1)

	cars := []string{"Ferrari", "Honda", "Ford"}
	fmt.Println("cars:", cars, "has old length", len(cars), "and capacity", cap(cars)) //capacity of cars is 3
	cars = append(cars, "Toyota")
	fmt.Println("cars:", cars, "has new length", len(cars), "and capacity", cap(cars)) //capacity of cars is doubled to 6

	fmt.Println("-----------------------")
	// nil
	var names []string //zero value of a slice is nil
	if names == nil {
		fmt.Println("slice is nil going to append")
		names = append(names, "John", "Sebastian", "Vinay")
		fmt.Println("names contents:", names)
	}

	fmt.Println("-----------------------")
	// append slice
	veggies := []string{"potatoes", "tomatoes", "brinjal"}
	fruits := []string{"oranges", "apples"}
	food := append(veggies, fruits...)
	fmt.Println("food:", food)
	fmt.Printf("food length: %d, capicity: %d\n", len(food), cap(food))

	fmt.Println("-----------------------")
	// passing a slice via function
	nos := []int{8, 7, 6}
	fmt.Println("slice before function call", nos)
	subtactOne(nos)                               //function modifies the slice
	fmt.Println("slice after function call", nos) //modifications are visible outside

	fmt.Println("-----------------------")
	// multidimensional slices
	pls := [][]string{
		{"C", "C++"},
		{"JavaScript"},
		{"Go", "Rust"},
	}

	for _, v1 := range pls {
		for _, v2 := range v1 {
			fmt.Printf("%s ", v2)
		}
		fmt.Print("\n")
	}

	fmt.Println("-----------------------")
	// memory optimisation
	countriesNeeded := getCountries()
	fmt.Println(countriesNeeded)
}
