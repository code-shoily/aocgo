// Package day02 - Solution for Advent of Code 2024/02
// Problem Link: http://adventofcode.com/2024/day/02
package day02

import (
	_ "embed"
	"fmt"
	"strings"

	"github.com/code-shoily/aocgo/io"
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

func solvePart1(data [][]int) (total int) {
	for _, reports := range data {
		if isSafe(reports) {
			total++
		}
	}

	return total
}

func solvePart2(data [][]int) (total int) {
	for _, reports := range data {
		if isSafe(reports) {
			total++
		} else if isSafeAfterRemoval(reports) {
			total++
		}
	}

	return total
}

func isSafeAfterRemoval(reports []int) bool {
	for i := range reports {

		modified := make([]int, 0, len(reports)-1)
		modified = append(modified, reports[:i]...)
		modified = append(modified, reports[i+1:]...)

		if isSafe(modified) {
			return true
		}
	}

	return false
}

func isSafe(reports []int) bool {
	if len(reports) < 2 {
		return true // Single element or empty is considered safe
	}

	firstDiff := reports[1] - reports[0]
	if !isValidDifference(firstDiff) {
		return false
	}

	isIncreasing := firstDiff > 0

	for i := 2; i < len(reports); i++ {
		diff := reports[i] - reports[i-1]

		if !isValidDifference(diff) || !isConsistentDirection(diff, isIncreasing) {
			return false
		}
	}

	return true
}

func isValidDifference(diff int) bool {
	absDiff := diff
	if absDiff < 0 {
		absDiff = -absDiff
	}
	return absDiff >= 1 && absDiff <= 3
}

func isConsistentDirection(diff int, shouldBeIncreasing bool) bool {
	return (diff > 0) == shouldBeIncreasing
}

func parse(input string) (data [][]int) {
	for line := range strings.SplitSeq(input, "\n") {
		data = append(data, io.SplitByInts(line, " "))
	}

	return data
}
