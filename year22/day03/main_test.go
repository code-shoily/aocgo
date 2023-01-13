package day03

import (
	"fmt"
	"reflect"
	"testing"
)

func TestCompartments(t *testing.T) {
	examples := []struct {
		given string
		a     string
		b     string
	}{
		{"vJrwpWtwJgWrhcsFMMfFFhFp", "vJrwpWtwJgWr", "hcsFMMfFFhFp"},
		{"jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL", "jqHRNqRjqzjGDLGL", "rsFMfFZSrLrFZsSL"},
		{"PmmdzqPrVvPwwTWBwg", "PmmdzqPrV", "vPwwTWBwg"},
	}

	for _, example := range examples {
		name := fmt.Sprintf("testing for input %v", example.given)
		t.Run(name, func(tt *testing.T) {
			if a, b := compartments(example.given); a != example.a || b != example.b {
				tt.Errorf("Fail - expected %s and %s but got %s and %s",
					example.a, example.b, a, b)
			}
		})
	}
}

func TestInBothCompartments(t *testing.T) {
	examples := []struct {
		given  string
		expect rune
	}{
		{"vJrwpWtwJgWrhcsFMMfFFhFp", 'p'},
		{"jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL", 'L'},
		{"PmmdzqPrVvPwwTWBwg", 'P'},
		{"wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn", 'v'},
		{"ttgJtRGJQctTZtZT", 't'},
		{"CrZsJsPPZsGzwwsLwLmpwMDw", 's'},
	}

	for _, example := range examples {
		name := fmt.Sprintf("testing for input %v", example.given)
		t.Run(name, func(tt *testing.T) {
			if got := inBothCompartments(example.given); got != example.expect {
				tt.Errorf("Fail - expected %v but got %v", example.expect, got)
			}
		})
	}
}

func TestGetPriority(t *testing.T) {
	examples := []struct {
		given  rune
		expect int
	}{
		{'p', 16},
		{'L', 38},
		{'P', 42},
		{'v', 22},
		{'t', 20},
		{'s', 19},
	}

	for _, example := range examples {
		name := fmt.Sprintf("testing for input %v", example.given)
		if got := getPriority(example.given); got != example.expect {
			t.Run(name, func(tt *testing.T) {
				tt.Errorf("Fail - expected %v, got %v", example.expect, got)
			})
		}
	}
}

func TestGroupRucksacks(t *testing.T) {
	examples := []struct {
		given  []string
		expect [][]string
	}{
		{
			[]string{"a", "b", "c"},
			[][]string{{"a", "b", "c"}},
		},
		{
			[]string{"a", "b", "c", "d", "e", "f"},
			[][]string{{"a", "b", "c"}, {"d", "e", "f"}},
		},
	}

	for _, example := range examples {
		name := fmt.Sprintf("testing for input %v", example.given)
		t.Run(name, func(tt *testing.T) {
			if got := groupRucksacks(example.given); !reflect.DeepEqual(example.expect, got) {
				tt.Errorf("Fail - expected %v but got %v", example.expect, got)
			}
		})
	}
}

func TestInAllGroups(t *testing.T) {
	a, b, c, expect := "aab", "bba", "cbx", 'b'

	if got := inAllGroups(a, b, c); got != expect {
		t.Errorf("Fail - expected %v, got %v", expect, got)
	}
}

func TestSolve(t *testing.T) {
	solve1, solve2 := solve(input)
	expected1, expected2 := 8233, 2821

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
