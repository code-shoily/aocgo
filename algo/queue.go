package algo

import "container/list"

// NewQueue creates a new Queue
func NewQueue() *Queue {
	return &Queue{list.New()}
}

type Queue struct {
	*list.List
}

// Enqueue inserts a new element on the queue
func (q *Queue) Enqueue(elem any) {
	q.PushBack(elem)
}

// Dequeue removes and returns the oldest element from the queue
func (q *Queue) Dequeue() (elem any, empty bool) {
	if q.IsEmpty() {
		return elem, true
	}
	front := q.Front()
	q.Remove(front)

	return front.Value, false
}

// MustDequeue removes and returns the oldest element from the queue, panics if empty
func (q *Queue) MustDequeue() any {
	if q.IsEmpty() {
		panic("empty queue")
	}
	front := q.Front()
	q.Remove(front)

	return front.Value
}

// Peek shows the oldest element in the queue without removing anything
func (q *Queue) Peek() (elem any, empty bool) {
	if q.IsEmpty() {
		return elem, true
	}
	return q.Front().Value, false
}

// IsEmpty checks if the queue is empty
func (q *Queue) IsEmpty() bool {
	return q.Len() == 0
}
