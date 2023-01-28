// Package day07 - Solution for Advent of Code 2022/07
// Problem Link: http://adventofcode.com/2022/day/07
package day07

import (
	_ "embed"
	"fmt"
	"github.com/code-shoily/aocgo/algo"
	"math"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

const (
	capacity    = 70_000_000
	spaceNeeded = 30_000_000
)

// Run prints out the result of the solution.
func Run() {
	fmt.Println(solve(input))
}

func solve(input string) (int, int) {
	data := parse(input)
	return solvePart1(data), solvePart2(data)
}

func solvePart1(data map[string]int) (total int) {
	for _, size := range data {
		if size <= 100_000 {
			total += size
		}
	}

	return total
}

func solvePart2(data map[string]int) int {
	emptyingNeeded := spaceNeeded - (capacity - data["root"])
	dirToDelete := math.MaxInt

	for _, size := range data {
		if size >= emptyingNeeded && size < dirToDelete {
			dirToDelete = size
		}
	}

	return dirToDelete
}

func parse(commands string) map[string]int {
	paths, stats := algo.Stack[string]{}, make(map[string]int)

	for _, line := range strings.Split(commands, "\n") {
		switch tokens := strings.Split(line, " "); {
		case line == "$ ls", strings.HasPrefix(line, "dir"):
		case line == "$ cd /":
			paths.Push("root")
		case line == "$ cd ..":
			paths.Pop()
		case strings.HasPrefix(line, "$ cd"):
			parent, _ := paths.Peek()
			paths.Push(parent + "/" + tokens[2])
		default:
			size, _ := strconv.Atoi(tokens[0])
			for _, ancestor := range paths {
				stats[ancestor] += size
			}
		}
	}

	return stats
}
