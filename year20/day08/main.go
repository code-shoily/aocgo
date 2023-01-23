// Package day08 - Solution for Advent of Code 2020/08
// Problem Link: http://adventofcode.com/2020/day/08
package day08

import (
	_ "embed"
	"fmt"
	"github.com/code-shoily/aocgo/utils"
	"strconv"
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
	c := NewConsole(instructions)
	c.run()
	return c.accumulator
}

func solvePart2(instructions []instruction) (acc int) {
	for idx := range instructions {
		instructions[idx].swap()
		c := NewConsole(instructions)
		c.run()
		if c.isDone {
			return c.accumulator
		}
		instructions[idx].swap()
	}
	panic("[logical error] unreachable zone")
}

func parse(input string) []instruction {
	var instructions []instruction
	for _, line := range utils.SplitLines(input) {
		instructions = append(instructions, parseInstruction(line))
	}

	return instructions
}

type instruction struct {
	cmd   string
	value int
}

func (i *instruction) swap() {
	switch i.cmd {
	case "jmp":
		i.cmd = "nop"
	case "nop":
		i.cmd = "jmp"
	}
}

func parseInstruction(line string) instruction {
	tokens := strings.Split(line, " ")

	if value, err := strconv.Atoi(tokens[1]); err == nil {
		return instruction{tokens[0], value}
	}

	panic("[logical error] unreachable zone")
}

func NewConsole(instructions []instruction) *console {
	return &console{instructions: instructions, visited: make(map[int]bool)}
}

type console struct {
	accumulator  int
	pointer      int
	instructions []instruction
	visited      map[int]bool
	isLoop       bool
	isDone       bool
}

func (c *console) run() {
	for !c.isLoop && !c.isDone {
		c.step()
	}
}

func (c *console) step() {
	if _, loop := c.visited[c.pointer]; loop {
		c.isLoop = true
		return
	}

	if c.pointer < -1 || c.pointer >= len(c.instructions) {
		c.isDone = true
		return
	}

	c.visited[c.pointer] = true
	currentInstruction := c.instructions[c.pointer]

	switch currentInstruction.cmd {
	case "nop":
		c.pointer++
	case "acc":
		c.accumulator += currentInstruction.value
		c.pointer++
	case "jmp":
		c.pointer += currentInstruction.value
	}
}
