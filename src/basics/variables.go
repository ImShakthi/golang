package basics

import "fmt"

func testVar() {
	var age = 18
	fmt.Println("Age=", age)
	age = 10
	fmt.Println("Age=", age)

	var width, height int = 5, 10
	fmt.Println("width=", width, ", height=", height)

	var w, h = 15, 100
	fmt.Println("width=", w, ", height=", h)

	var (
		name       = "sakthi"
		currentAge = 24
		exp        int
	)
	fmt.Println(name, currentAge, exp)

	myname, myage := "sakthi", 24
	fmt.Println(myname, myage)

	myage, myexp := 24, 2.5
	fmt.Println(myage, myexp)

}
