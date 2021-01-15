package department

import "fmt"

type department struct {
	name           string
	numOfEmployees int
}

func New(name string, numOfEmployees int) department {
	d := department{
		name,
		numOfEmployees,
	}
	return d
}

func (d department) ShowInfo() {
	fmt.Printf("Department %s with %d employee(s)\n", d.name, d.numOfEmployees)
}
