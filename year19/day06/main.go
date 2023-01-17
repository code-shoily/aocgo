// Package day06 - Solution for Advent of Code 2019/06
// Problem Link: http://adventofcode.com/2019/day/06
package day06

import (
	_ "embed"
	"fmt"
	"github.com/code-shoily/aocgo/algo"
	"github.com/code-shoily/aocgo/algo/graphs"
	"github.com/code-shoily/aocgo/seq"
	"strings"
)

//go:embed input.txt
var input string

// Run prints out the result of the solution.
func Run() {
	fmt.Println(solve(input))
}

func solve(input string) (int, int) {
	paths := parse(input)
	return solvePart1(paths), solvePart2(paths)
}

func solvePart1(paths map[string][]string) (total int) {
	for _, path := range paths {
		total += len(path)
	}

	return total
}

func solvePart2(paths map[string][]string) (moves int) {
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

func parse(input string) map[string][]string {
	graph := graphs.NewGraph[string](true)

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

func getAllOrbits(graph *graphs.Graph[string], source string) map[string][]string {
	// FIXME: This should be moved in graph module as DFS or something
	paths := map[string][]string{}
	vertices := graph.Vertices()

	v := vertices[source]
	stack := algo.Stack[string]{}
	stack.Push(v.ID())

	for !stack.IsEmpty() {
		vID, _ := stack.Pop()
		vertex := vertices[vID]
		outgoing, _ := vertex.Connections()

		for neighbour, _ := range outgoing {
			stack.Push(neighbour.ID())
			paths[neighbour.ID()] = append(paths[vertex.ID()], neighbour.ID())
		}
	}

	return paths
}
