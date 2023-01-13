package algo

type Stack[T any] []T

// IsEmpty returns true if the stack has nothing
func (stack *Stack[T]) IsEmpty() bool {
	return len(*stack) == 0
}

// Push pushes value on top of stack
func (stack *Stack[T]) Push(value T) {
	*stack = append(*stack, value)
}

// Pop pops out the top value of the stack
func (stack *Stack[T]) Pop() (value T, empty bool) {
	if stack.IsEmpty() {
		return value, true
	}
	target := len(*stack) - 1
	value = (*stack)[target]
	*stack = (*stack)[:target]
	return value, false
}

// Peek shows the top value of the stack
func (stack *Stack[T]) Peek() (value T, empty bool) {
	if stack.IsEmpty() {
		return value, true
	}
	target := len(*stack) - 1
	value = (*stack)[target]
	return value, false
}
