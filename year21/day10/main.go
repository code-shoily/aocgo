// Package day10 - Solution for Advent of Code 2021/10
// Problem Link: http://adventofcode.com/2021/day/10
package day10

import (
	_ "embed"
	"fmt"
	"github.com/code-shoily/aocgo/algo"
	"github.com/code-shoily/aocgo/utils"
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
	lines := utils.SplitLines(input)
	invalidCharCount, invalidSyntax := map[string]int{}, map[string]bool{}

	return solvePart1(lines, invalidCharCount, invalidSyntax), solvePart2(lines, invalidSyntax)
}

func solvePart1(lines []string, invalidCharCount map[string]int, invalidSyntax map[string]bool) int {
	for _, line := range lines {
		if invalidChar, found := getInvalidChar(line); found {
			invalidSyntax[line] = true
			invalidCharCount[invalidChar]++
		}
	}

	return getInvalidScore(invalidCharCount)
}

func solvePart2(lines []string, invalidSyntax map[string]bool) int {
	var results []int
	for _, line := range lines {
		if _, invalid := invalidSyntax[line]; !invalid {
			completions := getCompletion(line)
			results = append(results, getCompletionScore(completions))
		}
	}

	sort.Ints(results)

	return results[len(results)/2]
}

// Fetching Functions
func getCompletion(line string) (completion []string) {
	stack := algo.Stack[string]{}
	for _, char := range strings.Split(line, "") {
		value, empty := stack.Peek()
		if empty {
			stack.Push(char)
		} else if matchingPair(value, char) {
			stack.Pop()
		} else {
			stack.Push(char)
		}
	}

	for !stack.IsEmpty() {
		value, _ := stack.Pop()
		completion = append(completion, getClosingPair(value))
	}

	return completion
}

func getInvalidChar(line string) (c string, found bool) {
	stack := algo.Stack[string]{}
	for _, char := range strings.Split(line, "") {
		value, empty := stack.Peek()
		if empty {
			stack.Push(char)
		} else if matchingPair(value, char) {
			stack.Pop()
		} else if isClosing(char) {
			return char, true
		} else {
			stack.Push(char)
		}
	}
	return c, found
}

// Filtering Functions
func isClosing(char string) bool {
	return char == ")" || char == "}" || char == "]" || char == ">"
}

func getClosingPair(opening string) string {
	switch opening {
	case "(":
		return ")"
	case "{":
		return "}"
	case "[":
		return "]"
	case "<":
		return ">"
	}
	panic("Parsing error.")
}

func matchingPair(opening, closing string) bool {
	return closing == getClosingPair(opening)
}

// Scoring Functions
func getInvalidScore(tally map[string]int) (total int) {
	var invalidPoints = map[string]int{
		")": 3,
		"]": 57,
		"}": 1197,
		">": 25137,
	}

	for char, freq := range tally {
		if point, found := invalidPoints[char]; found {
			total += point * freq
		}
	}

	return total
}

func getCompletionScore(completion []string) (total int) {
	var incompletePoints = map[string]int{
		")": 1,
		"]": 2,
		"}": 3,
		">": 4,
	}
	for _, c := range completion {
		total = total*5 + incompletePoints[c]
	}

	return total
}
