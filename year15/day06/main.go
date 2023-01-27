// Package day06 - Solution for Advent of Code 2015/06
// Problem Link: http://adventofcode.com/2015/day/06
package day06

import (
	_ "embed"
	"fmt"
	"github.com/code-shoily/aocgo/io"
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

func solvePart1(instructions []Instruction) int {
	lights := makeLights()

	for _, instruction := range instructions {
		switch instruction.operation {
		case "turn on":
			setLights(lights, 1, instruction)
		case "turn off":
			setLights(lights, 0, instruction)
		case "toggle":
			toggleLights(lights, instruction)
		}
	}

	return totalLit(lights)
}

func solvePart2(instructions []Instruction) int {
	lights := makeLights()

	for _, instruction := range instructions {
		switch instruction.operation {
		case "turn on":
			updateLights(lights, 1, instruction)
		case "turn off":
			updateLights(lights, -1, instruction)
		case "toggle":
			updateLights(lights, 2, instruction)
		}
	}

	return totalLit(lights)
}

func setLights(lights [][]int, status int, instruction Instruction) {
	for i := instruction.x1; i <= instruction.x2; i++ {
		for j := instruction.y1; j <= instruction.y2; j++ {
			lights[i][j] = status
		}
	}
}

func toggleLights(lights [][]int, instruction Instruction) {
	for i := instruction.x1; i <= instruction.x2; i++ {
		for j := instruction.y1; j <= instruction.y2; j++ {
			var status int
			if lights[i][j] == 0 {
				status = 1
			}
			lights[i][j] = status
		}
	}
}

func updateLights(lights [][]int, by int, instruction Instruction) {
	for i := instruction.x1; i <= instruction.x2; i++ {
		for j := instruction.y1; j <= instruction.y2; j++ {
			if by == -1 && lights[i][j] == 0 {
				lights[i][j] = 0
			} else {
				lights[i][j] += by
			}
		}
	}
}

func makeLights() [][]int {
	lights := make([][]int, 1000)

	for idx := range lights {
		row := make([]int, 1000)
		lights[idx] = row
	}

	return lights
}
func totalLit(lights [][]int) (on int) {
	for _, row := range lights {
		for _, cell := range row {
			on += cell
		}
	}

	return on
}

type Instruction struct {
	operation string
	x1        int
	y1        int
	x2        int
	y2        int
}

func parse(input string) (data []Instruction) {
	for _, line := range io.SplitLines(input) {
		data = append(data, parseInstruction(line))
	}

	return data
}

func parseInstruction(line string) Instruction {
	var x1, y1, x2, y2 int
	var op string

	tokens := io.SplitBy(line, " ")
	if tokens[0] == "toggle" {
		fmt.Sscanf(line[7:], "%d,%d through %d,%d", &x1, &y1, &x2, &y2)
		return Instruction{"toggle", x1, y1, x2, y2}
	}

	op = strings.Join(tokens[:2], " ")
	startIdx := 8

	if op == "turn off" {
		startIdx = 9
	}

	fmt.Sscanf(line[startIdx:], "%d,%d through %d,%d", &x1, &y1, &x2, &y2)

	return Instruction{op, x1, y1, x2, y2}
}
