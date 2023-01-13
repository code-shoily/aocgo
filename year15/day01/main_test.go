package day01

import (
	"fmt"
	"testing"
)

func TestSolvePart1(t *testing.T) {
	var examples = []struct {
		given  []byte
		expect int
	}{
		{[]byte{'(', '(', ')', ')'}, 0},
		{[]byte{'(', ')', '(', ')'}, 0},
		{[]byte{'(', '(', '('}, 3},
		{[]byte{'(', '(', ')', '(', '(', ')', '('}, 3},
		{[]byte{')', ')', '(', '(', '(', '(', '('}, 3},
		{[]byte{'(', ')', ')'}, -1},
		{[]byte{')', ')', '('}, -1},
		{[]byte{')', ')', ')'}, -3},
		{[]byte{')', '(', ')', ')', '(', ')', ')'}, -3},
	}

	for _, tt := range examples {
		testname := fmt.Sprintf("Part 1 for %s => %d", tt.given, tt.expect)
		t.Run(testname, func(t *testing.T) {
			solution := solvePart1(tt.given)

			if solution != tt.expect {
				t.Errorf("Fail - expected %s, expect %d", tt.given, tt.expect)
			}
		})
	}
}

func TestSolvePart2(t *testing.T) {
	var examples = []struct {
		given []byte
		got   int
	}{
		{[]byte{')'}, 1},
		{[]byte{'(', ')', '(', ')', ')'}, 5},
	}

	for _, tt := range examples {
		testname := fmt.Sprintf("Part 2 for %s => %d", tt.given, tt.got)
		t.Run(testname, func(t *testing.T) {
			solution := solvePart2(tt.given)

			if solution != tt.got {
				t.Errorf("Fail - expected %s, expect %d", tt.given, tt.got)
			}
		})
	}
}

func TestSolve(t *testing.T) {
	solve1, solve2 := solve()
	expect1, expect2 := 232, 1783

	if solve1 != expect1 {
		t.Errorf("Part 1 failed, expected %d, expect %d", expect1, solve1)
	}

	if solve2 != expect2 {
		t.Errorf("Part 1 failed, expected %d, expect %d", expect2, solve2)
	}
}

func BenchmarkSolvePart1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		solvePart1(input)
	}
}

func BenchmarkSolvePart2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		solvePart2(input)
	}
}
