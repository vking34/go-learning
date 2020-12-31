package main

import "fmt"

type employee struct {
	salary  int
	country string
}

func main() {

	// creation
	salary := make(map[string]int)
	fmt.Println(salary)

	// addition
	salary["steve"] = 1200
	salary["mike"] = 750
	fmt.Println(salary)

	// access
	emp := "mike"
	employeeSalary := salary[emp]
	fmt.Println("Salary of", emp, "is", employeeSalary)

	// checking if a key exists
	employeeSalary1, existing := salary["travis"]
	fmt.Println("travis exists:", existing, "and salary:", employeeSalary1)

	// iterate
	for name, wage := range salary {
		fmt.Println("name:", name, ", wage:", wage)
	}

	// deleting items from a map
	delete(salary, "steve")
	fmt.Println(salary)

	// map of structs
	emp1 := employee{
		salary:  2000,
		country: "USA",
	}

	emp2 := employee{
		salary:  1200,
		country: "Vietnam",
	}

	employeeInfo := map[string]employee{
		"Steve": emp1,
		"Mike":  emp2,
	}

	for name, info := range employeeInfo {
		fmt.Println("name:", name, "country:", info.country, "salary:", info.salary)
	}

	// length of map
	fmt.Println("length is:", len(employeeInfo))

	// modify map
	modified := employeeInfo
	modified["Mike"] = employee{
		salary:  700,
		country: "VN",
	}
	fmt.Println("Emmployee salary changed:", employeeInfo)
}
