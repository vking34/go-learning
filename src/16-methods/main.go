package main

import "fmt"

type Employee struct {
	name     string
	salary   int
	currency string
}

func (e Employee) displaySalary() {
	fmt.Printf("Salary of %s is %s%d\n", e.name, e.currency, e.salary)
}

func (e Employee) changeName(newName string) {
	e.name = newName
}

func (e *Employee) changeSalary(newSalary int) {
	e.salary = newSalary
}

type address struct {
	district string
	city     string
}

func (a address) getFullAdress() {
	fmt.Printf("Full address: %s, %s\n", a.district, a.city)
}

type person struct {
	name string
	age  int
	address
}

type retangle struct {
	length int
	width  int
}

func area(r retangle) {
	fmt.Println("Area:", r.length*r.width)
}

func (r retangle) area() {
	fmt.Println("Area:", r.length*r.width)
}

func (r retangle) extend() {
	r.length++
	r.width++
}

func (r retangle) getSize() {
	fmt.Printf("Length: %d, width: %d\n", r.length, r.width)
}

type myInt int

func (a myInt) add(b myInt) myInt {
	return a + b
}

func main() {

	// Declaration
	emp1 := Employee{
		name:     "Vuong Le",
		salary:   800,
		currency: "$",
	}

	emp1.displaySalary()

	// Pointer receiver vs value receiver
	fmt.Println("Name before:", emp1.name)
	emp1.changeName("Truong")
	fmt.Println("Name after:", emp1.name)

	fmt.Println("Salary after:", emp1.salary)
	emp1.changeSalary(1000)
	fmt.Println("Salary after:", emp1.salary)

	// Methods of anonymous struct fields
	p1 := person{
		name: "Vuong",
		age:  23,
		address: address{
			city:     "Hanoi",
			district: "Hai Ba Trung",
		},
	}

	p1.getFullAdress()

	// Method accepts both value receiver and pointer receiver
	r := retangle{
		length: 9,
		width:  7,
	}

	area(r)
	r.area()

	p := &r
	p.area()

	r.extend()
	r.getSize()
	p.extend()
	p.getSize()

	// Methods with non-struct receivers
	n1 := myInt(3)
	n2 := myInt(4)
	sum := n1.add(n2)
	fmt.Println("Sum:", sum)
}
