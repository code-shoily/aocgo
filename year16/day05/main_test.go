package day05

import (
	"testing"
)

func TestSolve(t *testing.T) {
	solve1, solve2 := solve(input)
	expected1, expected2 := "f77a0e6e", "999828ec"

	if solve1 != expected1 {
		t.Errorf("Fail - part 1. Expected %s, got %s", expected1, solve1)
	}

	if solve2 != expected2 {
		t.Errorf("Fail - part 2. Expected %s, got %s", expected2, solve2)
	}
}
