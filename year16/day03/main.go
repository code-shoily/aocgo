// Package day03 - Solution for Advent of Code 2016/03
// Problem Link: http://adventofcode.com/2016/day/03
package day03

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

func solvePart1(data [][]int) (triangles int) {
	for _, triples := range data {
		if a, b, c := triples[0], triples[1], triples[2]; isTriangle(a, b, c) {
			triangles++
		}
	}

	return triangles
}

func solvePart2(data [][]int) (triangles int) {
	transposed := transpose(data)

	for _, line := range transposed {
		for _, triples := range chunkBy(line, 3) {
			if a, b, c := triples[0], triples[1], triples[2]; isTriangle(a, b, c) {
				triangles++
			}
		}
	}

	return triangles
}

func parse(input string) (data [][]int) {
	for _, value := range utils.SplitLines(input) {
		data = append(data, utils.SplitByInts(value, " "))
	}

	return data
}

func isTriangle(a, b, c int) bool {
	return a+b > c && b+c > a && c+a > b
}

func transpose(data [][]int) [][]int {
	// FIXME: Move this into its own module and make it generic
	transposed := make([][]int, len(data[0]))

	for idx := range transposed {
		transposed[idx] = make([]int, len(data))
	}

	for i := 0; i < len(data); i++ {
		for j := 0; j < len(data[0]); j++ {
			transposed[j][i] = data[i][j]
		}
	}

	return transposed
}

func chunkBy(seq []int, by int) (chunks [][]int) {
	// This only takes care of slices with lengths divisible by the chunk size
	// which satisfies this particular problem, a more general approach should
	// be taken and this function should be moved out.
	// FIXME: Tackle all edge cases and move it as a generic function in its own module
	for i := 0; i < len(seq); i += by {
		chunk := make([]int, by)
		for j := 0; j < by; j++ {
			chunk[j] = seq[i+j]
		}
		chunks = append(chunks, chunk)
	}

	return chunks
}
