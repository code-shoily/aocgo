// Package day06 - Solution for Advent of Code 2017/06
// Problem Link: http://adventofcode.com/2017/day/06
package day06

import (
	_ "embed"
	"fmt"
	"github.com/code-shoily/aocgo/io"
)

//go:embed input.txt
var input string

// Because I wanted to test smaller dataset without bleeding with find/replace
const size = 16

// Run prints out the result of the solution.
func Run() {
	fmt.Println(solve(input))
}

func solve(input string) (int, int) {
	data := parse(input)
	visits := map[[size]int]int{}

	var step int
	for {
		step++
		idx, maxValue := getMax(data)
		data[idx] = 0
		for i := (idx + 1) % size; maxValue > 0; i = (i + 1) % size {
			data[i]++
			maxValue--
		}

		if val, ok := visits[data]; ok {
			return step, step - val
		}

		visits[data] = step
	}
}

func parse(input string) (data [size]int) {
	for idx, val := range io.SplitByInts(input, "\t") {
		data[idx] = val
	}

	return data
}

func getMax(data [size]int) (index int, maxValue int) {
	for i, value := range data {
		if value > maxValue {
			maxValue = value
			index = i
		}
	}

	return index, maxValue
}
