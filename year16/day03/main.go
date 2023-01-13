// Package day03 - Solution for Advent of Code 2016/03
// Problem Link: http://adventofcode.com/2016/day/03
package day03

import (
	_ "embed"
	"fmt"
	"github.com/code-shoily/aocgo/seq"
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
	transposed := seq.Transpose(data)

	for _, line := range transposed {
		for _, triples := range seq.ChunkBy(line, 3) {
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
