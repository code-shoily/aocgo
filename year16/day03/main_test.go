package day03

import (
	"fmt"
	"reflect"
	"testing"
)

func TestTranspose(t *testing.T) {
	examples := []struct {
		given  [][]int
		expect [][]int
	}{
		{[][]int{{1, 2}}, [][]int{{1}, {2}}},
		{[][]int{{1, 2, 3}, {4, 5, 6}}, [][]int{{1, 4}, {2, 5}, {3, 6}}},
	}

	for _, example := range examples {
		name := fmt.Sprintf("testing for input %v", example.given)
		t.Run(name, func(tt *testing.T) {
			if got := transpose(example.given); !reflect.DeepEqual(example.expect, got) {
				tt.Errorf("Fail - expected %v but got %v", example.expect, got)
			}
		})
	}
}

func TestChunk(t *testing.T) {
	examples := []struct {
		sequence []int
		chunk    int
		expect   [][]int
	}{
		{[]int{1, 2, 3, 4}, 2, [][]int{{1, 2}, {3, 4}}},
		{[]int{1, 2, 3, 4, 5, 6}, 3, [][]int{{1, 2, 3}, {4, 5, 6}}},
	}

	for _, example := range examples {
		name := fmt.Sprintf("testing for sequence %v chunk %v", example.sequence, example.chunk)
		t.Run(name, func(tt *testing.T) {
			if got := chunkBy(example.sequence, example.chunk); !reflect.DeepEqual(example.expect, got) {
				tt.Errorf("Fail - expected %v but got %v", example.expect, got)
			}
		})
	}
}

func TestSolve(t *testing.T) {
	solve1, solve2 := solve(input)
	expected1, expected2 := 993, 1849

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
