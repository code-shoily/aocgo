package day21

import (
	"testing"
)

func TestSwapPositions(t *testing.T) {
	example := "abcde"
	expect := "ebcda"
	if got := swapPositions(example, 4, 0); got != expect {
		t.Errorf("Fail - expected %s but got %s", expect, got)
	}
}

func TestSwapLetters(t *testing.T) {
	example := "ebcda"
	expect := "edcba"
	if got := swapLetters(example, "d", "b"); got != expect {
		t.Errorf("Fail - expected %s but got %s", expect, got)
	}
}

func TestReverse(t *testing.T) {
	example := "edcba"
	expect := "abcde"
	if got := reverse(example, 0, 4); got != expect {
		t.Errorf("Fail - expected %s but got %s", expect, got)
	}
}

func TestRotateLeftStep(t *testing.T) {
	example := "abcde"
	expect := "bcdea"
	if got := rotateLeft(example, 1); got != expect {
		t.Errorf("Fail - expected %s but got %s", expect, got)
	}
}

func TestRotateRightStep(t *testing.T) {
	example := "abcde"
	expect := "eabcd"
	if got := rotateRight(example, 1); got != expect {
		t.Errorf("Fail - expected %s but got %s", expect, got)
	}
}

func TestMove(t *testing.T) {
	examples := []struct {
		given    string
		expect   string
		from, to int
	}{
		{"bcdea", "bdeac", 1, 4},
		{"bdeac", "abdec", 3, 0},
	}

	for _, example := range examples {
		if got := move(example.given, example.from, example.to); got != example.expect {
			t.Errorf("Fail - expected %s but got %s", example.expect, got)
		}
	}

}

func TestRotateRelative(t *testing.T) {
	examples := []struct {
		given  string
		by     string
		expect string
	}{
		{"abdec", "b", "ecabd"},
		{"ecabd", "d", "decab"},
	}

	for _, example := range examples {
		if got := rotateRelative(example.given, example.by); got != example.expect {
			t.Errorf("Fail - expected %s but got %s", example.expect, got)
		}
	}

}

func TestSolve(t *testing.T) {
	solve1, solve2 := solve(scrambleRules)
	expected1, expected2 := "aefgbcdh", "egcdahbf"

	if solve1 != expected1 {
		t.Errorf("Fail - part 1. Expected %v, got %v", expected1, solve1)
	}

	if solve2 != expected2 {
		t.Errorf("Fail - part 2. Expected %v, got %v", expected2, solve2)
	}
}

func BenchmarkSolve(b *testing.B) {
	for i := 0; i < b.N; i++ {
		solve(scrambleRules)
	}
}
