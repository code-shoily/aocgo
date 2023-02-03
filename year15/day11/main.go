// Package day11 - Solution for Advent of Code 2015/11
// Problem Link: http://adventofcode.com/2015/day/11
package day11

import (
	_ "embed"
	"fmt"
)

//go:embed input.txt
var input string

// Run prints out the result of the solution.
func Run() {
	fmt.Println(solve(input))
}

func solve(input string) (string, string) {
	firstPassword := nextValidPassword([]rune(input))
	secondPassword := nextValidPassword(firstPassword)
	return string(firstPassword), string(secondPassword)
}

func nextValidPassword(password []rune) []rune {
	for {
		if password = nextString(password); isPassword(password) {
			return password
		}
	}
}

func nextChar(ch rune) (rune, bool) {
	if ch == 'z' {
		return 'a', true
	}
	return ch + 1, false
}

func nextString(ch []rune) []rune {
	currentIndex := len(ch) - 1
	maybePassword := make([]rune, currentIndex+1)
	copy(maybePassword, ch)

	for currentIndex >= 0 {
		upNext, carry := nextChar(ch[currentIndex])
		maybePassword[currentIndex] = upNext

		if !carry {
			return maybePassword
		}

		currentIndex--
	}

	return maybePassword
}

func isPassword(chars []rune) bool {
	return !hasInvalidChar(chars) && hasIncreasing(chars) && repeatedCharCount(chars) >= 2
}

func hasIncreasing(chars []rune) (increasing bool) {
	i, j, k := 0, 1, 2
	for k < len(chars) {
		if chars[k]-chars[j] == 1 && chars[j]-chars[i] == 1 {
			return true
		}
		i++
		j++
		k++
	}
	return
}

func repeatedCharCount(chars []rune) (total int) {
	i, j := 0, 1
	for j < len(chars) {
		if chars[i] == chars[j] {
			total++
			i += 2
			j += 2
		} else {
			i++
			j++
		}
	}
	return total
}

func hasInvalidChar(chars []rune) (invalid bool) {
	for _, char := range chars {
		if char == 'i' || char == 'o' || char == 'l' {
			return true
		}
	}
	return
}
