// Package day09 - Solution for Advent of Code 2015/09
// Problem Link: http://adventofcode.com/2015/day/09
package day09

import (
	_ "embed"
	"fmt"
	"github.com/code-shoily/aocgo/algo/graphs"
	"github.com/code-shoily/aocgo/seq"
	"golang.org/x/exp/maps"
	"math"
	"strings"
)

//go:embed input.txt
var input string

// Run prints out the result of the solution.
func Run() {
	fmt.Println(solve(input))
}

func solve(input string) (int, int) {
	graph := parse(input)
	cities := maps.Keys(graph.Vertices())
	min, max := math.MaxInt, math.MinInt

	for _, path := range seq.Permutations(cities) {
		if distance, err := graphs.PathDistance(graph, path); err == nil {
			if distance < min {
				min = distance
			}
			if distance > max {
				max = distance
			}
		}
	}

	return min, max
}

func parse(input string) *graphs.Graph {
	graph := graphs.NewGraph(false)

	for _, line := range strings.Split(input, "\n") {
		var city1, city2 string
		var distance int

		_, err := fmt.Sscanf(line, "%s to %s = %d", &city1, &city2, &distance)
		err = graph.AddVertex(graphs.NewSimpleVertex(city1))
		err = graph.AddVertex(graphs.NewSimpleVertex(city2))
		err = graph.AddEdge(city1, city2, distance)

		if err != nil {
			panic("Invalid parsing")
		}
	}

	return graph
}
