package day01

import (
	"testing"
)

func TestSolve(t *testing.T) {
	solve1, solve2 := solve(input)
	expected1, expected2 := 1_014_624, 80_072_256

	if solve1 != expected1 {
		t.Errorf("Fail - part 1. Expected %d, got %d", expected1, solve1)
	}

	if solve2 != expected2 {
		t.Errorf("Fail - part 2. Expected %d, got %d", expected2, solve2)
	}
}

func BenchmarkSolvePart1(b *testing.B) {
	masses := parse(input)
	for i := 0; i < b.N; i++ {
		solvePart1(masses)
	}
}

func BenchmarkSolvePart2(b *testing.B) {
	masses := parse(input)
	for i := 0; i < b.N; i++ {
		solvePart2(masses)
	}
}
