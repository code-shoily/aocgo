// Package day08 - Solution for Advent of Code 2017/08
// Problem Link: http://adventofcode.com/2017/day/08
package day08

import (
	_ "embed"
	"fmt"
	"github.com/code-shoily/aocgo/seq"
	"golang.org/x/exp/maps"
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
	instructions := parse(input)
	registers := map[string]int{}
	largestEver := math.MinInt // Part 2

	for _, i := range instructions {
		if i.compareFunc(registers) {
			if updatedValue := i.updateFunc(registers); updatedValue > largestEver {
				largestEver = updatedValue
			}
		}
	}

	_, max := seq.GetMinMax(maps.Values(registers)) // Part 1

	return max, largestEver
}

func parse(input string) (instructions []instruction) {
	for _, line := range strings.Split(input, "\n") {
		var reg, updater, dep, cmp string
		var updateVal, cmpVal int
		fmt.Sscanf(line, "%s %s %d if %s %s %d",
			&reg, &updater, &updateVal, &dep, &cmp, &cmpVal)
		instructions = append(instructions, newInstruction(reg, updater, updateVal, dep, cmp, cmpVal))
	}

	return instructions
}

type instruction struct {
	updateFunc  func(map[string]int) int
	compareFunc func(map[string]int) bool
}

func newInstruction(reg string, updater string, updateVal int, dep string, cmp string, cmpVal int) instruction {
	updateFunc := func(registers map[string]int) int {
		switch updater {
		case "inc":
			registers[reg] += updateVal
		case "dec":
			registers[reg] -= updateVal
		}

		return registers[reg]
	}

	compareFunc := func(registers map[string]int) bool {
		switch cmp {
		case "==":
			return registers[dep] == cmpVal
		case "!=":
			return registers[dep] != cmpVal
		case "<":
			return registers[dep] < cmpVal
		case "<=":
			return registers[dep] <= cmpVal
		case ">":
			return registers[dep] > cmpVal
		case ">=":
			return registers[dep] >= cmpVal
		}
		panic("Logical Error: Unreachable Code")
	}

	return instruction{updateFunc, compareFunc}
}
