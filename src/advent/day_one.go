package advent

import (
	"fmt"
	"strconv"
)

// CalcFrequency is to calc frequency
func CalcFrequency() {
	fmt.Println("Frequency calculator")

	data := ReadDataFromFile("day_one.txt")
	noL := len(data)

	freqMap := make(map[int]int)
	sum := 0

	arr := []int{}
	for index := 0; index < len(data); index++ {
		num, _ := strconv.Atoi(data[index])
		arr = append(arr, num)
	}
	for i := 0; i < noL; i++ {
		sum += arr[i]
		freqMap[sum]++
	}
	fmt.Println("sum =", sum)
	fmt.Println(freqMap)

	for i := 0; i < noL; i++ {
		sum += arr[i]
		freqMap[sum]++
		if freqMap[sum] == 2 {
			fmt.Print(" Freq =", sum)
			break
		}
	}
	fmt.Println(freqMap)

}
