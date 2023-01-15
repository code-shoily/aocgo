// Package day04 - Solution for Advent of Code 2017/04
// Problem Link: http://adventofcode.com/2017/day/04
package day04

import (
	_ "embed"
	"fmt"
	"sort"
	"strings"
)

//go:embed input.txt
var input string

// Run prints out the result of the solution.
func Run() {
	fmt.Println(solve(input))
}

func solve(input string) (int, int) {
	phrases := parse(input)
	return solvePart1(phrases), solvePart2(phrases)
}

func solvePart1(phrases [][]string) (total int) {
	for _, phrase := range phrases {
		if hasNoDuplicates(phrase) {
			total++
		}
	}

	return total
}

func solvePart2(phrases [][]string) (total int) {
	for _, phrase := range phrases {
		if hasNoAnagrams(phrase) {
			total++
		}
	}

	return total
}

func parse(input string) (phrases [][]string) {
	for _, line := range strings.Split(input, "\n") {
		phrases = append(phrases, strings.Split(line, " "))
	}

	return phrases
}

func hasNoDuplicates(phrases []string) bool {
	frequencies := map[string]int{}

	for _, phrase := range phrases {
		if frequencies[phrase]++; frequencies[phrase] > 1 {
			return false
		}
	}

	return true
}

func hasNoAnagrams(phrases []string) bool {
	sortedPhrases := make([]string, 0, len(phrases))
	frequencies := map[string]int{}

	for _, phrase := range phrases {
		sortedPhrase := sortPhrase(phrase)
		sortedPhrases = append(sortedPhrases, sortedPhrase)
		if frequencies[sortedPhrase]++; frequencies[sortedPhrase] > 1 {
			return false
		}
	}

	return true
}

func sortPhrase(phrase string) string {
	chars := []byte(phrase)
	sort.Slice(chars, func(i, j int) bool {
		return chars[i] < chars[j]
	})

	return string(chars)
}
