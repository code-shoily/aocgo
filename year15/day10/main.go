// Package day10 - Solution for Advent of Code 2015/10
// Problem Link: http://adventofcode.com/2015/day/10
package day10

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

const input = "1113122113"

// Run prints out the result of the solution.
func Run() {
	fmt.Println(solve(input))
}

func solve(input string) (part1 int, _ int) {
	for i := 0; i < 50; i++ {
		if i == 40 {
			part1 = len(input)
		}
		input = encode(input)
	}

	return part1, len(input)
}

func encode(input string) string {
	var sb strings.Builder
	lastFreq, lastChar := 1, rune(input[0])
	for _, ch := range input[1:] {
		if ch == lastChar {
			lastFreq++
		} else {
			writeFrequency(&sb, lastFreq, lastChar)
			lastFreq, lastChar = 1, ch
		}
	}

	writeFrequency(&sb, lastFreq, lastChar)

	return sb.String()
}

func writeFrequency(sb *strings.Builder, freq int, ch rune) {
	sb.WriteString(strconv.Itoa(freq))
	sb.WriteRune(ch)
}
