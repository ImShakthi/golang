package snippets

import (
	"fmt"
	"regexp"
)

// PatternMatcher is function to play with pattern matcher
func PatternMatcher() {
	fmt.Println("--- Pattern Matcher --- ")
	//#5 @ 696,820: 26x19

	inputStr := "#5 @ 696,820: 26x19"
	pattern := "#(?P<CalimID>[0-9]*)\\s@\\s(?P<X>[0-9]*),(?P<Y>[0-9]*):\\s(?P<ROWS>[0-9]*)x(?P<COLS>[0-9]*)"
	claimMetaData := regexp.MustCompile(pattern)
	matches := claimMetaData.FindStringSubmatch(inputStr)
	// keys := claimMetaData.SubexpNames()

	// fmt.Println("matches =", matches)
	for _, match := range matches {
		fmt.Println(match)

	}
	// fmt.Println(claimMatcher.)
}
