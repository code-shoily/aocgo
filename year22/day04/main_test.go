package day04

import (
	"fmt"
	"testing"
)

func TestHasContained(t *testing.T) {
	examples := []struct {
		given  [4]int
		expect bool
	}{
		{[4]int{2, 4, 6, 8}, false},
		{[4]int{2, 3, 4, 5}, false},
		{[4]int{5, 7, 7, 9}, false},
		{[4]int{2, 8, 3, 7}, true},
		{[4]int{6, 6, 4, 6}, true},
		{[4]int{2, 6, 4, 8}, false},
	}

	for _, example := range examples {
		name := fmt.Sprintf("testing for %v", example.given)
		t.Run(name, func(tt *testing.T) {
			if got := hasContained(example.given); got != example.expect {
				tt.Errorf("Fail - expected %v but got %v", example.expect, got)
			}
		})
	}
}

func TestHasOverlap(t *testing.T) {
	examples := []struct {
		given  [4]int
		expect bool
	}{
		{[4]int{2, 4, 6, 8}, false},
		{[4]int{2, 3, 4, 5}, false},
		{[4]int{5, 7, 7, 9}, true},
		{[4]int{2, 8, 3, 7}, true},
		{[4]int{6, 6, 4, 6}, true},
		{[4]int{2, 6, 4, 8}, true},
	}

	for _, example := range examples {
		name := fmt.Sprintf("testing for %v", example.given)
		t.Run(name, func(tt *testing.T) {
			if got := hasOverlap(example.given); got != example.expect {
				tt.Errorf("Fail - expected %v but got %v", example.expect, got)
			}
		})
	}
}

func TestSolve(t *testing.T) {
	solve1, solve2 := solve(input)
	expected1, expected2 := 518, 909

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

func BenchmarkSolve(b *testing.B) {
	for i := 0; i < b.N; i++ {
		solve(input)
	}
}
