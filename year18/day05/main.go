// Package day05 - Solution for Advent of Code 2018/05
// Problem Link: http://adventofcode.com/2018/day/05
package day05

import (
	_ "embed"
	"fmt"
	"github.com/code-shoily/aocgo/algo"
	"math"
)

//go:embed input.txt
var input []byte

const casingOffset = 32 // 'a' - 'A'

// Run prints out the result of the solution.
func Run() {
	fmt.Println(solve(input))
}

func solve(input []byte) (int, int) {
	return solvePart1(input), solvePart2(input)
}

func solvePart1(polymer []byte) int {
	return react(polymer)
}

func solvePart2(polymer []byte) int {
	desiredPolymerLength := math.MaxInt
	for i := 'A'; i <= 'Z'; i++ {
		testPolymer := without(polymer, byte(i))
		if polymerLength := react(testPolymer); polymerLength < desiredPolymerLength {
			desiredPolymerLength = polymerLength
		}
	}
	return desiredPolymerLength
}

func react(polymer []byte) int {
	var reacted algo.Stack[byte]

	for len(polymer) > 0 {
		current := polymer[0]
		if unit, err := reacted.Peek(); !err {
			if willReact(unit, current) {
				reacted.Pop()
			} else {
				reacted.Push(current)
			}
		} else {
			reacted.Push(current)
		}
		polymer = polymer[1:]
	}

	return len(reacted)
}

func without(polymer []byte, unit byte) []byte {
	newPolymer := make([]byte, 0, len(polymer))
	for i := 0; i < len(polymer); i++ {
		if current := polymer[i]; current != unit && current != unit+casingOffset {
			newPolymer = append(newPolymer, current)
		}
	}

	return newPolymer
}

func willReact(unit1 byte, unit2 byte) bool {
	if unit1 > unit2 {
		return unit1-unit2 == casingOffset
	}
	return unit2-unit1 == casingOffset
}
