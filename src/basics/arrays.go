package basics

import "fmt"

// ArrayFn is
func ArrayFn() {
	fmt.Println("Array fn")
	arr := [5]int{1, 2, 3}
	fmt.Println(arr)

	strArr := [...]string{"ABC", "India", "three", "four", "five"}
	fmt.Println(strArr)

	for i := 0; i < len(strArr); i++ {
		fmt.Println(strArr[i])
	}

	for _, str := range strArr {
		fmt.Println(str)
	}

	twoStrArr := [...][3]string{
		{"abc", "def"},
		{"abcwe", "defwe"},
		{"qwabc", "qwdef"},
	}

	for _, strArr := range twoStrArr {
		for _, str := range strArr {
			fmt.Print(str, " ")
		}
		fmt.Println()
	}
	fmt.Println(twoStrArr)

	slicedStrArr := strArr[1:3]
	fmt.Println("len =", len(slicedStrArr), ", cap =", cap(slicedStrArr), ", >>", slicedStrArr)

	makStrArr := make([]int, 3, 3)
	fmt.Println("makStrArr =", makStrArr)
}
