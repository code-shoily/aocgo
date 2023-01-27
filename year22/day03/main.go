// Package day03 - Solution for Advent of Code 2022/03
// Problem Link: http://adventofcode.com/2022/day/03
package day03

import (
	_ "embed"
	"fmt"
	"github.com/code-shoily/aocgo/io"
)

//go:embed input.txt
var input string

// Run prints out the result of the solution.
func Run() {
	fmt.Println(solve(input))
}

func solve(input string) (int, int) {
	rucksacks := parse(input)
	return solvePart1(rucksacks), solvePart2(rucksacks)
}

func solvePart1(rucksacks []string) (total int) {
	for _, rucksack := range rucksacks {
		total += getPriority(inBothCompartments(rucksack))
	}

	return total
}

func solvePart2(rucksacks []string) (total int) {
	for _, rucksackGroup := range groupRucksacks(rucksacks) {
		a, b, c := rucksackGroup[0], rucksackGroup[1], rucksackGroup[2]
		total += getPriority(inAllGroups(a, b, c))
	}

	return total
}

func parse(input string) (data []string) {
	return io.SplitLines(input)
}

func compartments(rucksack string) (string, string) {
	rucksackSize := len(rucksack) / 2
	return rucksack[:rucksackSize], rucksack[rucksackSize:]
}

func inBothCompartments(rucksack string) rune {
	a, b := compartments(rucksack)
	aSet := make(map[rune]bool)

	for _, v := range a {
		aSet[v] = true
	}

	for _, v := range b {
		if _, ok := aSet[v]; ok {
			return v
		}
	}

	panic("No common item found")
}

func getPriority(r rune) int {
	switch {
	case 97 <= r && r <= 122:
		return int(r - 97 + 1)
	case 65 <= r && r <= 90:
		return int(r - 65 + 27)
	default:
		panic("Invalid input provided")
	}
}

func groupRucksacks(rucksacks []string) (groups [][]string) {
	for i := 0; i < len(rucksacks); i += 3 {
		groups = append(groups, rucksacks[i:i+3])
	}

	return groups
}

func inAllGroups(a, b, c string) rune {
	aSet := make(map[rune]bool)
	for _, v := range a {
		aSet[v] = true
	}
	bSet := make(map[rune]bool)
	for _, v := range b {
		if _, ok := aSet[v]; ok {
			bSet[v] = true
		}
	}

	for _, v := range c {
		if _, ok := bSet[v]; ok {
			return v
		}
	}

	panic("Invalid input detected")
}
