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
	for _, line := range strings.Split(input, "\n") {
		if mass, err := strconv.Atoi(line); err == nil {
			masses = append(masses, mass)
		}
	}

	return masses
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
