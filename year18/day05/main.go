// Package day05 - Solution for Advent of Code 2018/05
// Problem Link: http://adventofcode.com/2018/day/05
package day05

import (
	_ "embed"
	"fmt"
	"github.com/code-shoily/aocgo/algo"
	"github.com/code-shoily/aocgo/utils"
	"math"
	"strings"
)

//go:embed input.txt
var input string

// Run prints out the result of the solution.
func Run() {
	fmt.Println(solve(input))
}

func solve(input string) (int, int) {
	polymer := parse(input)
	return solvePart1(polymer), solvePart2(polymer)
}

func solvePart1(polymer []string) int {
	return react(polymer)
}

func react(polymer []string) int {
	var reacted algo.Stack[string]

	for len(polymer) > 0 {
		currentValue := polymer[0]
		if value, error := reacted.Peek(); !error {
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

func solvePart2(polymer []string) int {
	desiredPolymerLength := math.MaxInt
	for i := 'A'; i <= 'Z'; i++ {
		samplePolymer := removeUnit(polymer, string(i), string(i+('a'-'A')))
		if polymerLength := react(samplePolymer); polymerLength < desiredPolymerLength {
			desiredPolymerLength = polymerLength
		}
	}
	return desiredPolymerLength
}

func removeUnit(polymer []string, a string, b string) (newPolymer []string) {
	for i := 0; i < len(polymer); i++ {
		if polymer[i] == a || polymer[i] == b {
			continue
		}
		newPolymer = append(newPolymer, polymer[i])
	}

	return newPolymer
}

func parse(input string) (data []string) {
	return utils.SplitBy(input, "")
}

func areEqual(value string, currentValue string) bool {
	return value != currentValue && strings.ToLower(value) == strings.ToLower(currentValue)
}
