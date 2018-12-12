package main

import (
	"advent"
	"basics"
	"fmt"
	"rectangle"
)

func main() {
	isSandbox := false

	if isSandbox {
		sandbox()
	} else {
		adventOfCode()
	}

	//snippets.PatternMatcher()
}

func adventOfCode() {
	day := 1
	switch day {
	case 1:
		advent.CalcFrequency()
	case 2:
		advent.CalcThreshold()
	case 3:
		advent.ClaimMatrix()
	default:
	}
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
