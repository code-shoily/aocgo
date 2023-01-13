// Package day02 - Solution for Advent of Code 2022/02
// Problem Link: http://adventofcode.com/2022/day/02
package day02

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
	data := parse(input)
	return solvePart1(data), solvePart2(data)
}

func solvePart1(data []string) (total int) {
	score := map[string]int{
		"A X": 4, // 1 + 3 (selection + outcome)
		"A Y": 8, // 2 + 6
		"A Z": 3, // 3 + 0
		"B X": 1, // 1 + 0
		"B Y": 5, // 2 + 3
		"B Z": 9, // 3 + 6
		"C X": 7, // 1 + 6
		"C Y": 2, // 2 + 0
		"C Z": 6, // 3 + 3
	}

	for _, choice := range data {
		total += score[choice]
	}

	return total
}

func solvePart2(data []string) (total int) {
	score := map[string]int{
		"A X": 3, // 3 + 0 (selection + outcome)
		"A Y": 4, // 1 + 3
		"A Z": 8, // 2 + 6
		"B X": 1, // 1 + 0
		"B Y": 5, // 2 + 3
		"B Z": 9, // 3 + 6
		"C X": 2, // 2 + 0
		"C Y": 6, // 3 + 3
		"C Z": 7, // 1 + 6
	}

	for _, choice := range data {
		total += score[choice]
	}

	return total
}

func parse(input string) (data []string) {
	return utils.SplitLines(input)
}
