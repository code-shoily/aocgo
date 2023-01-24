package algo

import "container/list"

// NewQueue creates a new Queue
func NewQueue[T any]() *Queue[T] {
	return &Queue[T]{list.New()}
}

type Queue[T any] struct {
	*list.List
}

// Enqueue insers a new element on the queue
func (q *Queue[T]) Enqueue(elem T) {
	q.PushBack(elem)
}

// Dequeue removes and returns the oldest element from the queue
func (q *Queue[T]) Dequeue() (elem T, empty bool) {
	if q.IsEmpty() {
		return elem, true
	}
	front := q.Front()
	q.Remove(front)

	return front.Value.(T), false
}

// Peek shows the oldest element in the queue without removing anything
func (q *Queue[T]) Peek() (elem T, empty bool) {
	if q.IsEmpty() {
		return elem, true
	}
	return q.Front().Value.(T), false
}

// IsEmpty checks if the queue is empty
func (q *Queue[T]) IsEmpty() bool {
	return q.Len() == 0
}
