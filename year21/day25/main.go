// Package day25 - Solution for Advent of Code 2021/25
// Problem Link: http://adventofcode.com/2021/day/25
package day25

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

func solve(input string) int {
	data, steps := parse(input), 1

	for {
		if step(data) == 0 {
			return steps
		}
		steps++
	}
}

func parse(input string) (grid [][]string) {
	for _, line := range strings.Split(input, "\n") {
		grid = append(grid, strings.Split(line, ""))
	}
	return grid
}

func step(grid [][]string) int {
	movements := map[[2]int]string{}

	for i := range grid {
		for j := range grid[i] {
			if jNext := nextCol(grid, j); grid[i][j] == ">" && grid[i][jNext] == "." {
				movements[pos(i, j)] = "."
				movements[pos(i, jNext)] = ">"
			}
		}
	}

	mapToGrid(grid, movements)

	for i := range grid {
		for j := range grid[i] {
			if iNext := nextRow(grid, i); grid[i][j] == "v" && grid[iNext][j] == "." {
				movements[pos(i, j)] = "."
				movements[pos(iNext, j)] = "v"
			}
		}
	}

	mapToGrid(grid, movements)

	return len(movements)
}

func nextRow(grid [][]string, i int) int { return (i + 1) % len(grid) }
func nextCol(grid [][]string, j int) int { return (j + 1) % len(grid[0]) }
func pos(i, j int) [2]int                { return [2]int{i, j} }
func mapToGrid(grid [][]string, movements map[[2]int]string) {
	for position, cucumber := range movements {
		i, j := position[0], position[1]
		grid[i][j] = cucumber
	}
}
