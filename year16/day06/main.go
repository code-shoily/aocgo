// Package day06 - Solution for Advent of Code 2016/06
// Problem Link: http://adventofcode.com/2016/day/06
package day06

import (
	_ "embed"
	"fmt"
	"github.com/code-shoily/aocgo/seq"
	"github.com/code-shoily/aocgo/utils"
	"math"
	"strings"
)

//go:embed input.txt
var input string

// Run prints out the result of the solution.
func Run() {
	fmt.Println(solve(input))
}

func solve(input string) (string, string) {
	data := parse(input)
	tr := seq.Transpose(data)
	slice1 := make([]string, len(tr))
	slice2 := make([]string, len(tr))

	for idx, column := range tr {
		frequencies := seq.Frequencies(column)
		min, max := minMaxFrequencies(frequencies)
		slice1[idx] = max
		slice2[idx] = min
	}

	return strings.Join(slice1, ""), strings.Join(slice2, "")
}

func parse(input string) (data [][]string) {
	for _, line := range utils.SplitLines(input) {
		data = append(data, utils.SplitBy(line, ""))
	}

	return data
}

func minMaxFrequencies(frequencies map[string]int) (min string, max string) {
	maxSoFar, minSoFar := math.MinInt, math.MaxInt
	for char, frequency := range frequencies {
		if frequency > maxSoFar {
			maxSoFar = frequency
			max = char
		}
		if frequency < minSoFar {
			minSoFar = frequency
			min = char
		}
	}

	return min, max
}
