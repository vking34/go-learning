package main

import (
	"oop/department"
	"oop/employee"
)

func main() {
	// Structs instead of classes
	emp1 := employee.Employee{
		FirstName:   "Vuong",
		LastName:    "Le",
		TotalLeaves: 30,
		LeavesTaken: 20,
	}

	emp1.LeavesRemaining()

	// NewT function instead of constructor
	d1 := department.New("IT", 13)
	d1.ShowInfo()
}
