// Package day20 - Solution for Advent of Code 2015/20
// Problem Link: http://adventofcode.com/2015/day/20
package day20

import (
	_ "embed"
	"fmt"
	"github.com/code-shoily/aocgo/algo"
	"strconv"
)

//go:embed input.txt
var input string

// Run prints out the result of the solution.
func Run() {
	fmt.Println(solve(input))
}

func solve(input string) (int, int) {
	limit, _ := strconv.Atoi(input)

	deliverGifts := func(strategy func(int, bool) int, infinite bool) int {
		for houseNumber := 1; ; houseNumber++ {
			if gifts := strategy(houseNumber, infinite); gifts >= limit {
				return houseNumber
			}
		}
	}
	return deliverGifts(giftStrategy, true), deliverGifts(giftStrategy, false)
}

func giftStrategy(number int, infinite bool) (gifts int) {
	elves := algo.Factors(number)

	for _, elf := range elves {
		if infinite {
			gifts += elf * 10
		} else if number/elf <= 50 {
			gifts += elf * 11
		}
	}

	return
}
