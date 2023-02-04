// Package day14 - Solution for Advent of Code 2015/14
// Problem Link: http://adventofcode.com/2015/day/14
package day14

import (
	_ "embed"
	"fmt"
	"strings"
)

//go:embed input.txt
var input string

// Run prints out the result of the solution.
func Run() {
	fmt.Println(solve(input))
}

func solve(input string) (int, int) {
	return solvePart1(input), solvePart2(input)
}

func solvePart1(input string) (winner int) {
	data := parse(input)
	for _, reindeer := range data {
		for i := 1; i <= 2503; i++ {
			reindeer.moveOnce(i)
		}
		if distance := reindeer.distance; distance > winner {
			winner = distance
		}
	}
	return winner
}

func solvePart2(input string) (maxPoint int) {
	data := parse(input)

	for i := 1; i <= 2503; i++ {
		for _, reindeer := range data {
			reindeer.moveOnce(i)
		}

		var max int
		for _, reindeer := range data {
			if reindeer.distance > max {
				max = reindeer.distance
			}
		}

		for _, reindeer := range data {
			if reindeer.distance == max {
				reindeer.point++
				if reindeer.point > maxPoint {
					maxPoint = reindeer.point
				}
			}
		}
	}

	return
}

func parse(input string) (reindeer []*Reindeer) {
	for _, line := range strings.Split(input, "\n") {
		var name string
		var speed, fly, rest int
		fmt.Sscanf(line, "%s can fly %d km/s for %d seconds, but then must rest for %d seconds.",
			&name, &speed, &fly, &rest)
		reindeer = append(reindeer, &Reindeer{speed: speed, fly: fly, rest: rest})
	}

	return
}

type Reindeer struct {
	speed    int
	fly      int
	rest     int
	point    int
	distance int
}

func (r *Reindeer) moveOnce(at int) {
	if x := at % (r.fly + r.rest); x != 0 && x <= r.fly {
		r.distance += r.speed
	}
}
