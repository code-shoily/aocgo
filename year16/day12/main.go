// Package day12 - Solution for Advent of Code 2016/12
// Problem Link: http://adventofcode.com/2016/day/12
package day12

import (
	_ "embed"
	"fmt"
	"github.com/code-shoily/aocgo/year16/monorail"
	"strings"
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

func solvePart1(data []*monorail.Instruction) int {
	c := monorail.NewComputer(data, map[string]int{})
	c.Run()
	return c.Registers["a"]
}

func solvePart2(data []*monorail.Instruction) int {
	c := monorail.NewComputer(data, map[string]int{"c": 1})
	c.Run()
	return c.Registers["a"]
}

func parse(input string) (instructions []*monorail.Instruction) {
	for _, line := range strings.Split(input, "\n") {
		instructions = append(instructions, monorail.ParseInstruction(line))
	}

	return instructions
}
