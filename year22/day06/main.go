// Package day06 - Solution for Advent of Code 2022/06
// Problem Link: http://adventofcode.com/2022/day/06
package day06

import (
	_ "embed"
	"fmt"
)

//go:embed input.txt
var input string

// Run prints out the result of the solution.
func Run() {
	fmt.Println(solve(input))
}

func solve(input string) (int, int) {
	return findMarker(input, 4), findMarker(input, 14)
}

func findMarker(data string, size int) int {
	for i := 0; i+size < len(data); i++ {
		if hasNoDupes(data[i : i+size]) {
			return i + size
		}
	}
	panic("Processing error")
}

func hasNoDupes(text string) bool {
	elements := make(map[rune]bool)
	for _, letter := range text {
		elements[letter] = true
	}
	return len(elements) == len(text)
}
