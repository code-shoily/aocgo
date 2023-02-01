// Package day02 - Solution for Advent of Code 2019/02
// Problem Link: http://adventofcode.com/2019/day/02
package day02

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
	program := intcode.InitializeProgram(input, 0)
	program.ProvideInitialParameters(12, 2)
	program.Run()
	return program.Output()
}

func solvePart2(input string) int {
	for noun := 0; noun <= 99; noun++ {
		for verb := 0; verb <= 99; verb++ {
			program := intcode.InitializeProgram(input, 0)
			program.ProvideInitialParameters(noun, verb)
			program.Run()
			if program.Output() == 19_690_720 {
				return 100*noun + verb
			}
		}
	}
	panic("Invalid instruction")
}
