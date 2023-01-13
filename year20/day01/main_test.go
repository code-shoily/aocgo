package day01

import (
	"fmt"
	"testing"
)

func TestTwoSum(t *testing.T) {
	examples := []struct {
		given  []int
		target int
		a      int
		b      int
		ok     bool
	}{
		{[]int{5, 8, 10, 10}, 20, 10, 10, true},
		{[]int{1, 2, 3, 4}, 6, 2, 4, true},
		{[]int{3, 8, 5, 16}, 16, -1, -1, false},
	}

	for _, example := range examples {
		name := fmt.Sprintf("Testing for input %v", example.given)
		t.Run(name, func(t *testing.T) {
			a, b, ok := TwoSum(example.given, example.target)
			if a != example.a || b != example.b || ok != example.ok {
				t.Errorf("Fail - expected %v for %d and %d but got %v for %d and %d",
					example.ok, example.a, example.b, ok, a, b)
			}
		})
	}
}

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
