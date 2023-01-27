// Package day04 - Solution for Advent of Code 2019/04
// Problem Link: http://adventofcode.com/2019/day/04
package day04

import (
	_ "embed"
	"fmt"
	"github.com/code-shoily/aocgo/algo"
	"github.com/code-shoily/aocgo/io"
)

//go:embed input.txt
var input string

// Run prints out the result of the solution.
func Run() {
	fmt.Println(solve(input))
}

func solve(input string) (int, int) {
	from, to := parse(input)
	return solvePart1(from, to), solvePart2(from, to)
}

func solvePart1(from, to int) (total int) {
	for i := from; i <= to; i++ {
		if hasDuplicate(i) && isIncreasing(i) {
			total++
		}
	}

	return total
}

func solvePart2(from, to int) (total int) {
	for i := from; i <= to; i++ {
		if hasIsolatedDuplicate(i) && isIncreasing(i) {
			total++
		}
	}

	return total
}

func parse(input string) (int, int) {
	limits := io.SplitByInts(input, "-")

	return limits[0], limits[1]
}

func isIncreasing(number int) bool {
	digits := algo.IntToDigits(number)

	for i := 0; i < len(digits)-1; i++ {
		if digits[i] > digits[i+1] {
			return false
		}
	}

	return true
}

func hasDuplicate(number int) bool {
	digits := algo.IntToDigits(number)

	for i := 0; i < len(digits)-1; i++ {
		if digits[i] == digits[i+1] {
			return true
		}
	}

	return false
}

func hasIsolatedDuplicate(number int) bool {
	digits := algo.IntToDigits(number)

	for i := 0; i < len(digits)-1; i++ {
		if digits[i] == digits[i+1] {
			isPartOfLargerDupes :=
				(i > 0 && digits[i-1] == digits[i]) ||
					(i < len(digits)-2 && digits[i+1] == digits[i+2])
			if !isPartOfLargerDupes {
				return true
			}
		}
	}

	return false
}
