// Package day04 - Solution for Advent of Code 2015/04
// Problem Link: http://adventofcode.com/2015/day/04
package day04

import (
	_ "embed"
	"fmt"
	"github.com/code-shoily/aocgo/algo"
	"strings"
)

var input string = "bgvyzdsv"

// Run prints out the result of the solution.
func Run() {
	fmt.Println(solve(input))
}

func solve(input string) (int, int) {
	return outputForPrefix(input, "00000"), outputForPrefix(input, "000000")
}

func outputForPrefix(data string, prefix string) int {
	var i int
	for {
		text := fmt.Sprintf("%s%d", data, i)
		if hash := algo.GetMD5Hash(text); strings.HasPrefix(hash, prefix) {
			return i
		}
		i++
	}
}
