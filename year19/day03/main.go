// Package day03 - Solution for Advent of Code 2019/03
// Problem Link: http://adventofcode.com/2019/day/03
package day03

import (
	_ "embed"
	"fmt"
	"math"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

// Run prints out the result of the solution.
func Run() {
	fmt.Println(solve(input))
}

func solve(input string) (int, int) {
	wire1, wire2 := parse(input)
	return solvePart1(wire1, wire2), solvePart2(wire1, wire2)
}

func solvePart1(wire1, wire2 wire) int {
	shortestDistance := math.MaxInt

	for pos := range wire1.visits {
		if _, ok := wire2.visits[pos]; ok {
			if distance := manhattanDistance(pos); shortestDistance > distance {
				shortestDistance = distance
			}
		}
	}

	return shortestDistance
}

func solvePart2(wire1, wire2 wire) int {
	nearestIntersection := math.MaxInt

	for pos := range wire1.visits {
		if _, ok := wire2.visits[pos]; ok {
			if steps := wire1.visits[pos] + wire2.visits[pos]; steps < nearestIntersection {
				nearestIntersection = steps
			}
		}
	}

	return nearestIntersection
}

func parse(input string) (wire, wire) {
	var units1, units2 []position
	wires := strings.Split(input, "\n")

	for _, by := range strings.Split(wires[0], ",") {
		units1 = append(units1, pullWire(by)...)
	}

	for _, by := range strings.Split(wires[1], ",") {
		units2 = append(units2, pullWire(by)...)
	}

	return pulledWires(units1, units2)
}

type position [2]int

type wire struct {
	current position
	visits  map[position]int
}

func (w *wire) walk(units []position) {
	for step, unit := range units {
		w.current = translate(w.current, unit)
		w.visits[w.current] = step + 1
	}
}

func pulledWires(units1, units2 []position) (wire, wire) {
	wire1, wire2 := wire{visits: map[position]int{}}, wire{visits: map[position]int{}}

	wire1.walk(units1)
	wire2.walk(units2)

	return wire1, wire2
}

func pullWire(by string) (units []position) {
	tokens := strings.SplitN(by, "", 2)
	if steps, err := strconv.Atoi(tokens[1]); err == nil {
		repeatWith := func(vector position) []position {
			for i := 0; i < steps; i++ {
				units = append(units, vector)
			}
			return units
		}

		switch tokens[0] {
		case "L":
			units = repeatWith(position{-1, 0})
		case "R":
			units = repeatWith(position{1, 0})
		case "U":
			units = repeatWith(position{0, 1})
		case "D":
			units = repeatWith(position{0, -1})
		}
	}

	return units
}

func manhattanDistance(p position) int {
	return int(math.Abs(float64(0-p[0]))) + int(math.Abs(float64(0-p[1])))
}

func translate(current, by position) position {
	return position{current[0] + by[0], current[1] + by[1]}
}
