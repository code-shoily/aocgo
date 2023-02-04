// Package day16 - Solution for Advent of Code 2016/16
// Problem Link: http://adventofcode.com/2016/day/16
package day16

import (
	_ "embed"
	"fmt"
	"github.com/code-shoily/aocgo/algo"
	"github.com/code-shoily/aocgo/seq"
)

//go:embed input.txt
var input []byte

// Run prints out the result of the solution.
func Run() {
	fmt.Println(solve(input))
}

const (
	zero          = 48
	one           = 49
	part1Capacity = 272
	part2Capacity = 35_651_584
)

func solve(input []byte) (string, string) {
	solver := func(capacity int) string {
		return string(computeChecksum(fill(input, capacity)))
	}
	return solver(part1Capacity), solver(part2Capacity)
}

func fill(initialData []byte, capacity int) []byte {
	a := initialData

	for len(a) <= capacity {
		b := make([]byte, len(a))
		copy(b, a)
		a = append(
			append(a, zero),
			invertFlip(b)...,
		)
	}

	return a[:capacity]
}

func invertFlip(b []byte) []byte {
	out := make([]byte, 0, len(b))
	for i := len(b) - 1; i >= 0; i-- {
		switch b[i] {
		case zero:
			out = append(out, one)
		case one:
			out = append(out, zero)
		}
	}
	return out
}

func computeChecksum(state []byte) []byte {
	for algo.IsEven(len(state)) {
		var temp []byte
		chunks := seq.Chunk(state, 2, 2, true)
		for _, pair := range chunks {
			if pair[0] == pair[1] {
				temp = append(temp, one)
			} else {
				temp = append(temp, zero)
			}
		}
		state = temp
	}

	return state
}
