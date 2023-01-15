package day05

import (
	"fmt"
	"testing"
)

func TestLineIsVertical(t *testing.T) {
	examples := []struct {
		given  Line
		expect bool
	}{
		{Line{Point{1, 50}, Point{10, 50}}, true},
		{Line{Point{0, 0}, Point{16, 0}}, true},
		{Line{Point{11, 5}, Point{11, 50}}, false},
		{Line{Point{1, 2}, Point{1, 5}}, false},
		{Line{Point{1, 1}, Point{10, 10}}, false},
		{Line{Point{50, 50}, Point{100, 100}}, false},
	}

	for _, example := range examples {
		name := fmt.Sprintf("testing for input %v", example.given)

		t.Run(name, func(tt *testing.T) {
			if got := example.given.isVertical(); example.expect != got {
				tt.Errorf("Fail - expected %v but got %v", example.expect, got)
			}
		})
	}
}

func TestLineIsHorizontal(t *testing.T) {
	examples := []struct {
		given  Line
		expect bool
	}{
		{Line{Point{1, 50}, Point{10, 50}}, false},
		{Line{Point{0, 0}, Point{16, 0}}, false},
		{Line{Point{11, 5}, Point{11, 50}}, true},
		{Line{Point{1, 2}, Point{1, 5}}, true},
		{Line{Point{1, 1}, Point{10, 10}}, false},
		{Line{Point{50, 50}, Point{100, 100}}, false},
	}

	for _, example := range examples {
		name := fmt.Sprintf("testing for input %v", example.given)

		t.Run(name, func(tt *testing.T) {
			if got := example.given.isHorizontal(); example.expect != got {
				tt.Errorf("Fail - expected %v but got %v", example.expect, got)
			}
		})
	}
}

func TestLineIsDiagonal(t *testing.T) {
	examples := []struct {
		given  Line
		expect bool
	}{
		{Line{Point{1, 50}, Point{10, 50}}, false},
		{Line{Point{0, 0}, Point{16, 0}}, false},
		{Line{Point{11, 5}, Point{11, 50}}, false},
		{Line{Point{1, 2}, Point{1, 5}}, false},
		{Line{Point{1, 1}, Point{10, 10}}, true},
		{Line{Point{50, 50}, Point{100, 100}}, true},
	}

	for _, example := range examples {
		name := fmt.Sprintf("testing for input %v", example.given)

		t.Run(name, func(tt *testing.T) {
			if got := example.given.isDiagonal(); example.expect != got {
				tt.Errorf("Fail - expected %v but got %v", example.expect, got)
			}
		})
	}
}

func TestSolve(t *testing.T) {
	solve1, solve2 := solve(input)
	expected1, expected2 := 4655, 20_500

	if solve1 != expected1 {
		t.Errorf("Fail - part 1. Expected %d, got %d", expected1, solve1)
	}

	if solve2 != expected2 {
		t.Errorf("Fail - part 2. Expected %d, got %d", expected2, solve2)
	}
}

func BenchmarkSolve(b *testing.B) {
	for i := 0; i < b.N; i++ {
		solve(input)
	}
}
