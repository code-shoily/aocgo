// Package graphs contains graph data structures and algorithms needed to solve
// Advent of Code problems.
package graphs

import "errors"

var (
	ErrVertexExists       = errors.New("vertex exists")
	ErrVertexDoesNotExist = errors.New("vertex does not exist")
)

type graphSettings struct {
	isDirected bool
}

type Graph[T comparable] struct {
	vertices map[string]*Vertex[T]
	graphSettings
}

func (g *Graph[T]) Vertices() map[string]*Vertex[T] {
	return g.vertices
}

// AddVertex adds a single vertex to the graph
func (g *Graph[T]) AddVertex(v *Vertex[T]) error {
	if g.HasVertex(v.id) {
		return ErrVertexExists
	}
	g.vertices[v.id] = v
	return nil
}

// AddEdge creates a weighted edge between v1 and v2
func (g *Graph[T]) AddEdge(v1 string, v2 string, weight int) error {
	if !g.HasVertex(v1) || !g.HasVertex(v2) {
		return ErrVertexDoesNotExist
	}

	g.vertices[v1].AddConnection(g.vertices[v2], weight, !g.isDirected)

	return nil
}

// HasVertex checks if graph has vertex registered with given id
func (g *Graph[T]) HasVertex(key string) bool {
	_, exists := g.vertices[key]
	return exists
}

// NewGraph creates a new Graph with the given `isDirected` trait
func NewGraph[T comparable](isDirected bool) *Graph[T] {
	return &Graph[T]{
		vertices: make(map[string]*Vertex[T]),
		graphSettings: graphSettings{
			isDirected: isDirected,
		},
	}
}
