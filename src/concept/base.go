package concept

import (
	"fmt"
)

// Test funcc
func Test() {
	calculateExpense()
}

func vowelFinder() {
	fmt.Printf("Interfaces!")
	name := MyString("sakthivel")
	var v VowelFinder
	// MyString implements VowelFinder
	v = name
	fmt.Printf("\nVowels in %c ", v.FindVowels())
}

func calculateExpense() {
	fmt.Println("~~~~ Calculate Expense ~~~~")
	p1 := Permanent{101, 5000, 150}
	p2 := Permanent{102, 6500, 165}
	c1 := Contract{1001, 2500}
	c2 := Contract{1002, 2500}

	employees := []SalaryCalculator{p1, p2, c1, c2}
	fmt.Println("Salary=", calculateSalary(employees))
}

func calculateSalary(emps []SalaryCalculator) int {
	totalExpense := 0
	for _, e := range emps {
		totalExpense += e.CalculateSalary()
	}
	return totalExpense
}
