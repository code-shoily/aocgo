// Package day{{.Day}} - Solution for Advent of Code {{.Year}}/{{.Day}}
// Problem Link: http://adventofcode.com/{{.Year}}/day/{{.Day}}
package day{{.Day}}

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
	data := parse(input)
	return solvePart1(data), solvePart2(data)
}

func solvePart1(data []string) int {
	return len(data)
}

func solvePart2(data []string) int {
	return len(data)
}

func parse(input string) (data []string) {
	return utils.SplitLines(input)
}
