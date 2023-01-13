// Package day05 - Solution for Advent of Code 2015/05
// Problem Link: http://adventofcode.com/2015/day/05
package day05

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

func solvePart1(text [][]string) (output int) {
	for _, letters := range text {
		if isNiceV1(letters) {
			output++
		}
	}

	return output
}

func solvePart2(text [][]string) (output int) {
	for _, letters := range text {
		if isNiceV2(letters) {
			output++
		}
	}

	return output
}

func parse(input string) [][]string {
	lines := utils.SplitLines(input)
	output := make([][]string, 0, len(lines))
	for _, line := range lines {
		output = append(output, utils.SplitBy(line, ""))
	}

	return output
}

func isNiceV1(letters []string) bool {
	pairs := inPairs(letters)

	if hasDisallowedPair(pairs) {
		return false
	}

	if vowelCount(letters) < 3 {
		return false
	}

	for _, pair := range pairs {
		if pair[0] == pair[1] {
			return true
		}
	}

	return false
}

func inPairs(letters []string) [][2]string {
	output := make([][2]string, 0, len(letters)-1)
	for i := 0; i < len(letters)-1; i++ {
		var pair [2]string
		pair[0] = letters[i]
		pair[1] = letters[i+1]

		output = append(output, pair)
	}

	return output
}

func inTriples(letters []string) [][3]string {
	output := make([][3]string, 0, len(letters)-1)
	for i := 0; i < len(letters)-2; i++ {
		var triple [3]string
		triple[0] = letters[i]
		triple[1] = letters[i+1]
		triple[2] = letters[i+2]

		output = append(output, triple)
	}

	return output
}

func vowelCount(letters []string) (count int) {
	vowels := map[string]bool{
		"a": true,
		"e": true,
		"i": true,
		"o": true,
		"u": true,
	}
	for _, letter := range letters {
		if _, ok := vowels[letter]; ok {
			count++
		}
	}

	return count
}

func hasDisallowedPair(letters [][2]string) bool {
	disallowedPairs := map[[2]string]bool{
		[2]string{"a", "b"}: true,
		[2]string{"c", "d"}: true,
		[2]string{"p", "q"}: true,
		[2]string{"x", "y"}: true,
	}

	for _, letter := range letters {
		if _, ok := disallowedPairs[letter]; ok {
			return true
		}
	}
	return false
}

func isNiceV2(letters []string) bool {
	return hasNonOverlappingPair(letters) && hasSandwichedLetters(letters)
}

func hasNonOverlappingPair(letters []string) bool {
	var pairIndex = make(map[[2]string]int)
	for idx, pairs := range inPairs(letters) {
		if lastIdx, ok := pairIndex[pairs]; ok {
			if idx-lastIdx != 1 {
				return true
			} else {
				continue
			}
		}
		pairIndex[pairs] = idx
	}

	return false
}

func hasSandwichedLetters(letters []string) bool {
	for _, triples := range inTriples(letters) {
		if triples[0] == triples[2] {
			return true
		}
	}

	return false
}
