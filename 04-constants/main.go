package main

import "fmt"

// constants
func main() {
	const x = 50
	const y string = "b"
	fmt.Println("constant: ", x)
	fmt.Println("constant: ", y)

	const (
		name    = "vuong"
		age     = 34
		country = "VN"
	)

	fmt.Println("Name:", name)
	fmt.Println("Age:", age)
	fmt.Println("Country:", country)

	// string
	const n = "Sam"
	var person = n
	fmt.Printf("type %T value %v\n", person, person)

	// numberic
	var i = 5
	var f = 5.6
	var c = 5 + 6i
	fmt.Printf("i's type is %T, f's type is %T, c's type is %T\n", i, f, c)

	const a = 5
	var intVar int = a
	var int32Var int32 = a
	var float64Var float64 = a
	var complex64Var complex64 = a
	fmt.Println("intVar", intVar, "\nint32Var", int32Var, "\nfloat64Var", float64Var, "\ncomplex64Var", complex64Var)
}
