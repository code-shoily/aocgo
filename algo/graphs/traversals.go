package graphs

import (
	"errors"
	"github.com/code-shoily/aocgo/algo"
)

type Path []string
type PathMap map[string]Path

var (
	ErrBrokenPath = errors.New("broken path")
)

// DFS returns single Path PathMap reachable from S as {"D": ["a", "b", "D"]}. S isn't part of the Path.
func DFS(graph *Graph, source string) PathMap {
	paths := make(PathMap)
	vertices := graph.Vertices()
	stack := algo.Stack[*Vertex]{}
	stack.Push(vertices[source])

	for !stack.IsEmpty() {
		current, _ := stack.Pop()
		outgoing, _ := current.GetConnections()

		for neighbour, _ := range outgoing {
			stack.Push(neighbour)
			if _, exists := paths[neighbour.ID()]; !exists {
				paths[neighbour.ID()] = append(paths[current.ID()], neighbour.ID())
			}
		}
	}

	return paths
}

// PathDistance calculates the distance of a path within a graph, if valid, otherwirse returns error
func PathDistance(g *Graph, path Path) (distance int, err error) {
	previousNode := path[0]
	for _, currentNode := range path[1:] {
		v, _ := g.vertices[previousNode]
		w, _ := g.vertices[currentNode]

		weight, pathExists := v.outgoing[w]
		if !pathExists {
			return distance, ErrBrokenPath
		}
		distance += weight
		previousNode = currentNode
	}

	return distance, err
}
