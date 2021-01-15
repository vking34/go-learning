package main

import "fmt"

//
type VowelsFinder interface {
	FindVowels() []rune
}

type MyString string

func (str MyString) FindVowels() []rune {
	var vowels []rune

	for _, rune := range str {
		if rune == 'a' || rune == 'e' || rune == 'o' || rune == 'u' || rune == 'i' {
			vowels = append(vowels, rune)
		}
	}

	return vowels
}

func printVowels(vowels []rune) {
	fmt.Print("Vowels: ")
	for _, rune := range vowels {
		fmt.Printf("%s ", string(rune))
	}
	fmt.Print("\n")
}

//

type SalaryCalculator interface {
	CalculateSalary() int
}

type PermanentEmployee struct {
	empId    int
	basicPay int
	pf       int
}

type ContractEmployee struct {
	empId    int
	basicPay int
}

// Freelancer new struct
type Freelancer struct {
	empId       int
	ratePerHour int
	totalHours  int
}

func (p PermanentEmployee) CalculateSalary() int {
	return p.basicPay + p.pf
}

func (c ContractEmployee) CalculateSalary() int {
	return c.basicPay
}

func (f Freelancer) CalculateSalary() int {
	return f.ratePerHour * f.totalHours
}

func totalExpense(s []SalaryCalculator) {
	expense := 0
	for _, employee := range s {
		expense += employee.CalculateSalary()
	}
	fmt.Printf("Total expense per month $%d\n", expense)
}

func describe(c SalaryCalculator) {
	fmt.Printf("Interface type %T value %v\n", c, c)
}

func describeAny(i interface{}) {
	fmt.Printf("Type = %T, value = %v\n", i, i)
}

func assert(i interface{}) {
	v, ok := i.(int)
	fmt.Println(v, ok)
}

func main() {
	// Declaration
	name := MyString("Vuong")
	var v VowelsFinder
	v = name
	vowels := v.FindVowels()
	fmt.Printf("Vowels: %c\n", vowels)
	printVowels(vowels)

	// Practice
	pEmp1 := PermanentEmployee{
		empId:    1,
		basicPay: 500,
		pf:       50,
	}
	pEmp2 := PermanentEmployee{
		empId:    2,
		basicPay: 600,
		pf:       100,
	}
	pEmp3 := PermanentEmployee{
		empId:    3,
		basicPay: 450,
		pf:       100,
	}

	cEmp1 := ContractEmployee{
		empId:    4,
		basicPay: 300,
	}

	// add new employee
	fEmp1 := Freelancer{
		empId:       5,
		ratePerHour: 7,
		totalHours:  40,
	}

	employees := []SalaryCalculator{pEmp1, pEmp2, pEmp3, cEmp1, fEmp1}
	totalExpense(employees)

	// Interface as Tuple
	var calculator SalaryCalculator = pEmp1
	describe(calculator)

	// Empty interface
	s := "Hello world"
	describeAny(s)

	i := 34
	describeAny(i)

	describeAny(cEmp1)

	// Type Assertion
	var num interface{} = 34
	assert(num)

	var noun interface{} = "Screen"
	assert(noun)

}
