package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

func main() {
	fmt.Println(solve(input))
}

func solve(input string) (int, int) {
	instructions := parse(input)
	return solvePart1(instructions), solvePart2(instructions)
}

func solvePart1(instructions []Instruction) int {
	facing := North
	location := Point{0, 0}

	for _, instruction := range instructions {
		newFacing := facing.navigate(instruction.direction)
		moveBy := FaceFactors[newFacing]

		location = location.add(moveBy.multiply(instruction.steps))

		facing = newFacing
	}

	return location.manhattanDistance(Point{0, 0})
}

func solvePart2(instructions []Instruction) int {
	facing := North
	location := Point{0, 0}
	visited := make(map[Point]bool)

	for _, instruction := range instructions {
		newFacing := facing.navigate(instruction.direction)
		moveBy := FaceFactors[newFacing]

		for i := 0; i < instruction.steps; i++ {
			location = location.add(moveBy)
			if _, ok := visited[location]; ok {
				return location.manhattanDistance(Point{0, 0})
			}
			visited[location] = true
		}

		facing = newFacing
	}

	panic("No points visited twice")
}

type Instruction struct {
	direction string
	steps     int
}

func parse(input string) []Instruction {
	var locations []Instruction
	for _, location := range strings.Split(input, ", ") {
		locations = append(locations, parseLocation(location))
	}

	return locations
}

func parseLocation(location string) Instruction {
	parsedLine := strings.SplitN(location, "", 2)
	if steps, err := strconv.Atoi(parsedLine[1]); err == nil {
		return Instruction{parsedLine[0], steps}
	}

	panic("Invalid format")
}
