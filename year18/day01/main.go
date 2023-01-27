package day01

import (
	_ "embed"
	"fmt"
	"github.com/code-shoily/aocgo/io"
)

//go:embed input.txt
var input string

func Run() {
	fmt.Println(solve(input))
}

func solve(input string) (int, int) {
	frequencies := parse(input)

	return solvePart1(frequencies), solvePart2(frequencies)
}

func solvePart1(frequencies []int) (frequency int) {
	for _, freq := range frequencies {
		frequency += freq
	}
	return frequency
}

func solvePart2(frequencies []int) (frequency int) {
	visited := make(map[int]bool)

	for i := 0; ; i = (i + 1) % len(frequencies) {
		visited[frequency] = true
		frequency += frequencies[i]

		if _, ok := visited[frequency]; ok {
			return frequency
		}
	}
}

func parse(input string) (frequencies []int) {
	return io.SplitIntLines(input)
}
