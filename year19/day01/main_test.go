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
		{"12", []int{12}},
		{"12\n14", []int{12, 14}},
		{"132\n22123\n4003", []int{132, 22123, 4003}},
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

func TestGetFuel1(t *testing.T) {
	examples := []struct {
		given  int
		expect int
	}{
		{12, 2},
		{14, 2},
		{1969, 654},
		{100756, 33583},
	}

	for _, example := range examples {
		name := fmt.Sprintf("Testing for input %d", example.given)

		t.Run(name, func(t *testing.T) {
			if got := getFuel1(example.given); got != example.expect {
				t.Errorf("Fail - expected %d, got %d", example.expect, got)
			}
		})
	}
}

func TestGetFuel2(t *testing.T) {
	examples := []struct {
		given  int
		expect int
	}{
		{14, 2},
		{1969, 966},
		{100756, 50346},
	}

	for _, example := range examples {
		name := fmt.Sprintf("Testing for input %d", example.given)

		t.Run(name, func(t *testing.T) {
			if got := getFuel2(example.given); got != example.expect {
				t.Errorf("Fail - expected %d, got %d", example.expect, got)
			}
		})
	}
}

func TestSolve(t *testing.T) {
	solve1, solve2 := solve(input)
	expected1, expected2 := 3_421_505, 5_129_386

	if solve1 != expected1 {
		t.Errorf("Fail - part 1. Expected %d, got %d", expected1, solve1)
	}

	if solve2 != expected2 {
		t.Errorf("Fail - part 2. Expected %d, got %d", expected2, solve2)
	}
}

func BenchmarkSolvePart1(b *testing.B) {
	masses := parse(input)
	for i := 0; i < b.N; i++ {
		solvePart1(masses)
	}
}

func BenchmarkSolvePart2(b *testing.B) {
	masses := parse(input)
	for i := 0; i < b.N; i++ {
		solvePart2(masses)
	}
}
