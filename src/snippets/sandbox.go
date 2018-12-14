package snippets

import (
	"fmt"
	"strings"
)

//Sandbox is func
func Sandbox() {
	data := "dabAcCaCBAcCcaDA"
	fmt.Println(strings.Contains(data, string(data[2])))

	for lowCaps, bigCaps := 'a', 'A'; lowCaps <= 'z'; {

		cond := strings.Contains(data, string(lowCaps)) || strings.Contains(data, string(bigCaps))
		lowCaps++
		bigCaps++
	}
}
