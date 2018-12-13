package snippets

import (
	"advent"
	"fmt"
	"regexp"
)

// (?P<>[0-9]*)

const timeStampPattern = "\\[[0-9]*-[0-9]*-[0-9]*\\s[0-9]*:(?P<MINUTES>[0-9]*)\\]\\s"
const guardPattern = "Guard #(?P<GUARDID>[0-9]*) begins shift"
const wakesup = "wakes up"
const fallAsleep = "falls asleep"

var tsMetaData = regexp.MustCompile(timeStampPattern + guardPattern)
var wsMetaData = regexp.MustCompile(timeStampPattern + wakesup)
var fsMetaData = regexp.MustCompile(timeStampPattern + fallAsleep)

var s1 = "[1518-11-23 00:04] Guard #1663 begins shift"
var s2 = "[1518-11-23 00:19] falls asleep"
var s3 = "[1518-11-23 00:56] wakes up"

// PatternMatcher is function to play with pattern matcher
func PatternMatcher() {
	fmt.Println("--- Pattern Matcher --- ")

	input := s1
	if tsMetaData.Match([]byte(input)) {
		matches := tsMetaData.FindStringSubmatch(input)
		minutes := matches[1]
		guardID := matches[2]
		fmt.Println("minutes =", minutes, ", guardID =", guardID)
	} else if wsMetaData.Match([]byte(input)) {
		matches := wsMetaData.FindStringSubmatch(input)
		minutes := matches[1]
		fmt.Println("FALL minutes =", minutes)
	} else if fsMetaData.Match([]byte(input)) {
		matches := fsMetaData.FindStringSubmatch(input)
		minutes := matches[1]
		fmt.Println("WAKES minutes =", minutes)
	}

	guardMap := make(map[int]advent.NightWatch)
	fmt.Println(">>>>> ", guardMap[123])
}

/**
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
**/
