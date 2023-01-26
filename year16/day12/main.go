// Package day12 - Solution for Advent of Code 2016/12
// Problem Link: http://adventofcode.com/2016/day/12
package day12

import (
	_ "embed"
	"fmt"
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
	data := parse(input)
	return solvePart1(data), solvePart2(data)
}

func solvePart1(data []*instruction) int {
	c := newComputer(data, map[string]int{})
	c.run()
	return c.registers["a"]
}

func solvePart2(data []*instruction) int {
	c := newComputer(data, map[string]int{"c": 1})
	c.run()
	return c.registers["a"]
}

func parse(input string) (instructions []*instruction) {
	for _, line := range strings.Split(input, "\n") {
		instructions = append(instructions, parseInstruction(line))
	}

	return instructions
}

type instruction struct {
	cmd         string
	registerMap []bool
	registers   []string
	values      []int
}

func parseInstruction(line string) *instruction {
	tokens := strings.Split(line, " ")
	var (
		registerMap []bool
		registers   []string
		values      []int
	)
	for _, token := range tokens[1:] {
		value, isRegister := getValue(token)
		registerMap = append(registerMap, isRegister) // TODO: This is registerMap

		if isRegister {
			registers = append(registers, token)
		} else {
			values = append(values, value)
		}
	}

	return &instruction{tokens[0], registerMap, registers, values}

}

func getValue(token string) (value int, isRegister bool) {
	if value, err := strconv.Atoi(token); err == nil {
		return value, false
	}
	return -1, true
}

type computer struct {
	registers    map[string]int
	ptr          int
	instructions []*instruction
}

func (c *computer) run() {
	for c.ptr < len(c.instructions) {
		i := c.instructions[c.ptr]
		switch i.cmd {
		case "inc":
			c.ptr += c.incrementBy(i.registers[0], 1)
		case "dec":
			c.ptr += c.incrementBy(i.registers[0], -1)
		case "cpy":
			c.ptr += c.copy(i)
		case "jnz":
			c.ptr += c.jnz(i)
		}
	}
}

func (c *computer) incrementBy(register string, value int) int {
	c.registers[register] += value
	return 1
}

func (c *computer) copy(i *instruction) int {
	if i.registerMap[0] {
		copyValue := c.registers[i.registers[0]]
		c.registers[i.registers[1]] = copyValue
	} else {
		copyValue := i.values[0]
		c.registers[i.registers[0]] = copyValue
	}
	return 1
}

func (c *computer) jnz(i *instruction) int {
	switch a, b := i.registerMap[0], i.registerMap[1]; {
	case a && b:
		if c.registers[i.registers[0]] != 0 {
			return c.registers[i.registers[1]]
		}
	case a:
		if c.registers[i.registers[0]] != 0 {
			return i.values[0]
		}
	case b:
		if i.values[0] != 0 {
			return c.registers[i.registers[0]]
		}

	default:
		if i.values[0] != 0 {
			return i.values[1]
		}
	}

	return 1
}

func newComputer(instructions []*instruction, bootstrap map[string]int) *computer {
	return &computer{bootstrap, 0, instructions}
}
