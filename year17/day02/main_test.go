package day02

import "testing"

func TestSolve(t *testing.T) {
	solve1, solve2 := solve(input)
	expect1, expect2 := 32020, 236

	if solve1 != expect1 {
		t.Errorf("Fail - part 1. Expected %d, got %d", expect1, solve1)
	}

	if solve2 != expect2 {
		t.Errorf("Fail - part 2. Expected %d, got %d", expect2, solve2)
	}
}

func BenchmarkSolveP(b *testing.B) {
	for i := 0; i < b.N; i++ {
		solve(input)
	}
}
