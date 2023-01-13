// Package day05 - Solution for Advent of Code 2016/05
// Problem Link: http://adventofcode.com/2016/day/05
package day05

import (
	"fmt"
	"github.com/code-shoily/aocgo/algo"
	"strconv"
	"strings"
)

var input = "cxdnnyjw"

// Run prints out the result of the solution.
func Run() {
	fmt.Println(solve(input))
}

func solve(input string) (string, string) {
	return solvePart1(input), solvePart2(input)
}

func solvePart1(data string) string {
	passwordSlice := make([]string, 8)
	var start int64 = -1
	for i := 0; i < 8; i++ {
		idx, interestingString := outputForPrefix(data, "00000", start+1)
		start = idx
		passwordSlice[i] = interestingString[5:6]
	}

	return strings.Join(passwordSlice, "")
}

func solvePart2(data string) (password string) {
	passwordMap := make(map[int64]string)
	var start int64 = -1
	for i := 0; len(passwordMap) < 8; i++ {
		idx, interestingString := outputForPrefix(data, "00000", start+1)
		start = idx
		idx = toNumber(interestingString[5:6])
		if idx >= 0 && idx < 8 {
			if _, ok := passwordMap[idx]; !ok {
				passwordMap[idx] = interestingString[6:7]
			}
		}
	}

	for i := 0; i < 8; i++ {
		password += passwordMap[int64(i)]
	}

	return password
}

func outputForPrefix(data string, prefix string, from int64) (int64, string) {
	i := from
	for {
		text := fmt.Sprintf("%s%d", data, i)
		if hash := algo.GetMD5Hash(text); strings.HasPrefix(hash, prefix) {
			return i, hash
		}
		i++
	}
}

func toNumber(s string) int64 {
	if n, error := strconv.ParseInt(s, 16, 64); error == nil {
		return n
	}
	panic("Invalid processing")
}
