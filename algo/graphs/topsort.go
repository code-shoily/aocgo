package graphs

import (
	"container/heap"
	"github.com/code-shoily/aocgo/algo"
)

// TopologicalSort performs a topological sort based on FIFO and returns the sorted order (nil if cycle) and cycle
func TopologicalSort(g *Graph) (sorted []string, cycle bool) {
	var count int
	inDegrees := make(map[string]int)
	queue := algo.NewQueue()

	for _, vertex := range g.Vertices() {
		if inDegree := len(vertex.incoming); inDegree == 0 {
			queue.Enqueue(vertex.ID())
		} else {
			inDegrees[vertex.ID()] += inDegree
		}
	}

	for !queue.IsEmpty() {
		if vertexAny, empty := queue.Dequeue(); !empty {
			vertexID := vertexAny.(string)
			sorted = append(sorted, vertexID)

			for outgoing, _ := range g.Vertices()[vertexID].outgoing {
				outgoingID := outgoing.ID()
				inDegrees[outgoingID]--

				if inDegrees[outgoingID] == 0 {
					queue.Enqueue(outgoingID)
				}
			}
		}
		count++
	}

	if count != len(g.Vertices()) {
		return nil, true
	}

	return sorted, cycle
}

func lexicographicalSorter(g *Graph, h heap.Interface) (sorted []string, cycle bool) {
	var count int
	inDegrees := make(map[string]int)
	heap.Init(h)

	for _, vertex := range g.Vertices() {
		if inDegree := len(vertex.incoming); inDegree == 0 {
			heap.Push(h, vertex.ID())
		} else {
			inDegrees[vertex.ID()] += inDegree
		}
	}

	for h.Len() != 0 {
		vertexAny := heap.Pop(h)
		vertexID := vertexAny.(string)
		sorted = append(sorted, vertexID)

		for outgoing, _ := range g.Vertices()[vertexID].outgoing {
			outgoingID := outgoing.ID()
			inDegrees[outgoingID]--

			if inDegrees[outgoingID] == 0 {
				heap.Push(h, outgoingID)
			}
		}

		count++
	}

	if count != len(g.Vertices()) {
		return nil, true
	}

	return sorted, cycle
}

// LexicographicalTopologicalSort returns sort in ascending order of keys unless a cycle is detected.
func LexicographicalTopologicalSort(g *Graph) (sorted []string, cycle bool) {
	return lexicographicalSorter(g, &algo.Heap[string]{})
}

// LexicographicalReverseTopologicalSort returns sort in descending order of keys unless a cycle is detected.
func LexicographicalReverseTopologicalSort(g *Graph) (sorted []string, cycle bool) {
	return lexicographicalSorter(g, &algo.MaxHeap[string]{})
}
