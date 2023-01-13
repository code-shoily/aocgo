package day02

import (
	"fmt"
	"testing"
)

func TestNewCommand(t *testing.T) {
	examples := []struct {
		given  string
		expect Command
	}{
		{"forward 5", Command{"forward", 5}},
		{"up 15", Command{"up", 15}},
		{"down 0", Command{"down", 0}},
	}

	for _, example := range examples {
		name := fmt.Sprintf("testing for input %v", example.given)
		t.Run(name, func(t *testing.T) {
			if got := newCommand(example.given); got != example.expect {
				t.Errorf("Fail - expected %#v, got %#v", example.expect, got)
			}
		})
	}
}

func TestSolve(t *testing.T) {
	solve1, solve2 := solve(input)
	expected1, expected2 := 1_660_158, 1_604_592_846

	if solve1 != expected1 {
		t.Errorf("Fail - part 1. Expected %d, got %d", expected1, solve1)
	}

	if solve2 != expected2 {
		t.Errorf("Fail - part 2. Expected %d, got %d", expected2, solve2)
	}
}

func BenchmarkSolvePart1(b *testing.B) {
	data := parse(input)
	for i := 0; i < b.N; i++ {
		solvePart1(data)
	}
}

func BenchmarkSolvePart2(b *testing.B) {
	data := parse(input)
	for i := 0; i < b.N; i++ {
		solvePart2(data)
	}
}
