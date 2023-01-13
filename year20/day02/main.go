// Package day02 - Solution for Advent of Code 2020/02
// Problem Link: http://adventofcode.com/2020/day/02
package day02

import (
	_ "embed"
	"fmt"
	"github.com/code-shoily/aocgo/utils"
	"strings"
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

func solvePart1(data []PasswordData) (total int) {
	for _, pwd := range data {
		if pwd.isValidRule1() {
			total++
		}
	}

	return total
}

func solvePart2(data []PasswordData) (total int) {
	for _, pwd := range data {
		if pwd.isValidRule2() {
			total++
		}
	}

	return total
}

func parse(input string) (data []PasswordData) {
	for _, line := range utils.SplitLines(input) {
		data = append(data, newPasswordData(line))
	}

	return data
}

func parseLine(line string) []string {
	return strings.FieldsFunc(line, func(r rune) bool {
		return r == ' ' || r == ':'
	})
}

type PasswordData struct {
	start    int
	end      int
	char     string
	password string
}

func (pwd PasswordData) isValidRule1() bool {
	occurrence := strings.Count(pwd.password, pwd.char)
	return occurrence >= pwd.start && occurrence <= pwd.end
}

func (pwd PasswordData) isValidRule2() bool {
	a := string(pwd.password[pwd.start-1])
	b := string(pwd.password[pwd.end-1])
	char := pwd.char
	return (a == char && b != char) || (b == char && a != char)
}

func newPasswordData(line string) PasswordData {
	tokens := parseLine(line)
	charRange := utils.SplitByInts(tokens[0], "-")

	return PasswordData{charRange[0], charRange[1], tokens[1], tokens[2]}
}
