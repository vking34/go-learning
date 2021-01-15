package main

import "fmt"

type Describer interface {
	Describe()
}

type Job interface {
	Work()
}

type Person struct {
	name string
	age  int
}

// Embedding interface
type Employee interface {
	Describer
	Job
}

type Address struct {
	district string
	city     string
}

func (p Person) Describe() {
	fmt.Printf("%s is %d years old\n", p.name, p.age)
}

func (a *Address) Describe() {
	fmt.Printf("Address: %s, %s\n", a.district, a.city)
}

// Implementing multiple interfaces
func (p Person) Work() {
	fmt.Printf("%s is working...\n", p.name)
}

func main() {
	// Interface with pointer receiver and value receiver
	var d1 Describer
	p1 := Person{
		name: "Vuong",
		age:  23,
	}

	d1 = p1
	d1.Describe()

	d1 = &p1
	d1.Describe()

	a1 := Address{
		district: "Hai Ba Trung",
		city:     "Hanoi",
	}

	d1 = &a1
	d1.Describe()

	// Implementing multiple interfaces
	p1.Work()

	// Embedding interface
	var e1 Employee
	e1 = p1
	e1.Describe()
	e1.Work()
}
