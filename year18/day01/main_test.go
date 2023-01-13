package day01

import (
	"fmt"
	"reflect"
	"testing"
)

func TestParse(t *testing.T) {
	examples := []struct {
		given  string
		expect []int
	}{
		{"+1\n+1\n+1", []int{1, 1, 1}},
		{"+1\n+1\n-2", []int{1, 1, -2}},
		{"-1\n-2\n-3", []int{-1, -2, -3}},
	}

	for _, example := range examples {
		name := fmt.Sprintf("Testing for input %s", example.given)
		t.Run(name, func(t *testing.T) {
			if got := parse(example.given); !reflect.DeepEqual(got, example.expect) {
				t.Errorf("Fail - expected %v, got %v", example.expect, got)
			}
		})
	}
}

func TestSolvePart1(t *testing.T) {
	examples := []struct {
		given  []int
		expect int
	}{
		{[]int{+1, +1, +1}, 3},
		{[]int{+1, +1, -2}, 0},
		{[]int{-1, -2, -3}, -6},
	}

	for _, example := range examples {
		name := fmt.Sprintf("Testing for input %v", example.given)
		t.Run(name, func(t *testing.T) {
			if got := solvePart1(example.given); got != example.expect {
				t.Errorf("Fail - expected %v, got %v", example.expect, got)
			}
		})
	}
}

func TestSolvePart2(t *testing.T) {
	examples := []struct {
		given  []int
		expect int
	}{
		{[]int{+1, -1}, 0},
		{[]int{+3, +3, +4, -2, -4}, 10},
		{[]int{-6, +3, +8, +5, -6}, 5},
		{[]int{+7, +7, -2, -7, -4}, 14},
	}

	for _, example := range examples {
		name := fmt.Sprintf("Testing for input %v", example.given)
		t.Run(name, func(t *testing.T) {
			if got := solvePart2(example.given); got != example.expect {
				t.Errorf("Fail - expected %v, got %v", example.expect, got)
			}
		})
	}
}

func TestSolve(t *testing.T) {
	solve1, solve2 := solve(input)
	expected1, expected2 := 590, 83445

	if solve1 != expected1 {
		t.Errorf("Fail - part 1. Expected %d, got %d", expected1, solve1)
	}

	if solve2 != expected2 {
		t.Errorf("Fail - part 2. Expected %d, got %d", expected2, solve2)
	}
}

func BenchmarkSolvePart1(b *testing.B) {
	frequencies := parse(input)
	for i := 0; i < b.N; i++ {
		solvePart1(frequencies)
	}
}

func BenchmarkSolvePart2(b *testing.B) {
	frequencies := parse(input)
	for i := 0; i < b.N; i++ {
		solvePart2(frequencies)
	}
}
