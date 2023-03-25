// Package day07 - Solution for Advent of Code 2021/07
// Problem Link: http://adventofcode.com/2021/day/07
package day07

import (
	_ "embed"
	"fmt"
	"github.com/code-shoily/aocgo/algo"
	"math"
	"strconv"
	"strings"
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

func solvePart1(positions []int) (fuelCost int) {
	median := algo.Median(positions)

	for _, pos := range positions {
		fuelCost += int(math.Abs(float64(pos) - median))
	}

	return fuelCost
}

func solvePart2(positions []int) (fuelCost int) {
	mean := math.Round(algo.Mean(positions))

	for _, pos := range positions {
		diff := math.Abs(float64(pos)-mean) + 1
		fuelCost += int(diff * (diff - 1) / 2)
	}

	return fuelCost
}

func parse(input string) (positions []int) {
	for _, token := range strings.Split(input, ",") {
		if position, err := strconv.Atoi(token); err == nil {
			positions = append(positions, position)
		}
	}

	return positions
}
