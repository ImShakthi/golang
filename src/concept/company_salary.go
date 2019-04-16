package concept

// import ("fmt")

// SalaryCalculator interface
type SalaryCalculator interface {
	CalculateSalary() int
}

// Permanent struct
type Permanent struct {
	empID    int
	basicPay int
	pf       int
}

// Contract struct
type Contract struct {
	empID    int
	basicPay int
}

// CalculateSalary func
func (p Permanent) CalculateSalary() int {
	return p.basicPay + p.pf
}

// CalculateSalary func for contract
func (c Contract) CalculateSalary() int {
	return c.basicPay
}
