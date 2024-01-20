// Package day01 - Solution for Advent of Code 2023/01
// Problem Link: http://adventofcode.com/2023/day/01
package day01

import (
	_ "embed"
	"fmt"
	"regexp"
	"strings"

	"github.com/code-shoily/aocgo/seq"
)

//go:embed input.txt
var input string

// Run prints out the result of the solution.
func Run() {
	fmt.Println(solve(input))
}

func solve(input string) (int, int) {
	data := parse(input)
	return solvePart1(data), solvePart2(data)
}

func solvePart1(data []string) int {
	var result int
	for _, line := range data {
		result += calibrationValue(line, digits, digits)
	}
	return result
}

func solvePart2(data []string) int {
	var result int
	for _, line := range data {
		result += calibrationValue(line, words, wordsRev)
	}
	return result
}

func parse(input string) (data []string) {
	return strings.Split(input, "\n")
}

var digits = regexp.MustCompile("1|2|3|4|5|6|7|8|9")
var words = regexp.MustCompile("1|2|3|4|5|6|7|8|9|one|two|three|four|five|six|seven|eight|nine")
var wordsRev = regexp.MustCompile("1|2|3|4|5|6|7|8|9|eno|owt|eerht|ruof|evif|xis|neves|thgie|enin")

func calibrationValue(line string, forward *regexp.Regexp, backward *regexp.Regexp) int {
	first := forward.FindString(line)
	last := backward.FindString(seq.ReverseString(line))

	return toDigit(first)*10 + toDigit(last)
}

func toDigit(value string) int {
	switch value {
	case "1", "one", "eno":
		return 1
	case "2", "two", "owt":
		return 2
	case "3", "three", "eerht":
		return 3
	case "4", "four", "ruof":
		return 4
	case "5", "five", "evif":
		return 5
	case "6", "six", "xis":
		return 6
	case "7", "seven", "neves":
		return 7
	case "8", "eight", "thgie":
		return 8
	case "9", "nine", "enin":
		return 9
	default:
		return 0
	}
}
