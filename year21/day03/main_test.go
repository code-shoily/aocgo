package day03

import (
	"reflect"
	"testing"
)

func TestParse(t *testing.T) {
	given := "1101\n1001\n1111\n1000"
	expect := [][]int{
		{1, 1, 0, 1},
		{1, 0, 0, 1},
		{1, 1, 1, 1},
		{1, 0, 0, 0},
	}
	if got := parse(given); !reflect.DeepEqual(expect, got) {
		t.Errorf("Fail - expected %v but got %v", expect, got)
	}
}
func TestSolve(t *testing.T) {
	solve1, solve2 := solve(input)
	expected1, expected2 := 1_540_244, 4_203_981

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
