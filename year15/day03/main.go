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

var startingPosition = [2]int{0, 0}
var movementFor = map[string]position{
	">": {1, 0},
	"<": {-1, 0},
	"^": {0, 1},
	"v": {0, -1},
}

// Run prints out the result of the solution.
func Run() {
	fmt.Println(solve(input))
}

func solve(input string) (int, int) {
	directions := strings.Split(input, "")
	return solvePart1(directions), solvePart2(directions)
}

func solvePart1(directions []string) int {
	var visits = make(map[position]int)
	visited := startingPosition

	for _, dir := range directions {
		visits[visited]++
		movement := movementFor[dir]
		visited = [2]int{
			visited[0] + movement[0],
			visited[1] + movement[1],
		}
	}

	return len(visits)
}

func solvePart2(directions []string) int {
	var visits = make(map[position]int)
	visited := startingPosition

	for i := 0; i < len(directions); i += 2 {
		visits[visited]++
		movement := movementFor[directions[i]]
		visited = [2]int{
			visited[0] + movement[0],
			visited[1] + movement[1],
		}
	}

	visited = startingPosition

	for j := 1; j < len(directions); j += 2 {
		visits[visited]++
		movement := movementFor[directions[j]]
		visited = [2]int{
			visited[0] + movement[0],
			visited[1] + movement[1],
		}
	}

	return len(visits)
}
