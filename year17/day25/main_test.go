package day25

import (
	"testing"
)

func TestSolve(t *testing.T) {
	solve1 := solve(input)
	expected1 := 4287

	if solve1 != expected1 {
		t.Errorf("Fail - part 1. Expected %d, got %d", expected1, solve1)
	}
}

func BenchmarkSolve(b *testing.B) {
	for i := 0; i < b.N; i++ {
		solve(input)
	}
}
