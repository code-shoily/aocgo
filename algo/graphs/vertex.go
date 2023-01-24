package graphs

type Connections map[*Vertex]int

type Vertex struct {
	id       string
	content  any
	incoming Connections
	outgoing Connections
}

// ID returns the id for this vertex
func (v *Vertex) ID() string {
	return v.id
}

// Content returns the T content of this vertex
func (v *Vertex) Content() any {
	return v.content
}

// AddConnection creates a connection between 2 vertices. If reciprocal is true, then a doubly linked connection is formed
func (v *Vertex) AddConnection(w *Vertex, weight int, reciprocal bool) {
	v.outgoing[w] = weight
	w.incoming[v] = weight

	if reciprocal {
		w.AddConnection(v, weight, false)
	}
}

// GetConnections returns the outgoing and incoming connections
func (v *Vertex) GetConnections() (outgoing Connections, incoming Connections) {
	return v.outgoing, v.incoming
}

// String representation of this vertex
func (v *Vertex) String() string {
	return v.id
}

// NewVertex creates a new vertex with id and content
func NewVertex(id string, content any) *Vertex {
	return &Vertex{
		id,
		content,
		make(map[*Vertex]int),
		make(map[*Vertex]int),
	}
}

// NewSimpleVertex creates a vertex where id and value are one and the same
func NewSimpleVertex(value string) *Vertex {
	return NewVertex(value, value)
}
