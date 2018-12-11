package basics

import "fmt"

// LoopFn is
func LoopFn() {
	fmt.Println("For Loops")
	for i := 0; i < 10; i++ {
		fmt.Print(i, " ")
	}

	fmt.Println()
	i := 0
	for ; i < 10; i++ {
		fmt.Print(i, " ")
	}
}
