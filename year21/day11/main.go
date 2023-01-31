// Package day11 - Solution for Advent of Code 2021/11
// Problem Link: http://adventofcode.com/2021/day/11
package day11

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

func solve(input string) (flashes int, steps int) {
	data := parse(input)

	for ; steps < 100; steps++ {
		flashes += step(data)
	}

	for ; step(data) != 100; steps++ {
	}

	return flashes, steps + 1
}

func parse(input string) (grid [][]int) {
	for _, row := range strings.Split(input, "\n") {
		grid = append(grid, io.SplitByInts(row, ""))
	}
	return grid
}

func step(grid [][]int) (flashes int) {
	// Increase energy level
	for i := range grid {
		for j := range grid[i] {
			grid[i][j]++
		}
	}

	// Flash
	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] > 9 {
				flashes += flash(grid, i, j)
			}
		}
	}

	return flashes
}

func flash(grid [][]int, i, j int) (total int) {
	var adjacentLevels = [][2]int{
		{i - 1, j},     // north
		{i - 1, j + 1}, // north-east
		{i, j + 1},     // east
		{i + 1, j + 1}, // south-east
		{i + 1, j},     // south
		{i + 1, j - 1}, // south-west
		{i, j - 1},     // west
		{i - 1, j - 1}, // north-west
	}

	withinGrid := func(point [2]int) bool {
		a, b := point[0], point[1]
		return a >= 0 && a < len(grid) && b >= 0 && b < len(grid[a])
	}

	grid[i][j] = 0 // Flash!

	// Flash all eligible adjacent octopuses recursively
	for _, adj := range adjacentLevels {
		if withinGrid(adj) {
			if iAdj, jAdj := adj[0], adj[1]; grid[iAdj][jAdj] == 9 {
				total += flash(grid, iAdj, jAdj)
			} else if grid[iAdj][jAdj] > 0 { // Don't flash twice in a single step
				grid[iAdj][jAdj]++
			}
		}
	}

	return total + 1
}
