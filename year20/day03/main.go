// Package day03 - Solution for Advent of Code 2020/03
// Problem Link: http://adventofcode.com/2020/day/03
package day03

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
	coords := parse(input)
	return solvePart1(coords), solvePart2(coords)
}

func solvePart1(coords [][]string) (trees int) {
	return countTrees(coords, 3, 1)
}

func solvePart2(coords [][]string) int {
	slopes := [][]int{
		{1, 1},
		{3, 1},
		{5, 1},
		{7, 1},
		{1, 2},
	}

	result := 1

	for _, rightDown := range slopes {
		result *= countTrees(coords, rightDown[0], rightDown[1])
	}

	return result
}

func countTrees(coords [][]string, right, down int) (trees int) {
	horizontal, vertical := 0, 0

	for {
		if coords[horizontal][vertical] == "#" {
			trees++
		}

		horizontal += down
		vertical = (vertical + right) % len(coords[0])

		if horizontal >= len(coords) {
			break
		}
	}

	return trees
}

func parse(input string) (coords [][]string) {
	for _, lines := range strings.Split(input, "\n") {
		coords = append(coords, strings.Split(lines, ""))
	}

	return coords
}
