// Package day08 - Solution for Advent of Code 2015/08
// Problem Link: http://adventofcode.com/2015/day/08
package day08

import (
	_ "embed"
	"fmt"
	"strings"
)

//go:embed input.txt
var input string

// Run prints out the result of the solution.
func Run() {
	fmt.Println(solve(input))
}

func solve(input string) (int, int) {
	data := strings.Split(input, "\n")
	return totalEncodedDiffs(data, encodeAsInMem), totalEncodedDiffs(data, encodeAsString)
}

func totalEncodedDiffs(data []string, encode func(string) int) (total int) {
	for _, line := range data {
		total += encode(line)
	}
	return total
}

func encodeAsInMem(line string) (inMemLen int) {
	originalLen := len(line)
	for i := 0; i < originalLen-1; i++ {
		switch line[i] {
		case '"':
		case '\\':
			inMemLen++
			if nextChar := line[i+1]; nextChar == '\\' || nextChar == '"' {
				i++
			} else if nextChar == 'x' {
				i += 3
			}
		default:
			inMemLen++
		}
	}

	return originalLen - inMemLen
}

func encodeAsString(line string) (stringLen int) {
	for _, char := range line {
		if char == '"' || char == '\\' {
			stringLen += 2
		} else {
			stringLen++
		}
	}

	return (stringLen + 2) - len(line)
}
