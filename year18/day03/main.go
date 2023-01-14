// Package day03 - Solution for Advent of Code 2018/03
// Problem Link: http://adventofcode.com/2018/day/03
package day03

import (
	_ "embed"
	"fmt"
	"github.com/code-shoily/aocgo/utils"
)

//go:embed input.txt
var input string

// Run prints out the result of the solution.
func Run() {
	fmt.Println(solve(input))
}

func solve(input string) (int, int) {
	claims, tally := parse(input)
	return solvePart1(tally), solvePart2(claims, tally)
}

func solvePart1(tally FabricTally) (total int) {
	for _, count := range tally {
		if count > 1 {
			total++
		}
	}

	return total
}

func solvePart2(claims []Claim, tally FabricTally) int {
	doesNotOverlap := func(fabrics [][2]int) bool {
		for _, fabric := range fabrics {
			if tally[fabric] != 1 {
				return false
			}
		}
		return true
	}

	for _, claim := range claims {
		if doesNotOverlap(claim.fabrics) {
			return claim.id
		}
	}

	panic("[ERROR] No tally with single claim")
}

func parse(input string) (claims []Claim, tally FabricTally) {
	for _, line := range utils.SplitLines(input) {
		claims = append(claims, newClaim(line))
	}

	tally = make(map[[2]int]int)

	for _, claim := range claims {
		claim.fill(tally)
	}

	return claims, tally
}

type FabricTally = map[[2]int]int

type Claim struct {
	id      int
	x       int
	y       int
	width   int
	height  int
	fabrics [][2]int
}

func (claim *Claim) fill(tally FabricTally) {
	var idx int
	for i := claim.x; i < claim.x+claim.width; i++ {
		for j := claim.y; j < claim.y+claim.height; j++ {
			tally[[2]int{i, j}]++
			claim.fabrics[idx] = [2]int{i, j}
			idx++
		}
	}
}

func newClaim(line string) Claim {
	var id, x, y, width, height int
	fmt.Sscanf(line, "#%d @ %d,%d: %dx%d", &id, &x, &y, &width, &height)

	return Claim{
		id,
		x,
		y,
		width,
		height,
		make([][2]int, width*height),
	}
}
