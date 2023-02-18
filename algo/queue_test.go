package algo

import "testing"

func TestQueue_IsEmpty(t *testing.T) {
	queue := NewQueue()

	if empty := queue.IsEmpty(); !empty {
		t.Errorf("Fail - expected the queue to be empty")
	}

	queue.Enqueue("algorithms")

	if empty := queue.IsEmpty(); empty {
		t.Errorf("Fail - expected the queue to be non-empty")
	}

	queue.Dequeue()

	if empty := queue.IsEmpty(); !empty {
		t.Errorf("Fail - expected the queue to be empty")
	}
}

func TestQueue_EnqueueDeque(t *testing.T) {
	queue := NewQueue()
	queue.Enqueue(1)
	queue.Enqueue(2)
	queue.Enqueue(3)

	if e, notEmpty := queue.Dequeue(); e != 1 || notEmpty {
		t.Errorf("Fail - %v/%v expected, got %v/%v", 1, false, e, notEmpty)
	}
	if e, notEmpty := queue.Dequeue(); e != 2 || notEmpty {
		t.Errorf("Fail - %v/%v expected, got %v/%v", 1, false, e, notEmpty)
	}
	if e, notEmpty := queue.Dequeue(); e != 3 || notEmpty {
		t.Errorf("Fail - %v/%v expected, got %v/%v", 1, false, e, notEmpty)
	}
	if e, empty := queue.Dequeue(); e != nil || !empty {
		t.Errorf("Fail - %v/%v expected, got %v/%v", 1, true, e, empty)
	}
}

func TestQueue_MustDeque(t *testing.T) {
	queue := NewQueue()
	queue.Enqueue(1)
	queue.Enqueue(2)
	queue.Enqueue(3)

	if e := queue.MustDequeue(); e != 1 {
		t.Errorf("Fail - %v expected, got %v", 1, e)
	}
	if e := queue.MustDequeue(); e != 2 {
		t.Errorf("Fail - %v expected, got %v", 1, e)
	}
	if e := queue.MustDequeue(); e != 3 {
		t.Errorf("Fail - %v expected, got %v", 1, e)
	}
}

func TestQueue_Peek(t *testing.T) {
	queue := NewQueue()
	queue.Enqueue(1)
	if e, notEmpty := queue.Peek(); e != 1 || notEmpty {
		t.Errorf("Fail - %v/%v expected, got %v/%v", 1, false, e, notEmpty)
	}
	queue.Dequeue()
	//
	queue.Enqueue(2)
	if e, notEmpty := queue.Peek(); e != 2 || notEmpty {
		t.Errorf("Fail - %v/%v expected, got %v/%v", 1, false, e, notEmpty)
	}

	queue.Enqueue(3)
	if e, notEmpty := queue.Peek(); e != 2 || notEmpty {
		t.Errorf("Fail - %v/%v expected, got %v/%v", 1, false, e, notEmpty)
	}

	queue.Dequeue()
	queue.Dequeue()

	if e, empty := queue.Peek(); e != nil || !empty {
		t.Errorf("Fail - %v/%v expected, got %v/%v", 1, true, e, empty)
	}
}
