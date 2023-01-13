package day02

import (
	"fmt"
	"reflect"
	"testing"
)

func TestParse(t *testing.T) {
	examples := []struct {
		given  string
		expect []Dimension
	}{
		{"1x2x3", []Dimension{
			{1, 2, 3},
		}},
		{"1x2x3\n4x5x6", []Dimension{
			{1, 2, 3},
			{4, 5, 6},
		}},
	}

	for _, example := range examples {
		name := fmt.Sprintf("testing for %v", example.given)
		t.Run(name, func(tt *testing.T) {
			if got := parse(example.given); !reflect.DeepEqual(got, example.expect) {
				tt.Errorf("Fail - expected %v, got %v", example.expect, got)
			}
		})
	}
}

func TestCalculateWrappingPaper(t *testing.T) {
	examples := []struct {
		given  Dimension
		expect int
	}{
		{Dimension{2, 3, 4}, 58},
		{Dimension{1, 1, 10}, 43},
	}

	for _, example := range examples {
		name := fmt.Sprintf("testing for input %v", example.given)
		t.Run(name, func(tt *testing.T) {
			if got := calculateWrappingPaper(example.given); got != example.expect {
				tt.Errorf("Fail - expected %v got %v", example.expect, got)
			}
		})
	}
}

func TestCalculateRibbon(t *testing.T) {
	examples := []struct {
		given  Dimension
		expect int
	}{
		{Dimension{2, 3, 4}, 34},
		{Dimension{1, 1, 10}, 14},
	}

	for _, example := range examples {
		name := fmt.Sprintf("testing for input %v", example.given)
		t.Run(name, func(tt *testing.T) {
			if got := calculateRibbon(example.given); got != example.expect {
				tt.Errorf("Fail - expected %v got %v", example.expect, got)
			}
		})
	}
}

func TestSolve(t *testing.T) {
	solve1, solve2 := solve(input)
	expect1, expect2 := 1_606_483, 3_842_356

	if solve1 != expect1 {
		t.Errorf("Part 1 failed, expected %d, expect %d", expect1, solve1)
	}

	if solve2 != expect2 {
		t.Errorf("Part 1 failed, expected %d, expect %d", expect2, solve2)
	}
}

func BenchmarkSolvePart1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		solve(input)
	}
}
