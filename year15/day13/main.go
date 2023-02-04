// Package day13 - Solution for Advent of Code 2015/13
// Problem Link: http://adventofcode.com/2015/day/13
package day13

import (
	_ "embed"
	"fmt"
	"github.com/code-shoily/aocgo/algo/graphs"
	"github.com/code-shoily/aocgo/seq"
	"golang.org/x/exp/maps"
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
	data := parse(input)
	return solvePart1(data), solvePart2(data)
}

func solvePart1(graph *graphs.Graph) (max int) {
	everyone := maps.Keys(graph.Vertices())
	for _, people := range seq.Permutations(everyone) {
		people = append(people, people[0])
		var happiness int
		for _, pair := range seq.Chunk(people, 2, 1, true) {
			happiness += getHappinessChange(graph, pair[0], pair[1])
		}

		if happiness > max {
			max = happiness
		}
	}

	return
}

func solvePart2(graph *graphs.Graph) int {
	me := graphs.NewSimpleVertex("me")
	if err := graph.AddVertex(me); err != nil {
		panic(err)
	}

	return solvePart1(graph)
}

func parse(input string) *graphs.Graph {
	happinessGraph := graphs.NewGraph(true)

	for _, line := range strings.Split(input, "\n") {
		tokens := strings.Split(strings.TrimRight(line, "."), " ")
		person1, person2, direction := tokens[0], tokens[10], tokens[2]
		happiness, err := strconv.Atoi(tokens[3])

		if err != nil {
			panic("Invalid input")
		}

		if direction == "lose" {
			happiness = -1 * happiness
		}
		addHappinessBetween(happinessGraph, person1, person2, happiness)
	}

	return happinessGraph
}

func getHappinessChange(graph *graphs.Graph, person1, person2 string) int {
	person1Vertex, person2Vertex := graph.Vertices()[person1], graph.Vertices()[person2]
	outgoing, incoming := person1Vertex.GetConnections()
	return outgoing[person2Vertex] + incoming[person2Vertex]
}

func addHappinessBetween(graph *graphs.Graph, person1, person2 string, happiness int) {
	person1Vertex := graphs.NewSimpleVertex(person1)
	person2Vertex := graphs.NewSimpleVertex(person2)
	err := graph.AddVertex(person1Vertex)
	err = graph.AddVertex(person2Vertex)
	err = graph.AddEdge(person1, person2, happiness)
	if err != nil {
		panic(err)
	}
}
