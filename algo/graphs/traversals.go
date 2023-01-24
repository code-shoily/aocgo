package graphs

import "github.com/code-shoily/aocgo/algo"

type Path []string
type PathMap map[string]Path

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
