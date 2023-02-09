// Package day25 - Solution for Advent of Code 2015/25
// Problem Link: http://adventofcode.com/2015/day/25
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
	col, row := parse(input)
	iterations := indexOf(col, row)
	val := 20_151_125
	for i := 1; i < iterations; i++ {
		val = nextValue(val)
	}
	return val
}

func parse(input string) (col, row int) {
	part := strings.TrimSpace(strings.Split(input, ".")[1])
	fmt.Sscanf(part, "Enter the code at row %d, column %d", &row, &col)
	return
}

func indexOf(col, row int) (index int) {
	for i := 0; i <= col; i++ {
		index += i
	}

	for i := 0; i < row-1; i++ {
		index += col + i
	}

	return
}

func nextValue(previousValue int) int {
	return (previousValue * 252_533) % 33_554_393
}
