package graphs

type connections[T any] map[*vertex[T]]int

type vertex[T any] struct {
	id       string
	content  T
	incoming connections[T]
	outgoing connections[T]
}

// ID returns the id for this vertex
func (v *vertex[T]) ID() string {
	return v.id
}

// Content returns the T content of this vertex
func (v *vertex[T]) Content() T {
	return v.content
}

// AddConnection creates a connection between 2 vertices. If reciprocal is true, then a doubly linked connection is formed
func (v *vertex[T]) AddConnection(w *vertex[T], weight int, reciprocal bool) {
	v.outgoing[w] = weight
	w.incoming[v] = weight

	if reciprocal {
		w.AddConnection(v, weight, false)
	}
}

// Connections returns the outgoing and incoming connections
func (v *vertex[T]) Connections() (outgoing connections[T], incoming connections[T]) {
	return v.outgoing, v.incoming
}

// String representation of this vertex
func (v *vertex[T]) String() string {
	return v.id
}

// NewVertex creates a new vertex with id and content
func NewVertex[T any](id string, content T) *vertex[T] {
	return &vertex[T]{
		id,
		content,
		make(map[*vertex[T]]int),
		make(map[*vertex[T]]int),
	}
}

// NewSimpleVertex creates a vertex where id and value are one and the same
func NewSimpleVertex(value string) *vertex[string] {
	return NewVertex[string](value, value)
}
