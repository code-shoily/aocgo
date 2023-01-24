// Package day09 - Solution for Advent of Code 2020/09
// Problem Link: http://adventofcode.com/2020/day/09
package day09

import (
	_ "embed"
	"fmt"
	"github.com/code-shoily/aocgo/algo"
	"github.com/code-shoily/aocgo/seq"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

// Run prints out the result of the solution.
func Run() {
	fmt.Println(solve(input))
}

func solve(input string) (int64, int64) {
	data := parse(input)
	invalid := solvePart1(data)
	return invalid, solvePart2(data, invalid)
}

func solvePart1(data []int64) int64 {
	left, right := 0, 25
	for right < len(data) {
		if _, _, found := algo.TwoSum(data[left:right], data[right]); !found {
			return data[right]
		}
		left++
		right++
	}
	return -1
}

func solvePart2(data []int64, invalid int64) int64 {
	left, right := spliceSliceWhen(data, invalid)

	if a, b, found := algo.SubArraySum(left, invalid); found {
		return minMaxSum(data, a, b)
	}

	if a, b, found := algo.SubArraySum(right, invalid); found {
		return minMaxSum(data, a, b)
	}

	panic("[logical error] unreachable zone")
}

func parse(input string) (data []int64) {
	for _, value := range strings.Split(input, "\n") {
		if intValue, err := strconv.ParseInt(value, 10, 64); err == nil {
			data = append(data, intValue)
		}
	}
	return data
}

func spliceSliceWhen(data []int64, value int64) ([]int64, []int64) {
	var idx int
	for data[idx] != value {
		idx++
	}
	return data[:idx], data[idx+1:]
}

func minMaxSum(data []int64, a, b int) int64 {
	min, max := seq.GetMinMax(data[a : b+1])
	return min + max
}
