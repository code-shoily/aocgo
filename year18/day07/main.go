// Package day07 - Solution for Advent of Code 2018/07
// Problem Link: http://adventofcode.com/2018/day/07
package day07

import (
	_ "embed"
	"fmt"
	"github.com/code-shoily/aocgo/algo/graphs"
	"github.com/code-shoily/aocgo/io"
	"strings"
)

//go:embed input.txt
var input string

// Run prints out the result of the solution.
func Run() {
	fmt.Println(solve(input))
}

func solve(input string) (string, int) {
	data := parse(input)
	return solvePart1(data), solvePart2(data)
}

func solvePart1(g *graphs.Graph) string {
	if sorted, cycle := graphs.LexicographicalTopologicalSort(g); !cycle {
		return strings.Join(sorted, "")
	}
	panic("Cycle detected :O")
}

func solvePart2(g *graphs.Graph) int {
	return 0
}

func parse(input string) *graphs.Graph {
	var dependencyPairs [][2]string
	for _, line := range io.SplitLines(input) {
		var from, to string
		fmt.Sscanf(line, "Step %s must be finished before step %s can begin.", &from, &to)
		dependencyPairs = append(dependencyPairs, [2]string{from, to})
	}

	return MakeGraph(dependencyPairs)
}

func MakeGraph(deps [][2]string) *graphs.Graph {
	graph := graphs.NewGraph(true)
	vertices := map[string]*graphs.Vertex{}
	for _, dep := range deps {
		from := getVertex(dep[0], vertices)
		to := getVertex(dep[1], vertices)

		graph.AddVertex(from)
		graph.AddVertex(to)
		graph.AddEdge(from.ID(), to.ID(), 1)
	}

	return graph
}

func getVertex(key string, vertexMap map[string]*graphs.Vertex) *graphs.Vertex {
	if vertex, exists := vertexMap[key]; exists {
		return vertex
	}
	vertexMap[key] = graphs.NewSimpleVertex(key)

	return vertexMap[key]
}
