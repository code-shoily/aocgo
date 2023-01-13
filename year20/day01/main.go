package day01

import (
	_ "embed"
	"fmt"
	"github.com/code-shoily/aocgo/algo"
	"github.com/code-shoily/aocgo/utils"
	"sort"
)

//go:embed input.txt
var input string

func Run() {
	fmt.Println(solve(input))
}

func solve(input string) (int, int) {
	expenses := parse(input)
	sort.Ints(expenses)
	return solvePart1(expenses), solvePart2(expenses)
}

func solvePart1(expenses []int) int {
	if a, b, ok := algo.TwoSum(expenses, 2020); ok {
		return a * b
	}

	panic("Nothing sums to 2020")
}

func solvePart2(expenses []int) int {
	if a, b, c, ok := algo.ThreeSum(expenses, 2020); ok {
		return a * b * c
	}

	panic("Nothing sums to 2020")
}

func parse(input string) (expenses []int) {
	return utils.SplitIntLines(input)
}
