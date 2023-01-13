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
	sweeps := parse(input)

	return nthIncreasing(sweeps, 1), nthIncreasing(sweeps, 3)
}

func nthIncreasing(sweeps []int, n int) (incr int) {
	for i := 0; i < len(sweeps)-n; i++ {
		if sweeps[i] < sweeps[i+n] {
			incr++
		}
	}

	return incr
}

func parse(input string) (sweeps []int) {
	for _, line := range strings.Split(input, "\n") {
		if sweep, error := strconv.Atoi(line); error == nil {
			sweeps = append(sweeps, sweep)
		}
	}

	return sweeps
}
