// Package day06 - Solution for Advent of Code 2021/06
// Problem Link: http://adventofcode.com/2021/day/06
package day06

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
	timers := parse(input)
	return multiply(timers, 80), multiply(timers, 256)
}

func parse(input string) []int {
	frequency := make([]int, 9)
	for _, t := range utils.SplitByInts(input, ",") {
		frequency[t]++
	}

	return frequency
}

func multiply(timers []int, days int) (lanternFish int) {
	for day := 0; day < days; day++ {
		newTimers := make([]int, len(timers))
		for i := 0; i < 9; i++ {
			timer := timers[i]
			if i == 0 {
				newTimers[8] = timer
				newTimers[6] = timer
			} else {
				newTimers[i-1] += timer
			}
		}
		timers = newTimers
	}

	for _, timer := range timers {
		lanternFish += timer
	}

	return lanternFish
}
