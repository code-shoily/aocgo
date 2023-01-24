package algo

import (
	"container/heap"
	"testing"
)

func TestMinHeap(t *testing.T) {
	h := &Heap[int]{}
	heap.Init(h)
	heap.Push(h, 3)
	heap.Push(h, 5)
	heap.Push(h, 1)

	if min := h.Min(); min != 1 {
		t.Errorf("Fail - expected %v got %v", 1, min)
	}

	if min := heap.Pop(h); min != 1 {
		t.Errorf("Fail - expected %v got %v", 1, min)
	}

	if min := h.Min(); min != 3 {
		t.Errorf("Fail - expected %v got %v", 3, min)
	}

	if min := heap.Pop(h); min != 3 {
		t.Errorf("Fail - expected %v got %v", 3, min)
	}

	if min := h.Min(); min != 5 {
		t.Errorf("Fail - expected %v got %v", 5, min)
	}

	if min := heap.Pop(h); min != 5 {
		t.Errorf("Fail - expected %v got %v", 5, min)
	}

	if empty := h.IsEmpty(); !empty {
		t.Errorf("Fail - wrong size, expected %v but got %v", true, empty)
	}
}

func TestMaxHeap(t *testing.T) {
	h := &MaxHeap[int]{}
	heap.Init(h)
	heap.Push(h, 1)
	heap.Push(h, 3)
	heap.Push(h, 5)

	if min := h.Max(); min != 5 {
		t.Errorf("Fail - expected %v got %v", 5, min)
	}

	if min := heap.Pop(h); min != 5 {
		t.Errorf("Fail - expected %v got %v", 5, min)
	}

	if min := h.Max(); min != 3 {
		t.Errorf("Fail - expected %v got %v", 3, min)
	}

	if min := heap.Pop(h); min != 3 {
		t.Errorf("Fail - expected %v got %v", 3, min)
	}

	if min := h.Max(); min != 1 {
		t.Errorf("Fail - expected %v got %v", 1, min)
	}

	if min := heap.Pop(h); min != 1 {
		t.Errorf("Fail - expected %v got %v", 1, min)
	}

	if empty := h.IsEmpty(); !empty {
		t.Errorf("Fail - wrong size, expected %v but got %v", true, empty)
	}
}

func TestHeapStringSortAscending(t *testing.T) {
	h := &Heap[string]{}
	heap.Init(h)

	heap.Push(h, "x")
	heap.Push(h, "m")
	heap.Push(h, "a")
	heap.Push(h, "e")

	var sorted [4]string
	expect := [...]string{"a", "e", "m", "x"}

	for i := 0; i != 4; i++ {
		sorted[i] = heap.Pop(h).(string)
	}

	if sorted != expect {
		t.Errorf("Fail - expected %v but got %v", expect, sorted)
	}
}

func TestHeapStringSortDescending(t *testing.T) {
	h := &MaxHeap[string]{}
	heap.Init(h)

	heap.Push(h, "x")
	heap.Push(h, "m")
	heap.Push(h, "a")
	heap.Push(h, "e")

	var sorted [4]string
	expect := [...]string{"x", "m", "e", "a"}

	for i := 0; i != 4; i++ {
		sorted[i] = heap.Pop(h).(string)
	}

	if sorted != expect {
		t.Errorf("Fail - expected %v but got %v", expect, sorted)
	}
}
