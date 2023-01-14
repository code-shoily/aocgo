package day03

import (
	"testing"
)

func TestSolve(t *testing.T) {
	solve1, solve2 := solve(input)
	expected1, expected2 := 110_389, 552

	if solve1 != expected1 {
		t.Errorf("Fail - part 1. Expected %d, got %d", expected1, solve1)
	}

	if solve2 != expected2 {
		t.Errorf("Fail - part 2. Expected %d, got %d", expected2, solve2)
	}
}

func BenchmarkSolvePart1(b *testing.B) {
	_, tally := parse(input)
	for i := 0; i < b.N; i++ {
		solvePart1(tally)
	}
}

func BenchmarkSolvePart2(b *testing.B) {
	claims, tally := parse(input)
	for i := 0; i < b.N; i++ {
		solvePart2(claims, tally)
	}
}

func BenchmarkSolve(b *testing.B) {
	for i := 0; i < b.N; i++ {
		solve(input)
	}
}
