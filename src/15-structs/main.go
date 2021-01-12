package main

import (
	"15-structs/computer"
	"fmt"
)

// Employee is a normal struct
type Employee struct {
	firstName string
	lastName  string
	age       int
	salary    int
}

// Person struct with anonymous fields
type Person struct {
	string
	int
}

// Address struct used for the nested struct
type Address struct {
	city     string
	district string
}

// Member struct for nested struct example
type Member struct {
	name    string
	age     int
	address Address
}

// Building struct for promoted field
type Building struct {
	name   string
	floors int
	Address
}

type name struct {
	firstName string
	lastName  string
}

func main() {

	// Creating struct
	emp1 := Employee{
		firstName: "Vuong",
		salary:    800,
		age:       23,
		lastName:  "Le",
	}

	emp2 := Employee{"Son", "Hong", 24, 750}

	fmt.Println("Employee 1:", emp1)
	fmt.Println("Employee 2:", emp2)

	// Anonymous struct
	emp3 := struct {
		firstName string
		lastNanme string
		address   string
	}{
		firstName: "Truong",
		lastNanme: "Nguyen",
		address:   "Dong Thanh",
	}
	fmt.Println("Employee 3:", emp3)

	// Accessing field
	fmt.Println("First name of employee 1:", emp1.firstName)
	fmt.Println("Age of employee 1:", emp1.age)

	// Pointer
	emp4 := &Employee{
		firstName: "Tu",
		lastName:  "Pham",
		age:       23,
		salary:    700,
	}
	fmt.Println("First name of employee 4:", (*emp4).firstName)
	fmt.Println("Age of employee 1:", emp1.age)

	// Anonymous fields
	p1 := Person{
		string: "Huy",
		int:    23,
	}
	fmt.Println("Person 1, name:", p1.string, "age:", p1.int)

	// Nested struct
	m1 := Member{
		name: "Vuong",
		age:  23,
		address: Address{
			city:     "Hanoi",
			district: "Hai Ba Trung",
		},
	}
	fmt.Println("Member 1:", m1)
	fmt.Println("From city:", m1.address.city)

	// Promoted fields
	b1 := Building{
		name:   "HAPU",
		floors: 20,
		Address: Address{
			city:     "Hanoi",
			district: "Thanh Xuan",
		},
	}

	fmt.Println("Building 1 in district:", b1.district) // promoted field

	// Use exported struct
	spec := computer.Spec{
		Maker: "Apple",
		Price: 1000,
	}
	fmt.Println("Computer:", spec)

	// Struct equality
	name1 := name{
		firstName: "Steve",
		lastName:  "Jobs",
	}
	name2 := name{
		firstName: "Steve",
		lastName:  "Jobs",
	}
	if name1 == name2 {
		fmt.Println("name1 and name2 are equal")
	} else {
		fmt.Println("name1 and name2 are not equal")
	}

	name3 := name{
		firstName: "Steve",
		lastName:  "Jobs",
	}
	name4 := name{
		firstName: "Steve",
	}

	if name3 == name4 {
		fmt.Println("name3 and name4 are equal")
	} else {
		fmt.Println("name3 and name4 are not equal")
	}
}
