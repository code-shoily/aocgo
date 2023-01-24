package graphs

//FIXME: Decompose those large ass test funcs.

import (
	"fmt"
	"testing"
)

func TestNewGraph(t *testing.T) {
	examples := []bool{true, false}

	for _, isDirected := range examples {
		name := fmt.Sprintf("testing for isDirected %v", isDirected)

		t.Run(name, func(tt *testing.T) {
			g := NewGraph(isDirected)
			if g.isDirected != isDirected {
				tt.Errorf("Fail - isDirected expected %v but got %v", isDirected, g.isDirected)
			}
		})
	}
}

func TestGraph_AddVertex(t *testing.T) {
	g := NewGraph(false)
	a := NewSimpleVertex("a")
	b := NewSimpleVertex("b")

	if err := g.AddVertex(a); err != nil {
		t.Errorf("Failed when adding %v, found error %v", a.id, err)
	}

	if err := g.AddVertex(b); err != nil {
		t.Errorf("Failed when adding %v, found error %v", b.id, err)
	}
}

func TestGraph_AddVertexErrorsOnDuplicate(t *testing.T) {
	g := NewGraph(false)
	a := NewSimpleVertex("a")
	a2 := NewSimpleVertex("a")

	if err := g.AddVertex(a); err != nil {
		t.Errorf("Failed when adding %v, found error %v", a.id, err)
	}

	if err := g.AddVertex(a2); err == nil {
		t.Errorf("Failed when adding %v, did not find error %v", a2.id, err)
	}

	if vertices := len(g.vertices); vertices != 1 {
		t.Errorf("Fail - should be vertex# %v but gound %v", 1, vertices)
	}
}

func TestGraph_HasVertex(t *testing.T) {
	g := NewGraph(false)
	a := NewSimpleVertex("a")
	b := NewSimpleVertex("b")
	g.AddVertex(a)

	if hasA := g.HasVertex(a.id); !hasA {
		t.Errorf("Fail - expected to be %v but got %v", true, hasA)
	}

	if hasB := g.HasVertex(b.id); hasB {
		t.Errorf("Fail - expected to be %v but got %v", false, hasB)
	}
}

func TestGraph_AddEdgeDirected(t *testing.T) {
	g := NewGraph(true)

	vertices := []*Vertex{
		NewSimpleVertex("a"),
		NewSimpleVertex("b"),
		NewSimpleVertex("c"),
		NewSimpleVertex("d"),
	}

	// Add vertices
	for _, v := range vertices {
		if err := g.AddVertex(v); err != nil {
			t.Errorf("Fail - could not add vertex, received error %v", err)
		}
	}

	// Add edges
	// Poor graphical representation:
	// a -> b -> c
	// |    /
	// v   /
	// d ->
	vertexIDPairs := [][2]string{
		{"a", "b"},
		{"a", "d"},
		{"b", "c"},
		{"d", "b"},
	}

	for _, pair := range vertexIDPairs {
		if err := g.AddEdge(pair[0], pair[1], 42); err != nil {
			t.Errorf("Fail - could not add %v -> %v (%v). Got error %v",
				pair[0], pair[1], 42, err)
		}
	}

	edges := getGraphEdges(g)

	if count := edgeCount(edges); count != len(vertexIDPairs) {
		t.Errorf("Fail - edge len is expected to be %v but found %v",
			count,
			len(vertexIDPairs))
	}

	edgeVertices := vertexIDPairs

	for _, pair := range edgeVertices {
		if w := edges[pair[0]][pair[1]]; w != 42 {
			t.Errorf("Fail - edge %v %v -> %v not found",
				pair[0],
				pair[1],
				42)
		}
	}

}

func TestGraph_AddEdgeUndirected(t *testing.T) {
	g := NewGraph(false)

	vertices := []*Vertex{
		NewSimpleVertex("a"),
		NewSimpleVertex("b"),
		NewSimpleVertex("c"),
		NewSimpleVertex("d"),
	}

	// Add vertices
	for _, v := range vertices {
		if err := g.AddVertex(v); err != nil {
			t.Errorf("Fail - could not add vertex, received error %v", err)
		}
	}

	// Add edges
	// Poor graphical representation:
	// a --- b --- c
	// |    /
	// d __|
	vertexIDPairs := [][2]string{
		{"a", "b"},
		{"a", "d"},
		{"b", "c"},
		{"d", "b"},
	}

	for _, pair := range vertexIDPairs {
		if err := g.AddEdge(pair[0], pair[1], 42); err != nil {
			t.Errorf("Fail - could not add %v -> %v (%v). Got error %v",
				pair[0], pair[1], 42, err)
		}
	}

	edges := getGraphEdges(g)

	if count := edgeCount(edges); count != 2*len(vertexIDPairs) {
		t.Errorf("Fail - edge len is expected to be %v but found %v",
			count,
			2*len(vertexIDPairs))
	}

	edgeVertices := [][2]string{
		{"a", "b"},
		{"b", "a"},
		{"a", "d"},
		{"d", "a"},
		{"b", "c"},
		{"c", "b"},
		{"d", "b"},
		{"b", "d"},
	}

	for _, pair := range edgeVertices {
		if w := edges[pair[0]][pair[1]]; w != 42 {
			t.Errorf("Fail - edge %v %v -> %v not found",
				pair[0],
				pair[1],
				42)
		}
	}
}

func getGraphEdges(g *Graph) map[string]map[string]int {
	edges := map[string]map[string]int{}
	for _, v := range g.vertices {
		edges[v.id] = map[string]int{}
		for w, weight := range v.outgoing {
			edges[v.id][w.id] = weight
		}
	}

	return edges
}

func edgeCount(edges map[string]map[string]int) (count int) {
	for _, m := range edges {
		count += len(m)
	}

	return count
}
