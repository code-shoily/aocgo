package algo

import (
	"reflect"
	"testing"
)

func TestStack_IsEmpty(t *testing.T) {
	var stack Stack[string]

	if !stack.IsEmpty() {
		t.Errorf("Fail - expected %v to be empty. It's not", stack)
	}

	stack.Push("A")

	if stack.IsEmpty() {
		t.Errorf("Fail - expected %v to not be empty. It is", stack)
	}

	stack.Pop()

	if !stack.IsEmpty() {
		t.Errorf("Fail - expected %v to be empty. It's not", stack)
	}
}

func TestStack_Push(t *testing.T) {
	var stack Stack[int]

	stack.Push(1)
	stack.Push(2)
	stack.Push(3)

	expect := []int{1, 2, 3}

	if !reflect.DeepEqual([]int(stack), expect) {
		t.Errorf("Fail - expected %v, got %v", expect, stack)
	}
}

func TestStack_Peek(t *testing.T) {
	var stack Stack[int]

	if _, err := stack.Peek(); !err {
		t.Errorf("Fail - expected error to be true")
	}

	stack.Push(1)
	stack.Push(2)
	stack.Push(3)

	expect := 3

	if value, err := stack.Peek(); value != expect || err {
		t.Errorf("Fail - expected %v, got %v", expect, value)
	}
}

func TestStack_Pop(t *testing.T) {
	var stack Stack[int]

	if _, err := stack.Pop(); !err {
		t.Errorf("Fail - expected error to be true")
	}

	stack.Push(1)
	stack.Push(2)
	stack.Push(3)

	expect, content := 3, []int{1, 2}

	if value, err := stack.Pop(); value != expect || err {
		t.Errorf("Fail - expected %v, got %v", expect, value)
	}

	if !reflect.DeepEqual([]int(stack), content) {
		t.Errorf("Fail - expected %v, got %v", content, stack)
	}
}
