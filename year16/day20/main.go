// Package day20 - Solution for Advent of Code 2016/20
// Problem Link: http://adventofcode.com/2016/day/20
package day20

import (
	_ "embed"
	"fmt"
	"sort"
	"strings"
)

//go:embed input.txt
var input string

const (
	MinValue = 0
	MaxValue = 429_496_7295
)

// Run prints out the result of the solution.
func Run() {
	fmt.Println(solve(input))
}

func solve(input string) (int, int) {
	data := parse(input)
	return solvePart1(data), solvePart2(data)
}

func solvePart1(data [][2]int) int {
	mergedPairs := mergeBlockedIPRanges(data)

	if mergedPairs[0][0]-1 >= 0 {
		return mergedPairs[0][0] - 1
	}
	return mergedPairs[0][1] + 1
}

func solvePart2(data [][2]int) (allowedIPs int) {
	mergedPairs := mergeBlockedIPRanges(data)

	if minBlocked := mergedPairs[0][0]; minBlocked > MinValue {
		allowedIPs += mergedPairs[0][0] - MinValue
	}
	if maxBlocked := mergedPairs[len(mergedPairs)-1][1]; maxBlocked < MaxValue {
		allowedIPs += MaxValue - maxBlocked
	}

	for i := 0; i < len(mergedPairs)-1; i++ {
		allowedIPs += mergedPairs[1][0] - mergedPairs[0][1] - 1
	}

	return allowedIPs
}

func parse(input string) (ranges [][2]int) {
	for _, line := range strings.Split(input, "\n") {
		var from, to int
		if _, err := fmt.Sscanf(line, "%d-%d", &from, &to); err != nil {
			panic("Invalid input format")
		}
		ranges = append(ranges, [2]int{from, to})
	}

	sort.SliceStable(ranges, func(i, j int) bool {
		return ranges[i][0] < ranges[j][0]
	})

	return ranges
}

func mergeBlockedIPRanges(ranges [][2]int) (mergedRanges [][2]int) {
	a := ranges[0]
	for _, b := range ranges[1:] {
		if b[0]-1 <= a[1] {
			if b[1] > a[1] {
				a[1] = b[1]
			}
		} else {
			mergedRanges = append(mergedRanges, a)
			a = b
		}
	}

	mergedRanges = append(mergedRanges, a)

	return mergedRanges
}
