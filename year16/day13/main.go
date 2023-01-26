// Package day13 - Solution for Advent of Code 2016/13
// Problem Link: http://adventofcode.com/2016/day/13
package day13

import (
	_ "embed"
	"fmt"
	"github.com/code-shoily/aocgo/algo"
	"github.com/code-shoily/aocgo/utils"
)

//go:embed input.txt
var input string

// Run prints out the result of the solution.
func Run() {
	fmt.Println(solve(input))
}

func solve(input string) (int, int) {
	return doBFS(
		cubicle{1, 1, 0},
		cubicle{31, 39, 0},
		utils.SplitIntLines(input)[0],
	)
}

type cubicle struct {
	x    int
	y    int
	dist int
}

func (c cubicle) cubicleAt(direction [2]int) cubicle {
	return cubicle{c.x + direction[0], c.y + direction[1], c.dist + 1}
}

func (c cubicle) getLoc() [2]int {
	return [2]int{c.x, c.y}
}

func (c cubicle) isOutside() bool {
	return c.x < 0 || c.y < 0
}

var diffs = [4][2]int{
	{1, 0},  // Up
	{-1, 0}, // Down
	{0, -1}, // Left
	{0, 1},  // Right
}

func doBFS(src, dst cubicle, input int) (distance int, steps int) {
	visits, queue := map[[2]int]bool{}, algo.NewQueue()

	queue.Enqueue(src)
	visits[src.getLoc()] = true

	for !queue.IsEmpty() {
		c, _ := queue.Dequeue()
		current := c.(cubicle)

		if current.getLoc() == dst.getLoc() {
			return current.dist, steps + 1
		}

		for _, diff := range diffs {
			adj := current.cubicleAt(diff)
			loc := adj.getLoc()

			if _, visited := visits[loc]; !visited {
				if !adj.isOutside() && isOpenSpace(loc, input) {
					queue.Enqueue(adj)
					visits[loc] = true

					if adj.dist <= 50 {
						steps++
					}
				}
			}
		}
	}

	return
}

func isOpenSpace(loc [2]int, input int) bool {
	x, y := loc[0], loc[1]
	setBitCount := algo.CountSetBits(x*x + 3*x + 2*x*y + y + y*y + input)

	return algo.IsEven(setBitCount)
}
