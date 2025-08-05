package day01

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
			name:  "simple two column input",
			input: "1 2\n3 4\n5 6",
			expected: [][]int{
				{1, 3, 5}, // left column
				{2, 4, 6}, // right column
			},
		},
		{
			name:  "variable spacing",
			input: "10   20\n30    40\n  50     60  ",
			expected: [][]int{
				{10, 30, 50},
				{20, 40, 60},
			},
		},
		{
			name:  "with empty lines",
			input: "1 2\n\n3 4\n\n5 6\n",
			expected: [][]int{
				{1, 3, 5},
				{2, 4, 6},
			},
		},
		{
			name:  "single line",
			input: "100 200",
			expected: [][]int{
				{100},
				{200},
			},
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
	expected1, expected2 := 2742123, 21328497

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
