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
	sea := parse(input)

	for step := 1; ; step++ {
		if move(sea) == 0 {
			return step
		}
	}
}

func parse(input string) (sea [][]string) {
	for _, line := range strings.Split(input, "\n") {
		sea = append(sea, strings.Split(line, ""))
	}
	return sea
}

func move(sea [][]string) int {
	nextPositions := map[[2]int]string{}

	for i := range sea {
		for j := range sea[i] {
			if jNext := nextCol(sea, j); sea[i][j] == ">" && sea[i][jNext] == "." {
				nextPositions[[2]int{i, j}] = "."
				nextPositions[[2]int{i, jNext}] = ">"
			}
		}
	}

	updateSeaMap(sea, nextPositions)

	for i := range sea {
		for j := range sea[i] {
			if iNext := nextRow(sea, i); sea[i][j] == "v" && sea[iNext][j] == "." {
				nextPositions[[2]int{i, j}] = "."
				nextPositions[[2]int{iNext, j}] = "v"
			}
		}
	}

	updateSeaMap(sea, nextPositions)

	return len(nextPositions)
}

func nextRow(sea [][]string, i int) int { return (i + 1) % len(sea) }
func nextCol(sea [][]string, j int) int { return (j + 1) % len(sea[0]) }

func updateSeaMap(sea [][]string, movements map[[2]int]string) {
	for position, cucumber := range movements {
		i, j := position[0], position[1]
		sea[i][j] = cucumber
	}
}
