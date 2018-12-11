package basics

import (
	"fmt"
	"unsafe"
)

func dataTypeTest() {
	// var valid bool
	valid := true
	fmt.Println(valid)

	// 	Numeric Types
	// int8, int16, int32, int64, int
	var age int8 = 10
	fmt.Println(age)

	fmt.Println(unsafe.Sizeof(age))
	// uint8, uint16, uint32, uint64, uint
	// float32, float64
	// complex64, complex128
	// byte
	// rune

	const pie = 3.14
	fmt.Println(pie)

	var comp = 5 + 6i
	fmt.Println(comp)

	// const a = 5
	// var intVar int = a
	// var int32Var int32 = a
	// var float64Var float64 = a
	// var complex64Var complex64 = a
	// fmt.Println("intVar", intVar, "\nint32Var", int32Var, "\nfloat64Var", float64Var, "\ncomplex64Var", complex64Var)

	var a = 5.9 / 8
	fmt.Println("%T >> %v", a, a)
}
