package day11

import (
	"fmt"
	"testing"
)

func TestNextPassword(t *testing.T) {
	examples := []struct {
		given    string
		expected string
	}{
		{"abxx", "abxy"},
		{"xy", "xz"},
		{"xz", "ya"},
		{"ya", "yb"},
		{"zzz", "aaa"},
	}

	for _, example := range examples {
		name := fmt.Sprintf("testing for data %v", example.given)

		t.Run(name, func(tt *testing.T) {
			if got := nextString([]rune(example.given)); string(got) != example.expected {
				tt.Errorf("Fail - expected %v but got %v", example.expected, string(got))
			}
		})
	}

}
func TestNext(t *testing.T) {
	examples := []struct {
		given    rune
		expected rune
		carry    bool
	}{
		{'a', 'b', false},
		{'c', 'd', false},
		{'e', 'f', false},
		{'z', 'a', true},
	}

	for _, example := range examples {
		name := fmt.Sprintf("testing for data %v", example.given)

		t.Run(name, func(tt *testing.T) {
			if got, carry := nextChar(example.given); got != example.expected || carry != example.carry {
				tt.Errorf("Fail - expected %v/%v but got %v/%v", example.expected, example.carry, got, carry)
			}
		})
	}
}

func TestSolve(t *testing.T) {
	solve1, solve2 := solve(input)
	expected1, expected2 := "cqjxxyzz", "cqkaabcc"

	if solve1 != expected1 {
		t.Errorf("Fail - part 1. Expected %s, got %s", expected1, solve1)
	}

	if solve2 != expected2 {
		t.Errorf("Fail - part 2. Expected %s, got %s", expected2, solve2)
	}
}

func BenchmarkSolve(b *testing.B) {
	for i := 0; i < b.N; i++ {
		solve(input)
	}
}
