package graphs

import (
	"reflect"
	"testing"
)

func TestNewVertex(t *testing.T) {
	v := NewVertex("a", 'a')

	if vID := v.ID(); vID != "a" {
		t.Errorf("Fail - expected %v but got %v", "a", vID)
	}
}

func TestNewSimpleVertex(t *testing.T) {
	v := NewSimpleVertex("a")

	if vID := v.ID(); vID != "a" {
		t.Errorf("Fail - expected %v but got %v", "a", vID)
	}
}

func TestVertex_AddConnectionUnidirectional(t *testing.T) {
	v := NewSimpleVertex("a")
	w := NewSimpleVertex("b")
	v.AddConnection(w, 10, false)

	expectedVOutgoing := Connections[string]{
		w: 10,
	}

	expectedVIncoming := Connections[string]{}

	expectedWOutgoing := Connections[string]{}

	expectedWIncoming := Connections[string]{
		v: 10,
	}

	vOutgoing, vIncoming := v.GetConnections()
	wOutgoing, wIncoming := w.GetConnections()

	if !reflect.DeepEqual(expectedVOutgoing, vOutgoing) {
		t.Errorf("Fail v-> - expected outgoing to be %v but got %v", expectedVOutgoing, vOutgoing)
	}

	if !reflect.DeepEqual(expectedVIncoming, vIncoming) {
		t.Errorf("Fail v<- - expected outgoing to be %v but got %v", expectedVIncoming, vIncoming)
	}

	if !reflect.DeepEqual(expectedWOutgoing, wOutgoing) {
		t.Errorf("Fail w-> - expected outgoing to be %v but got %v", expectedWOutgoing, wOutgoing)
	}

	if !reflect.DeepEqual(expectedWIncoming, wIncoming) {
		t.Errorf("Fail w<- - expected outgoing to be %v but got %v", expectedWIncoming, wIncoming)
	}
}

func TestVertex_AddConnectionReciprocal(t *testing.T) {
	v := NewSimpleVertex("a")
	w := NewSimpleVertex("b")
	v.AddConnection(w, 10, true)

	expectedVOutgoing := Connections[string]{
		w: 10,
	}

	expectedVIncoming := Connections[string]{
		w: 10,
	}

	expectedWOutgoing := Connections[string]{
		v: 10,
	}

	expectedWIncoming := Connections[string]{
		v: 10,
	}

	vOutgoing, vIncoming := v.GetConnections()
	wOutgoing, wIncoming := w.GetConnections()

	if !reflect.DeepEqual(expectedVOutgoing, vOutgoing) {
		t.Errorf("Fail v-> - expected outgoing to be %v but got %v", expectedVOutgoing, vOutgoing)
	}

	if !reflect.DeepEqual(expectedVIncoming, vIncoming) {
		t.Errorf("Fail v<- - expected outgoing to be %v but got %v", expectedVIncoming, vIncoming)
	}

	if !reflect.DeepEqual(expectedWOutgoing, wOutgoing) {
		t.Errorf("Fail w-> - expected outgoing to be %v but got %v", expectedWOutgoing, wOutgoing)
	}

	if !reflect.DeepEqual(expectedWIncoming, wIncoming) {
		t.Errorf("Fail w<- - expected outgoing to be %v but got %v", expectedWIncoming, wIncoming)
	}
}
