package day02

import (
	_ "embed"
	"fmt"
	"github.com/code-shoily/aocgo/seq"
	"github.com/code-shoily/aocgo/utils"
)

//go:embed input.txt
var input string

// Run returns the solution
func Run() {
	fmt.Println(solve(input))
}

func solve(input string) (int, int) {
	data := parse(input)
	return solvePart1(data), solvePart2(data)
}

func parse(input string) (data [][]int) {
	for _, line := range utils.SplitLines(input) {
		var cells []int
		data = append(data, append(cells, utils.SplitByInts(line, "\t")...))
	}

	return data
}

func solvePart1(data [][]int) (result int) {
	for _, row := range data {
		min, max := seq.GetMinMax(row)
		result += max - min
	}

	return result
}

func solvePart2(data [][]int) (result int) {
	for _, row := range data {
		for i, a := range row {
			for _, b := range row[i+1:] {
				if a%b == 0 {
					result += a / b
					break
				}

				if b%a == 0 {
					result += b / a
					break
				}
			}
		}
	}

	return result
}
