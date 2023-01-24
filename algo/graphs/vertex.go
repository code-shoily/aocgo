package graphs

type Connections[T any] map[*Vertex[T]]int

type Vertex[T any] struct {
	id       string
	content  T
	incoming Connections[T]
	outgoing Connections[T]
}

// ID returns the id for this vertex
func (v *Vertex[T]) ID() string {
	return v.id
}

// Content returns the T content of this vertex
func (v *Vertex[T]) Content() T {
	return v.content
}

// AddConnection creates a connection between 2 vertices. If reciprocal is true, then a doubly linked connection is formed
func (v *Vertex[T]) AddConnection(w *Vertex[T], weight int, reciprocal bool) {
	v.outgoing[w] = weight
	w.incoming[v] = weight

	if reciprocal {
		w.AddConnection(v, weight, false)
	}
}

// GetConnections returns the outgoing and incoming connections
func (v *Vertex[T]) GetConnections() (outgoing Connections[T], incoming Connections[T]) {
	return v.outgoing, v.incoming
}

// String representation of this vertex
func (v *Vertex[T]) String() string {
	return v.id
}

// NewVertex creates a new vertex with id and content
func NewVertex[T any](id string, content T) *Vertex[T] {
	return &Vertex[T]{
		id,
		content,
		make(map[*Vertex[T]]int),
		make(map[*Vertex[T]]int),
	}
}

// NewSimpleVertex creates a vertex where id and value are one and the same
func NewSimpleVertex(value string) *Vertex[string] {
	return NewVertex[string](value, value)
}
