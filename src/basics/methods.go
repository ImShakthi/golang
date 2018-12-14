package basics

import "fmt"

type person struct {
	name string
	age  int
}
type rectangle struct {
	length float64
	width  float64
}

type circle struct {
	radius float64
}

func (c circle) area() float64 {
	return 3.14 * c.radius
}

func (r rectangle) area() float64 {
	return r.length * r.width
}

func (p person) updateAge() {
	p.age++
	fmt.Println("Value updated age =", p.age)
}

func (p *person) updateAge2() {
	p.age++
	fmt.Println("refer updated age =", p.age)
}

// MethodsFn method function
func MethodsFn() {
	fmt.Println("--- Methods function ---")
	r := rectangle{10.5, 12.0}
	fmt.Println(r.area())

	c := circle{5}
	fmt.Println(c.area())

	fmt.Println(">>>>> POINTER <<<<<")
	p := person{"sakthi", 24}
	p.updateAge()
	fmt.Println(p)

	(&p).updateAge2()
	fmt.Println(p)

	testStr := "aAbBcCdD"

	for index := 0; index < len(testStr); index++ {
		fmt.Print(testStr[index], " ")
	}
	//65, 97
}
