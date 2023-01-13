package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

func main() {
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
	for _, freqString := range strings.Split(input, "\n") {
		if freq, err := strconv.Atoi(freqString); err == nil {
			frequencies = append(frequencies, freq)
		}
	}

	return frequencies
}
