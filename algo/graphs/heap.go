package graphs

import (
	"golang.org/x/exp/constraints"
)

// Heap is implementation of MinHeap
type Heap[T constraints.Ordered] []T

// Len is the length of the heap.
func (h *Heap[T]) Len() int {
	return len(*h)
}

// Less is the comparator function used to determine order of the heap.
func (h *Heap[T]) Less(i, j int) bool {
	return (*h)[i] < (*h)[j]
}

// Swap is the method to swap two elements of this heap.
func (h *Heap[T]) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

// Push inserts an element to this heap.
func (h *Heap[T]) Push(elem any) {
	*h = append(*h, elem.(T))
}

// Pop extracts the minimum value of this heap
func (h *Heap[T]) Pop() any {
	old := *h
	size := len(old)
	elem := old[size-1]
	*h = old[0 : size-1]
	return elem
}

// Min returns the minimum value without removing anything.
func (h *Heap[T]) Min() any {
	return (*h)[0]
}

// IsEmpty checks if this heap is empty.
func (h *Heap[T]) IsEmpty() bool {
	return h.Len() == 0
}

// MaxHeap is implementation of MaxHeap.
type MaxHeap[T constraints.Ordered] []T

// Len is the length of the heap.
func (h *MaxHeap[T]) Len() int {
	return len(*h)
}

// Less is the comparator function used to determine order of the heap. (Here Less is More)
func (h *MaxHeap[T]) Less(i, j int) bool {
	return (*h)[i] > (*h)[j]
}

// Swap is the method to swap two elements of this heap.
func (h *MaxHeap[T]) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

// Push inserts an element to this heap.
func (h *MaxHeap[T]) Push(elem any) {
	*h = append(*h, elem.(T))
}

// Pop extracts the maximum value of this heap
func (h *MaxHeap[T]) Pop() any {
	old := *h
	size := len(old)
	elem := old[size-1]
	*h = old[0 : size-1]
	return elem
}

// Max returns the maximum value without removing anything.
func (h *MaxHeap[T]) Max() any {
	return (*h)[0]
}

// IsEmpty checks if this heap is empty.
func (h *MaxHeap[T]) IsEmpty() bool {
	return h.Len() == 0
}
