// Package day25 - Solution for Advent of Code 2016/25
// Problem Link: http://adventofcode.com/2016/day/25
package day25

import (
	_ "embed"
	"fmt"
	"github.com/code-shoily/aocgo/algo"
	"regexp"
	"strconv"
)

//go:embed input.txt
var input string

// Run prints out the result of the solution.
func Run() {
	fmt.Println(solve())
}

func solve() int {
	for aValue := 0; ; aValue++ {
		if emitsSignal(aValue) {
			return aValue
		}
	}
}

func emitsSignal(registerA int) bool {
	var outputs string

	a := registerA
	d := a + 2572
	b := 0

	pattern := regexp.MustCompile("^(01){1,}$")

	for {
		a = d
		for a != 0 {
			a, b = a/2, a%2
			outputs += strconv.Itoa(b)
			if size := len(outputs); algo.IsEven(size) {
				if !pattern.MatchString(outputs) {
					return false
				} else if size > 50 {
					return true
				}
			}
		}
	}
}
