package day10

import (
	"testing"
)

func TestSolve(t *testing.T) {
	solve1, solve2 := solve(input)
	expected1, expected2 := 360_154, 5_103_798

	if solve1 != expected1 {
		t.Errorf("Fail - part 1. Expected %d, got %d", expected1, solve1)
	}

	if solve2 != expected2 {
		t.Errorf("Fail - part 2. Expected %d, got %d", expected2, solve2)
	}
}

func BenchmarkSolve(b *testing.B) {
	for i := 0; i < b.N; i++ {
		solve(input)
	}
}
