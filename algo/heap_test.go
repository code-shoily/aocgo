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
