// Package day24 - Solution for Advent of Code 2015/24
// Problem Link: http://adventofcode.com/2015/day/24
package day24

import (
	_ "embed"
	"fmt"
	"github.com/code-shoily/aocgo/algo"
	"github.com/code-shoily/aocgo/io"
	"github.com/code-shoily/aocgo/seq"
	"math"
)

//go:embed input.txt
var input string

// Run prints out the result of the solution.
func Run() {
	fmt.Println(solve(input))
}

func solve(input string) (int, int) {
	data := io.SplitIntLines(input)
	total := seq.Sum(data)
	return getQE(data, total/3), getQE(data, total/4)
}

func getQE(packages []int, target int) int {
	minSize, minQE := math.MaxInt, math.MaxInt
	subsets := algo.SubsetSum(packages, target)

	for _, s := range subsets {
		if size := len(s); size < minSize {
			minSize = size
		}
	}

	for _, s := range subsets {
		if qe := seq.Product(s); len(s) == minSize && qe < minQE {
			minQE = seq.Product(s)
		}
	}

	return minQE
}
