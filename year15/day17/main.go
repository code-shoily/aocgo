// Package day17 - Solution for Advent of Code 2015/17
// Problem Link: http://adventofcode.com/2015/day/17
package day17

import (
	_ "embed"
	"fmt"
	"github.com/code-shoily/aocgo/algo"
	"github.com/code-shoily/aocgo/io"
	"math"
)

//go:embed input.txt
var input string

// Run prints out the result of the solution.
func Run() {
	fmt.Println(solve(input))
}

const target = 150

func solve(input string) (int, int) {
	containerSets := algo.SubsetSum(io.SplitIntLines(input), target)
	minimum, frequency := math.MaxInt, map[int]int{}

	for _, containerSet := range containerSets {
		if size := len(containerSet); size <= minimum {
			minimum = size
			frequency[minimum]++
		}
	}

	return len(containerSets), frequency[minimum]
}
