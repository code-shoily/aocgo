package main

import (
	_ "embed"
	"fmt"
)

//go:embed input.txt
var input []byte

func solvePart1(input []byte) int {
	floor := 0
	for _, instruction := range input {
		if instruction == '(' {
			floor += 1
		} else {
			floor -= 1
		}
	}

	return floor
}

func solvePart2(input []byte) int {
	floor := 0
	for idx, instruction := range input {
		if instruction == '(' {
			floor += 1
		} else {
			floor -= 1
			if floor == -1 {
				return idx + 1
			}
		}
	}

	return -1
}

func solve() (int, int) {
	return solvePart1(input), solvePart2(input)
}

func main() {
	fmt.Println(solve())
}
