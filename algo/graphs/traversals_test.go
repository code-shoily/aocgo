package graphs

import (
	"fmt"
	"reflect"
	"testing"
)

func setUpDirectedGraph() *Graph {
	graph := NewGraph(true)

	for _, name := range []string{"A", "B", "C", "D", "E", "F"} {
		v := NewSimpleVertex(name)
		graph.AddVertex(v)
	}

	/*
	          C -7-> D
	          |      ^
	          3      1
	          v      |
	   A -2-> B <-8- E
	   |            /
	   --7--> F -2->
	*/
	graph.AddEdge("A", "B", 2)
	graph.AddEdge("A", "F", 7)
	graph.AddEdge("C", "B", 3)
	graph.AddEdge("C", "D", 7)
	graph.AddEdge("E", "D", 1)
	graph.AddEdge("E", "B", 8)
	graph.AddEdge("F", "E", 2)

	return graph
}

func TestDFS(t *testing.T) {
	graph := setUpDirectedGraph()

	examples := []struct {
		source string
		expect PathMap
	}{
		{"A", PathMap{
			"B": {"B"},
			"D": {"F", "E", "D"},
			"E": {"F", "E"},
			"F": {"F"},
		}},
		{"B", PathMap{}},
		{"C", PathMap{
			"B": {"B"},
			"D": {"D"},
		}},
		{"D", PathMap{}},
		{"E", PathMap{
			"B": {"B"},
			"D": {"D"},
		}},
		{"F", PathMap{
			"B": {"E", "B"},
			"D": {"E", "D"},
			"E": {"E"},
		}},
	}

	for _, example := range examples {
		name := fmt.Sprintf("testing for input %v", example.source)
		t.Run(name, func(tt *testing.T) {
			if got := DFS(graph, example.source); !reflect.DeepEqual(got, example.expect) {
				tt.Errorf("Fail - expected %v but got %v", example.expect, got)
			}
		})
	}

	fmt.Println(DFS(graph, "B"))
}
