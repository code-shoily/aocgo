package day02

import (
	"reflect"
	"testing"
)

func TestParse(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		expected [][]int
	}{
		{
			name:  "simple multi-line input",
			input: "1 2 3\n7 8\n5 6 7 8",
			expected: [][]int{
				{1, 2, 3},
				{7, 8},
				{5, 6, 7, 8},
			},
		},
		{
			name:     "single line",
			input:    "12 45 67",
			expected: [][]int{{12, 45, 67}},
		},
		{
			name:     "variable spacing",
			input:    "1  2   3\n  4    5     6  ",
			expected: [][]int{{1, 2, 3}, {4, 5, 6}},
		},
		{
			name:     "single numbers per line",
			input:    "1\n2\n3",
			expected: [][]int{{1}, {2}, {3}},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := parse(tc.input)
			if !reflect.DeepEqual(result, tc.expected) {
				t.Errorf("Expected %v, got %v", tc.expected, result)
			}
		})
	}
}

func TestSolve(t *testing.T) {
	solve1, solve2 := solve(input)
	expected1, expected2 := 486, 540

	if solve1 != expected1 {
		t.Errorf("Fail - part 1. Expected %v, got %v", expected1, solve1)
	}

	if solve2 != expected2 {
		t.Errorf("Fail - part 2. Expected %v, got %v", expected2, solve2)
	}
}

func BenchmarkSolve(b *testing.B) {
	for i := 0; i < b.N; i++ {
		solve(input)
	}
}

/* func BenchmarkSolvePart1(b *testing.B) {
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
} */
