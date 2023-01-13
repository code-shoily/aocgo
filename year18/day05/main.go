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
		samplePolymer := removeUnit(polymer, byte(i))
		if polymerLength := react(samplePolymer); polymerLength < desiredPolymerLength {
			desiredPolymerLength = polymerLength
		}
	}
	return desiredPolymerLength
}

func react(polymer []byte) int {
	var reacted algo.Stack[byte]

	for len(polymer) > 0 {
		currentValue := polymer[0]
		if value, err := reacted.Peek(); !err {
			if areEqual(value, currentValue) {
				reacted.Pop()
			} else {
				reacted.Push(currentValue)
			}
		} else {
			reacted.Push(currentValue)
		}
		polymer = polymer[1:]
	}

	return len(reacted)
}

func removeUnit(polymer []byte, a byte) []byte {
	newPolymer := make([]byte, 0, len(polymer))
	for i := 0; i < len(polymer); i++ {
		unit := polymer[i]
		if unit == a || unit == a+casingOffset {
			continue
		}
		newPolymer = append(newPolymer, unit)
	}

	return newPolymer
}

func areEqual(value byte, currentValue byte) bool {
	if value > currentValue {
		return value-currentValue == casingOffset
	}
	return currentValue-value == casingOffset
}
