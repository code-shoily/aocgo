// Package day21 - Solution for Advent of Code 2016/21
// Problem Link: http://adventofcode.com/2016/day/21
package day21

import (
	_ "embed"
	"fmt"
	"github.com/code-shoily/aocgo/seq"
	"strings"
)

//go:embed input.txt
var scrambleRules string

// Run prints out the result of the solution.
func Run() {
	fmt.Println(solve(scrambleRules))
}

func solve(input string) (string, string) {
	initial, target := "abcdefgh", "fbgdceah"

	solvePart1 := func() string {
		return scramble(initial, input)
	}

	solvePart2 := func() string {
		for _, password := range seq.StringPermutations(initial) {
			if scramble(password, input) == target {
				return password
			}
		}
		panic("This code should be unreachable")
	}

	return solvePart1(), solvePart2()
}

func scramble(password string, rules string) string {
	for _, line := range strings.Split(rules, "\n") {
		switch {
		case strings.HasPrefix(line, "swap position"):
			var x, y int
			fmt.Sscanf(line, "swap position %d with position %d", &x, &y)
			password = swapPositions(password, x, y)
		case strings.HasPrefix(line, "swap letter"):
			var x, y string
			fmt.Sscanf(line, "swap letter %s with letter %s", &x, &y)
			password = swapLetters(password, x, y)
		case strings.HasPrefix(line, "reverse"):
			var from, to int
			fmt.Sscanf(line, "reverse positions %d through %d", &from, &to)
			password = reverse(password, from, to)
		case strings.HasPrefix(line, "rotate right"):
			var step int
			fmt.Sscanf(line, "rotate right %d step", &step)
			password = rotateRight(password, step)
		case strings.HasPrefix(line, "rotate left"):
			var step int
			fmt.Sscanf(line, "rotate left %d step", &step)
			password = rotateLeft(password, step)
		case strings.HasPrefix(line, "rotate based"):
			var x string
			fmt.Sscanf(line, "rotate based on position of letter %s", &x)
			password = rotateRelative(password, x)
		case strings.HasPrefix(line, "move"):
			var from, to int
			fmt.Sscanf(line, "move position %d to position %d", &from, &to)
			password = move(password, from, to)
		}
	}
	return password
}

func swapPositions(pwd string, x, y int) string {
	letters := []byte(pwd)
	letters[x], letters[y] = letters[y], letters[x]

	return string(letters)
}

func swapLetters(pwd, x, y string) string {
	return swapPositions(pwd, indexOf(pwd, x), indexOf(pwd, y))
}

func reverse(pwd string, x, y int) string {
	letters := []byte(pwd)
	temp := make([]byte, y-x+1)
	copy(temp, letters[x:y+1])

	for i, j := y, 0; i >= x; i-- {
		letters[i] = temp[j]
		j++
	}

	return string(letters)
}

func rotateRight(pwd string, step int) string {
	reference, letters := []byte(pwd), []byte(pwd)

	for i := range reference {
		diff := i + step

		if diff >= len(reference) {
			diff = diff % len(reference)
		}
		letters[diff] = reference[i]
	}

	return string(letters)
}

func rotateLeft(pwd string, step int) string {
	reference, letters := []byte(pwd), []byte(pwd)

	for i := range reference {
		diff := i - step

		if diff < 0 {
			diff = len(reference) + diff
		}
		letters[diff] = reference[i]
	}

	return string(letters)
}

func rotateRelative(pwd string, letter string) string {
	rotateBy := indexOf(pwd, letter)
	return rotateRight(pwd, 1+rotateBy+rotateBy/4)
}

func move(pwd string, from, to int) string {
	letters := []byte(pwd)
	store := letters[from]

	if from < to {
		for i, replace := to, letters[to]; i > from; i-- {
			temp := letters[i-1]
			letters[i-1] = replace
			replace = temp
		}
	} else {
		for i, replace := to, letters[to]; i < from; i++ {
			temp := letters[i+1]
			letters[i+1] = replace
			replace = temp
		}
	}

	letters[to] = store

	return string(letters)
}

func indexOf(pwd string, char string) int {
	for idx, ch := range pwd {
		if string(ch) == char {
			return idx
		}
	}
	return -1
}
