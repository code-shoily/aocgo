// Package day01 - Solution for Advent of Code 2024/01
// Problem Link: http://adventofcode.com/2024/day/01
package day01

import (
	_ "embed"
	"fmt"
	"log"
	"math"
	"sort"
	"strconv"
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

func solvePart1(data [][]int) (result int) {
	sort.IntSlice(data[0]).Sort()
	sort.IntSlice(data[1]).Sort()

	for idx := range data[0] {
		result += int(math.Abs(float64(data[0][idx] - data[1][idx])))
	}

	return result
}

func solvePart2(data [][]int) (result int) {
	var frequencies = seq.Frequencies(data[1])
	for _, v := range data[0] {
		frequency, ok := frequencies[v]
		if ok {
			result += v * frequency
		}
	}
	return result
}

func parse(input string) [][]int {
	var leftVals, rightVals []int

	for _, line := range strings.Split(input, "\n") {
		tokens := strings.Fields(line)

		leftVal, err := strconv.Atoi(tokens[0])
		if err != nil {
			log.Fatalf("Failed to parse left value %q: %v", tokens[0], err)
		}

		rightVal, err := strconv.Atoi(tokens[1])
		if err != nil {
			log.Fatalf("Failed to parse right value %q: %v", tokens[1], err)
		}
		leftVals = append(leftVals, leftVal)
		rightVals = append(rightVals, rightVal)
	}

	return [][]int{leftVals, rightVals}
}
