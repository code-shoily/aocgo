package day02

import (
	_ "embed"
	"fmt"
	"github.com/code-shoily/aocgo/utils"
	"strings"
)

//go:embed input.txt
var input string

var dialPad1 = [][]string{
	{" ", " ", " ", " ", " "},
	{" ", "1", "2", "3", " "},
	{" ", "4", "5", "6", " "},
	{" ", "7", "8", "9", " "},
	{" ", " ", " ", " ", " "},
}

var dialPad2 = [][]string{
	{" ", " ", " ", " ", " ", " ", " "},
	{" ", " ", " ", "1", " ", " ", " "},
	{" ", " ", "2", "3", "4", " ", " "},
	{" ", "5", "6", "7", "8", "9", " "},
	{" ", " ", "A", "B", "C", " ", " "},
	{" ", " ", " ", "D", " ", " ", " "},
	{" ", " ", " ", " ", " ", " ", " "},
}

// Run prints the solution for the day
func Run() {
	fmt.Println(solve(input))
}

func solve(input string) (string, string) {
	instructions := parse(input)

	return solveFor(instructions, dialPad1, 2, 2), solveFor(instructions, dialPad2, 3, 1)
}

func parse(input string) (instructions [][]string) {
	for _, line := range utils.SplitLines(input) {
		var instructionRow []string
		instructions = append(instructions, append(instructionRow, utils.SplitBy(line, "")...))
	}

	return instructions
}

func solveFor(instructions [][]string, dialPad [][]string, ud int, lr int) string {
	var passwordChar []string

	for _, instructionRow := range instructions {
		for _, instruction := range instructionRow {
			switch instruction {
			case "L":
				if dialPad[ud][lr-1] != " " {
					lr--
				}
			case "R":
				if dialPad[ud][lr+1] != " " {
					lr++
				}
			case "U":
				if dialPad[ud-1][lr] != " " {
					ud--
				}
			case "D":
				if dialPad[ud+1][lr] != " " {
					ud++
				}
			}
		}
		passwordChar = append(passwordChar, dialPad[ud][lr])
	}

	return strings.Join(passwordChar, "")
}
