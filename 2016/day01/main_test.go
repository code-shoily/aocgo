package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestParseInput(t *testing.T) {
	examples := []struct {
		given  string
		expect []Instruction
	}{
		{"L1", []Instruction{{"L", 1}}},
		{"L1, R2", []Instruction{{"L", 1}, {"R", 2}}},
		{"L10, R122", []Instruction{{"L", 10}, {"R", 122}}},
	}

	for _, example := range examples {
		name := fmt.Sprintf("testing for %s input", example.given)
		t.Run(name, func(t *testing.T) {
			if got := parse(example.given); !reflect.DeepEqual(got, example.expect) {
				t.Errorf("Fail - expected %v, got %v", example.expect, got)
			}
		})
	}
}

func TestSolvePart1(t *testing.T) {
	examples := []struct {
		given  []Instruction
		expect int
	}{
		{[]Instruction{{"R", 2}, {"L", 3}}, 5},
		{[]Instruction{{"R", 5}, {"L", 5}, {"R", 5}, {"R", 3}}, 12},
		{[]Instruction{{"R", 2}, {"R", 2}, {"R", 2}}, 2},
	}

	for _, example := range examples {
		name := fmt.Sprintf("testing for %v input", example.given)
		t.Run(name, func(t *testing.T) {
			if got := solvePart1(example.given); got != example.expect {
				t.Errorf("Fail - expected %v, got %v", example.expect, got)
			}
		})
	}
}

func TestSolvePart2(t *testing.T) {
	examples := []struct {
		given  []Instruction
		expect int
	}{
		{[]Instruction{{"R", 8}, {"R", 4}, {"R", 4}, {"R", 8}}, 4},
	}

	for _, example := range examples {
		name := fmt.Sprintf("testing for %v input", example.given)
		t.Run(name, func(t *testing.T) {
			if got := solvePart2(example.given); got != example.expect {
				t.Errorf("Fail - expected %v, got %v", example.expect, got)
			}
		})
	}
}

func TestSolve(t *testing.T) {
	solve1, solve2 := solve(input)

	if solve1 != 253 {
		t.Errorf("Fail - part 1 is failing, got %d, expected 253", solve1)
	}

	if solve2 != 126 {
		t.Errorf("Fail - part 2 is failing, got %d, expected 126", solve1)
	}

}

func BenchmarkSolvePart1(b *testing.B) {
	instructions := parse(input)
	for i := 0; i < b.N; i++ {
		solvePart1(instructions)
	}
}

func BenchmarkSolvePart2(b *testing.B) {
	instructions := parse(input)
	for i := 0; i < b.N; i++ {
		solvePart2(instructions)
	}
}
