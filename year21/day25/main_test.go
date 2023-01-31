package day25

import (
	"testing"
)

func TestSolve(t *testing.T) {
	solution := solve(input)
	expected := 504

	if solution != expected {
		t.Errorf("Fail - part 1. Expected %d, got %d", expected, solution)
	}
}

func BenchmarkSolve(b *testing.B) {
	for i := 0; i < b.N; i++ {
		solve(input)
	}
}
