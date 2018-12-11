package rectangle

import "fmt"

// Area is method to get area of rectangle
func Area(length, width float64) (area float64) {
	fmt.Println("Rectangle...")
	area = length * width
	return
}

// Diagonal is method to get perimeter
func Diagonal(length, width float64) (perimeter float64) {
	perimeter = (length + width) * 2
	return
}
