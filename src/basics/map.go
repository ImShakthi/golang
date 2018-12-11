package basics

import (
	"fmt"
)

// MapFn is
func MapFn() {
	fmt.Println(">> Map <<")
	colorCodeMap := make(map[string]string)
	colorCodeMap["major"] = "red"
	colorCodeMap["minor"] = "blue"
	colorCodeMap["cleared"] = "green"
	colorCodeMap["info"] = "white"

	fmt.Println(colorCodeMap)

	sampleMap := map[string]int{
		"red":  1230,
		"blue": 1,
	}
	fmt.Println(sampleMap["red"], ">>", sampleMap["sakhi"])
	value, ok := colorCodeMap["major"]
	fmt.Println(value, ok)
	fmt.Println(colorCodeMap["major"])

	for key, value := range colorCodeMap {
		fmt.Println("key =", key, ", value =", value)
	}

	fmt.Println("Before delete =", colorCodeMap, ">>>> length =", len(colorCodeMap))
	newCcMap := (colorCodeMap)
	delete(colorCodeMap, "major")
	fmt.Println("After delete =", colorCodeMap, ">>>> length =", len(colorCodeMap), " #newColorCodeMap =", newCcMap)
}
