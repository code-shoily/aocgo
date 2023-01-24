// Package day06 - Solution for Advent of Code 2019/06
// Problem Link: http://adventofcode.com/2019/day/06
package day06

import (
	_ "embed"
	"fmt"
	"github.com/code-shoily/aocgo/algo/graphs"
	"github.com/code-shoily/aocgo/seq"
	"strings"
)

//go:embed input.txt
var input string

type orbits graphs.PathMap

// Run prints out the result of the solution.
func Run() {
	fmt.Println(solve(input))
}

func solve(input string) (int, int) {
	paths := parse(input)
	return solvePart1(paths), solvePart2(paths)
}

func solvePart1(paths orbits) (total int) {
	for _, path := range paths {
		total += len(path)
	}

	return total
}

func solvePart2(paths orbits) (moves int) {
	from, to := paths["YOU"], paths["SAN"]
	fromSet, toSet := seq.MakeSet(from), seq.MakeSet(to)

	for i := len(from) - 1; i >= 0; i-- {
		if _, ok := toSet[from[i]]; ok {
			break
		}
		moves++
	}

	for i := len(to) - 1; i >= 0; i-- {
		if _, ok := fromSet[to[i]]; ok {
			break
		}
		moves++
	}

	return moves
}

func parse(input string) orbits {
	graph := graphs.NewGraph(true)

	for _, line := range strings.Split(input, "\n") {
		tokens := strings.Split(line, ")")
		v, w := tokens[0], tokens[1]
		graph.AddVertex(graphs.NewSimpleVertex(v))
		graph.AddVertex(graphs.NewSimpleVertex(w))
		if err := graph.AddEdge(v, w, 1); err != nil {
			panic("Invalid processing")
		}
	}

	return getAllOrbits(graph, "COM")
}

func getAllOrbits(graph *graphs.Graph, source string) orbits {
	return orbits(graphs.DFS(graph, source))
}
