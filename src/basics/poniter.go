package basics

import "fmt"

// PointerFn is simple pointer method
func PointerFn() {
	fmt.Println("--- Pointer basics --- ")

	val := 10
	var ptr *int
	fmt.Println("Ptr =", ptr)

	ptr = &val
	fmt.Println("Ptr val =", *ptr)

	changePtr(ptr)
	fmt.Println("val =", val, ", *ptr =", *ptr)

	var arr = [...]int{23, 34, 41}

	fmt.Println(arr)
	updateArr(arr[:])
	fmt.Println(arr)
}

func updateArr(arr []int) {
	arr[0] = 1
}

func changePtr(ptr *int) {
	*ptr = 1810
}
