package day01

import (
	_ "embed"
	"fmt"
)

//go:embed input.txt
var input []byte

func Run() {
	fmt.Println(solve())
}

func solvePart1(input []byte) (floor int) {
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

	panic("Will never reach basement.")
}

func solve() (int, int) {
	return solvePart1(input), solvePart2(input)
}
