// Package day07 - Solution for Advent of Code 2020/07
// Problem Link: http://adventofcode.com/2020/day/07
package day07

import (
	_ "embed"
	"fmt"
	"github.com/code-shoily/aocgo/algo/graphs"
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
	luggages := parse(input)
	return solvePart1(luggages), solvePart2(luggages)
}

func solvePart1(luggages []luggage) int {
	reversedGraph := luggageToGraph(luggages, true)
	source := reversedGraph.Vertices()["shiny gold"]
	return len(graphs.DFS(reversedGraph, source.ID()))
}

func solvePart2(luggages []luggage) int {
	graph := luggageToGraph(luggages, false)
	source := graph.Vertices()["shiny gold"]
	return countBags(graph, source)
}

func parse(input string) (luggages []luggage) {
	for _, info := range strings.Split(input, "\n") {
		tokens := strings.Split(strings.Trim(info, "."), " bags contain ")

		for _, container := range strings.Split(tokens[1], ",") {
			colour, quantity := parseBag(container)
			luggages = append(luggages, *newLuggage(tokens[0], colour, quantity))
		}
	}
	return luggages
}

func parseBag(container string) (bag string, qty int) {
	tokens := strings.Split(strings.TrimSpace(container), " ")
	if quantity, err := strconv.Atoi(tokens[0]); err == nil {
		return strings.Join(tokens[1:len(tokens)-1], " "), quantity
	}
	return
}

type luggage struct {
	container string
	bag       string
	quantity  int
}

func newLuggage(containerColour, bagColour string, quantity int) *luggage {
	return &luggage{containerColour, bagColour, quantity}
}

func luggageToGraph(luggage []luggage, transposed bool) *graphs.Graph[string] {
	graph := graphs.NewGraph[string](true)

	for _, ll := range luggage {
		graph.AddVertex(graphs.NewSimpleVertex(ll.container))
		graph.AddVertex(graphs.NewSimpleVertex(ll.bag))

		if transposed {
			graph.AddEdge(ll.bag, ll.container, ll.quantity)
		} else {
			graph.AddEdge(ll.container, ll.bag, ll.quantity)
		}
	}

	return graph
}

func countBags(graph *graphs.Graph[string], from *graphs.Vertex[string]) (bags int) {
	connection, _ := from.Connections()

	for to, distance := range connection {
		bags += distance + distance*countBags(graph, to)
	}

	return bags
}
