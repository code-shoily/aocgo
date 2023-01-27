// Package day02 - Solution for Advent of Code 2021/02
// Problem Link: http://adventofcode.com/2021/day/02
package day02

import (
	_ "embed"
	"fmt"
	"github.com/code-shoily/aocgo/io"
	"strconv"
)

//go:embed input.txt
var input string

// Run prints out the result of the solution.
func Run() {
	fmt.Println(solve(input))
}

func solve(input string) (int, int) {
	commands := parse(input)
	return solvePart1(commands), solvePart2(commands)
}

func solvePart1(commands []Command) int {
	position := new(Position)
	for _, command := range commands {
		position.applyCommandV1(command)
	}

	return position.depth * position.horizontal
}

func solvePart2(commands []Command) int {
	position := new(Position)
	for _, command := range commands {
		position.applyCommandV2(command)
	}

	return position.depth * position.horizontal
}

func parse(input string) (data []Command) {
	for _, line := range io.SplitLines(input) {
		data = append(data, newCommand(line))
	}

	return data
}

type Command struct {
	direction string
	steps     int
}

type Position struct {
	horizontal int
	depth      int
	aim        int
}

func (p *Position) applyCommandV1(command Command) {
	switch command.direction {
	case "forward":
		p.horizontal += command.steps
	case "up":
		p.depth -= command.steps
	case "down":
		p.depth += command.steps
	}
}

func (p *Position) applyCommandV2(command Command) {
	switch command.direction {
	case "forward":
		p.horizontal += command.steps
		p.depth += p.aim * command.steps
	case "up":
		p.aim -= command.steps
	case "down":
		p.aim += command.steps
	}
}

func newCommand(line string) Command {
	tokens := io.SplitBy(line, " ")
	steps, err := strconv.Atoi(tokens[1])
	if err != nil {
		panic("Invalid input")
	}
	return Command{tokens[0], steps}
}
