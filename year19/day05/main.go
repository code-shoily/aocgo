// Package day05 - Solution for Advent of Code 2019/05
// Problem Link: http://adventofcode.com/2019/day/05
package day05

import (
	_ "embed"
	"fmt"
	"github.com/code-shoily/aocgo/year19/intcode"
)

//go:embed input.txt
var input string

// Run prints out the result of the solution.
func Run() {
	fmt.Println(solve(input))
}

func solve(input string) (int, int) {
	return solvePart1(input), solvePart2(input)
}

func solvePart1(input string) int {
	program := intcode.InitializeProgram(input, 1)
	program.Run()
	return program.Output()
}

func solvePart2(input string) int {
	program := intcode.InitializeProgram(input, 5)
	program.Run()
	return program.Output()
}
