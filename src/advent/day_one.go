package advent

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

// CalcFrequency is to calc frequency
func CalcFrequency() {
	fmt.Println("Frequency calculator")

	if file, err := os.Open("/Users/sakthivel/data/sakthi/sandbox/go-workspace/src/advent/inputs/day_one.txt"); err != nil {
		log.Fatal(err)
		defer file.Close()
	} else if scanner := bufio.NewScanner(file); scanner == nil {
		log.Fatal(" scanner is nil")
	} else {
		sum := 0
		num := 0
		for scanner.Scan() {
			num, _ = strconv.Atoi(scanner.Text())
			sum += num
		}
		fmt.Println("sum =", sum)
	}

}
