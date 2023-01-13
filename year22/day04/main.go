// Package day04 - Solution for Advent of Code 2022/04
// Problem Link: http://adventofcode.com/2022/day/04
package day04

import (
	_ "embed"
	"fmt"
	"github.com/code-shoily/aocgo/utils"
)

//go:embed input.txt
var input string

// Run prints out the result of the solution.
func Run() {
	fmt.Println(solve(input))
}

func solve(input string) (int, int) {
	sections := parse(input)
	return solvePart1(sections), solvePart2(sections)
}

func solvePart1(sections [][4]int) (total int) {
	for _, ranges := range sections {
		if hasContained(ranges) {
			total++
		}
	}

	return total
}

func solvePart2(sections [][4]int) (total int) {
	for _, ranges := range sections {
		if hasOverlap(ranges) {
			total++
		}
	}

	return total
}

func parse(input string) (data [][4]int) {
	for _, line := range utils.SplitLines(input) {
		data = append(data, parseSections(line))
	}

	return data
}

func parseSections(line string) [4]int {
	var from1, to1, from2, to2 int
	fmt.Sscanf(line, "%d-%d,%d-%d", &from1, &to1, &from2, &to2)

	return [4]int{from1, to1, from2, to2}
}

func hasContained(ranges [4]int) bool {
	x1 := ranges[0]
	y1 := ranges[1]
	x2 := ranges[2]
	y2 := ranges[3]

	return x1 <= x2 && y1 >= y2 || x2 <= x1 && y2 >= y1
}

func hasOverlap(ranges [4]int) bool {
	x1 := ranges[0]
	y1 := ranges[1]
	x2 := ranges[2]
	y2 := ranges[3]

	return x1 <= x2 && y1 >= x2 || x2 <= x1 && y2 >= x1
}
