package main

import (
	"fmt"
	"golang/src/basics"
	"golang/src/concept"
	"golang/src/rectangle"
	//"snippets"
)

func main() {
	concept.Test()
	// sandbox()
	//snippets.Sandbox()
	//snippets.PatternMatcher()
	// basics.MethodsFn()
}

func sandbox() {

	choice := 8
	switch choice {
	case 1:
		fmt.Println("--- RECTANGLE ---")
		length, width := 10.3, 5.23
		fmt.Println("Area =", rectangle.Area(length, width))
		fmt.Println("Diagonal =", rectangle.Diagonal(length, width))

	case 2:
		fmt.Println("----------------------")
		basics.LoopFn()

	case 3:
		fmt.Println("----------------------")
		basics.SwitchFn()

	case 4:
		fmt.Println("----------------------")
		basics.ArrayFn()

	case 5:
		fmt.Println("----------------------")
		basics.VariadicFn()

	case 6:
		fmt.Println("----------------------")
		basics.MapFn()

	case 7:
		fmt.Println("----------------------")
		basics.PointerFn()

	case 8:
		fmt.Println("-----------------------")
		basics.EmployeeMgmtFn()
	default:
		fmt.Println("Selected feature not implemented...")
	}

}
