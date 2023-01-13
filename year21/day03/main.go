// Package day03 - Solution for Advent of Code 2021/03
// Problem Link: http://adventofcode.com/2021/day/03
package day03

import (
	_ "embed"
	"fmt"
	"github.com/code-shoily/aocgo/utils"
	"math"
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

func solvePart1(data [][]int) int {
	return gammaRate(data) * epsilonRate(data)
}

func solvePart2(data [][]int) int {
	return co2Rating(data) * o2Rating(data)
}

func parse(input string) (data [][]int) {
	for _, line := range utils.SplitLines(input) {
		data = append(data, utils.SplitByInts(line, ""))
	}

	return data
}

func gammaRate(data [][]int) int {
	columns, rows := len(data[0]), len(data)
	bits := make([]int, columns)

	for col := 0; col < columns; col++ {
		var ones, zeroes int
		for row := 0; row < rows; row++ {
			if data[row][col] == 1 {
				ones++
			} else {
				zeroes++
			}
		}
		if ones > zeroes {
			bits[col] = 1
		}
	}

	return toDecimal(bits)
}

func epsilonRate(data [][]int) int {
	// Since the input was 12 bit per row, hard-coded this
	// had it been other value (data[0]), would've been
	// int(math.Pow(2, float64(len(data[0])))-1)
	return 0x0fff & ^gammaRate(data)
}

func o2Rating(data [][]int) int {
	return gasRating(data, mostFrequentAt)
}

func co2Rating(data [][]int) int {
	return gasRating(data, leastFrequentAt)
}

func filterByBitAt(data [][]int, bit int, at int) (output [][]int) {
	for i := 0; i < len(data); i++ {
		if data[i][at] == bit {
			output = append(output, data[i])
		}
	}

	return output
}

func mostFrequentAt(data [][]int, at int) int {
	var zero, one int
	for i := 0; i < len(data); i++ {
		if data[i][at] == 1 {
			one++
		} else {
			zero++
		}
	}

	if one >= zero {
		return 1
	}
	return 0
}

func leastFrequentAt(data [][]int, at int) int {
	if bit := mostFrequentAt(data, at); bit == 1 {
		return 0
	}
	return 1
}

func toDecimal(bits []int) (decimal int) {
	size := len(bits)
	for i := 0; i < len(bits); i++ {
		decimal += bits[i] * int(math.Pow(2, float64(size-i-1)))
	}
	return decimal
}

type bitFrequencyGetter = func([][]int, int) int

func gasRating(data [][]int, frequencyFn bitFrequencyGetter) int {
	var col int
	for {
		bit := frequencyFn(data, col)
		data = filterByBitAt(data, bit, col)

		if len(data) == 1 {
			return toDecimal(data[0])
		}

		col++
	}

	panic("Shouldn't be here")
}
