// Package day05 - Solution for Advent of Code 2017/05
// Problem Link: http://adventofcode.com/2017/day/05
package day05

import (
	_ "embed"
	"fmt"
	"github.com/code-shoily/aocgo/utils"
)

//go:embed input.txt
var input string

// Run prints out the result of the solution.
func Run() {
	fmt.Println(solve(input))
}

func solve(input string) (int, int) {
	data1 := parse(input)
	data2 := make([]int, len(data1))
	copy(data2, data1)

	return solvePart1(data1), solvePart2(data2)
}

func solvePart1(data []int) (steps int) {
	var current int
	for {
		if current >= len(data) {
			return steps
		}

		if offset := data[current]; offset == 0 {
			data[current]++
		} else {
			data[current]++
			current += offset
		}

		steps++
	}
}

func solvePart2(data []int) (steps int) {
	var current int
	for {
		if current >= len(data) {
			return steps
		}

		switch offset := data[current]; {
		case offset == 0:
			data[current]++
		case offset >= 3:
			data[current]--
			current += offset
		default:
			data[current]++
			current += offset
		}

		steps++
	}
}

func parse(input string) (data []int) {
	return utils.SplitIntLines(input)
}
