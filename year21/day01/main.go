package day01

import (
	_ "embed"
	"fmt"
	"github.com/code-shoily/aocgo/utils"
)

//go:embed input.txt
var input string

func Run() {
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
	return utils.SplitIntLines(input)
}
