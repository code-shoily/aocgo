// Package day03 - Solution for Advent of Code 2015/03
// Problem Link: http://adventofcode.com/2015/day/03
package day03

import (
	_ "embed"
	"fmt"
	"strings"
)

//go:embed input.txt
var input string

type position = [2]int

var firstHouse = [2]int{0, 0}

// Run prints out the result of the solution.
func Run() {
	fmt.Println(solve(input))
}

func solve(input string) (int, int) {
	directions := strings.Split(input, "")
	return solvePart1(directions), solvePart2(directions)
}

func solvePart1(directions []string) int {
	var houses = make(map[position]int)
	currentHouse := firstHouse

	for _, direction := range directions {
		houses[currentHouse]++
		updateHouse(&currentHouse, direction)
	}

	return len(houses)
}

func solvePart2(directions []string) int {
	var houses = make(map[position]int)
	currentHouse := firstHouse

	for i := 0; i < len(directions); i += 2 {
		houses[currentHouse]++
		updateHouse(&currentHouse, directions[i])
	}

	currentHouse = firstHouse

	for i := 1; i < len(directions); i += 2 {
		houses[currentHouse]++
		updateHouse(&currentHouse, directions[i])
	}

	return len(houses)
}

func updateHouse(currentHouse *[2]int, direction string) {
	visitUnit := map[string]position{
		">": {1, 0},
		"<": {-1, 0},
		"^": {0, 1},
		"v": {0, -1},
	}[direction]

	currentHouse[0] += visitUnit[0]
	currentHouse[1] += visitUnit[1]
}
