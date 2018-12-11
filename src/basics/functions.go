package basics

import (
	"fmt"
)

func main() {
	price, no := 100, 2
	fmt.Println("Function module....")
	fmt.Println("Total price =", calculateBill(price, no))

	length, width := 10.0, 5.0
	fmt.Println("Length =", length, ", width=", width)
	area, perimeter := reactPropsParams(length, width)
	fmt.Println("Area =", area, ", Perimeter =", perimeter)

	a, _ := reactProps(length, width)
	fmt.Println("Area =", a)
}

func calculateBill(price int, no int) int {
	return price * no
}

func reactProps(length, width float64) (float64, float64) {
	area := length * width
	perimeter := (length + width) * 2
	return area, perimeter
}

func reactPropsParams(length, width float64) (area, perimeter float64) {
	area = length * width
	perimeter = (length + width) * 2
	return
}
