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
	masses := parse(input)

	return solvePart1(masses), solvePart2(masses)
}

func solvePart1(masses []int) (totalFuel int) {
	for _, mass := range masses {
		totalFuel += getFuel1(mass)
	}

	return totalFuel
}

func solvePart2(masses []int) (totalFuel int) {
	for _, mass := range masses {
		totalFuel += getFuel2(mass)
	}

	return totalFuel
}

func parse(input string) (masses []int) {
	return utils.SplitIntLines(input)
}

func getFuel1(mass int) int {
	return mass/3 - 2
}

func getFuel2(mass int) (fuel int) {
	for {
		mass = getFuel1(mass)
		if mass <= 0 {
			return fuel
		}
		fuel += mass
	}
}
