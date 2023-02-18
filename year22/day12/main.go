// Package day12 - Solution for Advent of Code 2022/12
// Problem Link: http://adventofcode.com/2022/day/12
package day12

import (
	_ "embed"
	"fmt"
	"github.com/code-shoily/aocgo/algo"
	"strings"
)

//go:embed input.txt
var input string

// Run prints out the result of the solution.
func Run() {
	fmt.Println(solve(input))
}

func solve(input string) (int, int) {
	data, src, dst := parse(input)
	return solvePart1(data, src, dst), solvePart2(data, dst)
}

func solvePart1(heightMap [][]rune, src, dst [2]int) int {
	satisfies := func(hill [2]int) bool { return hill == dst }
	return doBFS(heightMap, src, false, satisfies)
}

func solvePart2(heightMap [][]rune, dst [2]int) int {
	satisfies := func(hill [2]int) bool { return heightMap[hill[0]][hill[1]] == 'a' }
	return doBFS(heightMap, dst, true, satisfies)
}

func parse(input string) (data [][]rune, start [2]int, stop [2]int) {
	for x, line := range strings.Split(input, "\n") {
		hills := make([]rune, len(line))
		for y, hill := range line {
			if hill == 'S' {
				start = [2]int{x, y}
				hill = 'a'
			}
			if hill == 'E' {
				stop = [2]int{x, y}
				hill = 'z'
			}
			hills[y] = hill
		}
		data = append(data, hills)
	}

	return
}

type frame struct {
	height [2]int
	step   int
}

func doBFS(heightMap [][]rune, src [2]int, inverse bool, satisfies func([2]int) bool) int {
	queue, visited := algo.NewQueue(), make(map[[2]int]bool)
	queue.Enqueue(frame{src, 0})
	visited[src] = true

	for !queue.IsEmpty() {
		currentStep := queue.MustDequeue().(frame)
		location := currentStep.height
		for _, hill := range steps(heightMap, location, inverse) {
			if satisfies([2]int{hill[0], hill[1]}) {
				return currentStep.step + 1
			}
			if _, visit := visited[hill]; !visit {
				queue.Enqueue(frame{hill, currentStep.step + 1})
				visited[hill] = true
			}
		}
	}
	panic("unreachable code")
}

func steps(heightMap [][]rune, src [2]int, inverse bool) (hills [][2]int) {
	dirs := [][2]int{
		{0, -1}, // up
		{0, 1},  // down
		{-1, 0}, // left
		{1, 0},  // right
	}

	canStep := func(neighbour [2]int) bool {
		if step := heightMap[neighbour[0]][neighbour[1]] - heightMap[src[0]][src[1]]; inverse {
			return step >= -1
		} else {
			return step <= 1
		}
	}

	for _, dir := range dirs {
		x, y := src[0]+dir[0], src[1]+dir[1]
		withinRange := x >= 0 && y >= 0 && x < len(heightMap) && y < len(heightMap[0])

		if neighbour := [2]int{x, y}; withinRange && canStep(neighbour) {
			hills = append(hills, neighbour)
		}
	}

	return hills
}
