// Package day08 - Solution for Advent of Code 2020/08
// Problem Link: http://adventofcode.com/2020/day/08
package day08

import (
	_ "embed"
	"fmt"
	"strings"
)

//go:embed input.txt
var input string

// Run prints out the result of the solution.
func Run() {
	fmt.Println(solve(input))
}

func solve(input string) (int, int) {
	instructions := parse(input)
	return solvePart1(instructions), solvePart2(instructions)
}

func solvePart1(instructions []instruction) int {
	c := newConsole(instructions)
	c.run()
	return c.accumulator
}

func solvePart2(instructions []instruction) (acc int) {
	for idx := range instructions {
		instructions[idx].swap()
		c := newConsole(instructions)
		c.run()
		if c.completed {
			return c.accumulator
		}
		instructions[idx].swap() // Swap back to original
	}
	panic("[logical error] unreachable zone")
}

func parse(input string) (instructions []instruction) {
	for _, line := range strings.Split(input, "\n") {
		instructions = append(instructions, newInstruction(line))
	}

	return instructions
}
