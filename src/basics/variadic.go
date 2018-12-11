package basics

import "fmt"

// VariadicFn is
func VariadicFn() {
	colors := []string{"red", "blue", "green", "yellow"}
	fmt.Println("> Variadic function <")
	getSum(1, 2, 3)
	change(colors...)
}

func getSum(nums ...int) (sum int) {
	sum = 0
	for _, n := range nums {
		sum += n
	}
	fmt.Println("Sum =", sum)
	return
}

func change(colors ...string) {
	for index, color := range colors {
		fmt.Println(index, ">", color)
	}
}
