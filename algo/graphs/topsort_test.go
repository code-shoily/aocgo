package graphs

import (
	"reflect"
	"testing"
)

func MakeGraph() *Graph {
	g := NewGraph(true)
	for _, vertexID := range []string{"0", "1", "2", "3", "4", "5"} {
		g.AddVertex(NewSimpleVertex(vertexID))
	}
	g.AddEdge("2", "3", 1)
	g.AddEdge("3", "1", 1)
	g.AddEdge("4", "0", 1)
	g.AddEdge("4", "1", 1)
	g.AddEdge("5", "2", 1)
	g.AddEdge("5", "0", 1)

	return g
}

func MakeGraphWithCycle() *Graph {
	g := NewGraph(true)
	for _, vertexID := range []string{"0", "1", "2", "3", "4", "5"} {
		g.AddVertex(NewSimpleVertex(vertexID))
	}
	g.AddEdge("2", "3", 1)
	g.AddEdge("3", "1", 1)
	g.AddEdge("4", "0", 1)
	g.AddEdge("4", "1", 1)
	g.AddEdge("5", "2", 1)
	g.AddEdge("5", "0", 1)
	g.AddEdge("3", "2", 1)

	return g
}

func TestTopologicalSort(t *testing.T) {
	graph := MakeGraph()
	expect := []string{"4", "5", "2", "0", "3", "1"}
	if sorted, cycle := TopologicalSort(graph); !reflect.DeepEqual(sorted, expect) || cycle {
		t.Errorf("Fail - expected %v/%v but got %v/%v", expect, false, sorted, cycle)
	}
}

func TestTopologicalSortCycled(t *testing.T) {
	graph := MakeGraphWithCycle()
	if _, cycle := TopologicalSort(graph); !cycle {
		t.Errorf("Fail - expected %v but got %v", true, cycle)
	}
}

func TestLexicographicalTopologicalSort(t *testing.T) {
	graph := MakeGraph()
	expect := []string{"4", "5", "0", "2", "3", "1"}
	if sorted, cycle := LexicographicalTopologicalSort(graph); !reflect.DeepEqual(sorted, expect) || cycle {
		t.Errorf("Fail - expected %v/%v but got %v/%v", expect, false, sorted, cycle)
	}
}

func TestLexicographicalTopologicalSortCycled(t *testing.T) {
	graph := MakeGraphWithCycle()
	if _, cycle := LexicographicalTopologicalSort(graph); !cycle {
		t.Errorf("Fail - expected %v but got %v", true, cycle)
	}
}

func TestLexicographicalReverseTopologicalSort(t *testing.T) {
	graph := MakeGraph()
	expect := []string{"5", "4", "2", "3", "1", "0"}
	if sorted, cycle := LexicographicalReverseTopologicalSort(graph); !reflect.DeepEqual(sorted, expect) || cycle {
		t.Errorf("Fail - expected %v/%v but got %v/%v", expect, false, sorted, cycle)
	}
}

func TestLexicographicalReverseTopologicalSortCycled(t *testing.T) {
	graph := MakeGraphWithCycle()
	if _, cycle := LexicographicalReverseTopologicalSort(graph); !cycle {
		t.Errorf("Fail - expected %v but got %v", true, cycle)
	}
}
