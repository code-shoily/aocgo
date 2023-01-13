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
	captcha := parse(input)
	size := len(captcha)

	return solvePart1(captcha, size), solvePart2(captcha, size)
}

func solvePart1(input []int, size int) int {
	return solveCaptcha(input, size, 1)
}

func solvePart2(input []int, size int) int {
	return solveCaptcha(input, size, size/2)
}

func solveCaptcha(input []int, size, offset int) (total int) {
	for i := 0; i < size; i++ {
		if input[i] == input[(i+offset)%size] {
			total += input[i]
		}
	}

	return total
}

func parse(input string) (values []int) {
	return utils.SplitByInts(input, "")
}
