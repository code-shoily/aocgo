package day06

import (
	"testing"
)

func TestSolve(t *testing.T) {
	solve1, solve2 := solve(input)
	expected1, expected2 := "qzedlxso", "ucmifjae"

	if solve1 != expected1 || solve2 != expected2 {
		t.Errorf("Fail - expected %v/%v but got %v/%v", expected1, expected2, solve1, solve2)
	}
}

func BenchmarkSolve(b *testing.B) {
	for i := 0; i < b.N; i++ {
		solve(input)
	}
}
