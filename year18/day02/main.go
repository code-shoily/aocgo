package day02

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

func solve(input string) (int, string) {
	ids := io.SplitLines(input)
	return solvePart1(ids), solvePart2(ids)
}

func solvePart1(ids []string) int {
	var twice, thrice int
	for _, id := range ids {
		hasTwice, hasThrice := twiceOrThrice(id)
		if hasTwice {
			twice++
		}

		if hasThrice {
			thrice++
		}
	}

	return twice * thrice
}

func solvePart2(ids []string) string {
	visitedPartialIds := make(map[string]int)
	for _, id := range ids {
		for i := 0; i < len(id); i++ {
			subId := id[:i] + id[i+1:]
			if idx, ok := visitedPartialIds[subId]; ok && idx == i {
				return subId
			}
			visitedPartialIds[subId] = i
		}
	}

	panic("ERROR: No matching ID found")
}

func twiceOrThrice(id string) (twice bool, thrice bool) {
	frequency := make(map[string]int)

	for _, ch := range io.SplitBy(id, "") {
		frequency[ch]++
	}

	for _, freq := range frequency {
		switch freq {
		case 2:
			twice = true
		case 3:
			thrice = true
		}
	}

	return twice, thrice
}
