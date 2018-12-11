package basics

import "fmt"

// SwitchFn is
func SwitchFn() {
	fmt.Println("Switch Func")

	name := "sak"

	switch name {

	case "sakthi", "sak":
		fmt.Println("Name is Sakthivel.........")
	default:
		fmt.Println("def")
	}
}
