package basics

import (
	"fmt"
)

// Employee is model
type Employee struct {
	firstName string
	lastName  string
	age       int
	salary    float64
}

// EmployeeMgmtFn is employee management function
func EmployeeMgmtFn() {
	fmt.Println("--- Employee details ---")
	emp := Employee{"Sakthivel", "Balasubramaniam", 24, 12345.0}
	fmt.Println(emp.firstName)
}
