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
		{"1", []int{1}},
		{"1212", []int{1, 2, 1, 2}},
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
		given  string
		expect int
	}{
		{"1122", 3},
		{"1111", 4},
		{"1234", 0},
		{"91212129", 9},
	}

	for _, example := range examples {
		name := fmt.Sprintf("Testing for input %s", example.given)
		t.Run(name, func(t *testing.T) {
			given := parse(example.given)
			size := len(given)
			if got := solvePart1(given, size); got != example.expect {
				t.Errorf("Fail - expected %v, got %v", example.expect, got)
			}
		})
	}
}

func TestSolvePart2(t *testing.T) {
	examples := []struct {
		given  string
		expect int
	}{
		{"1212", 6},
		{"1221", 0},
		{"123425", 4},
		{"123123", 12},
		{"12131415", 4},
	}

	for _, example := range examples {
		name := fmt.Sprintf("Testing for input %s", example.given)
		t.Run(name, func(t *testing.T) {
			given := parse(example.given)
			size := len(given)
			if got := solvePart2(given, size); got != example.expect {
				t.Errorf("Fail - expected %v, got %v", example.expect, got)
			}
		})
	}
}

func TestSolve(t *testing.T) {
	solve1, solve2 := solve(input)
	expect1, expect2 := 1089, 1156

	if solve1 != 1089 {
		t.Errorf("Fail - part 1. Expected %d, got %d", expect1, solve1)
	}

	if solve2 != 1156 {
		t.Errorf("Fail - part 2. Expected %d, got %d", expect2, solve2)
	}
}

func BenchmarkSolvePart1(b *testing.B) {
	data := parse(input)
	for i := 0; i < b.N; i++ {
		solvePart1(data, len(data))
	}
}

func BenchmarkSolvePart2(b *testing.B) {
	data := parse(input)
	for i := 0; i < b.N; i++ {
		solvePart2(data, len(data))
	}
}
