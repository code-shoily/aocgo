package seq

import (
	"fmt"
	"reflect"
	"testing"
)

func TestChunkByWithoutDiscard(t *testing.T) {
	examples := []struct {
		sequence []int
		chunk    int
		expect   [][]int
	}{
		{[]int{1, 2, 3}, 1, [][]int{{1}, {2}, {3}}},
		{[]int{1, 2, 3, 4}, 2, [][]int{{1, 2}, {3, 4}}},
		{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9}, 3, [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}},
	}

	for _, example := range examples {
		name := fmt.Sprintf("testing for sequence %v chunk %v", example.sequence, example.chunk)
		t.Run(name, func(tt *testing.T) {
			if got := Chunk(example.sequence, example.chunk, example.chunk, false); !reflect.DeepEqual(example.expect, got) {
				tt.Errorf("Fail - expected %v but got %v", example.expect, got)
			}
		})
	}
}

func TestChunkByWithDiscard(t *testing.T) {
	examples := []struct {
		sequence []int
		chunk    int
		expect   [][]int
	}{
		{[]int{1, 2, 3}, 1, [][]int{{1}, {2}, {3}}},
		{[]int{1, 2, 3}, 2, [][]int{{1, 2}}},
		{[]int{1, 2, 3, 4}, 3, [][]int{{1, 2, 3}}},
		{[]int{1, 2, 3, 4, 5, 6}, 3, [][]int{{1, 2, 3}, {4, 5, 6}}},
	}

	for _, example := range examples {
		name := fmt.Sprintf("testing for sequence %v chunk %v", example.sequence, example.chunk)
		t.Run(name, func(tt *testing.T) {
			if got := Chunk(example.sequence, example.chunk, example.chunk, true); !reflect.DeepEqual(example.expect, got) {
				tt.Errorf("Fail - expected %v but got %v", example.expect, got)
			}
		})
	}
}

func TestChunkByWithInterval(t *testing.T) {
	examples := []struct {
		sequence []int
		chunk    int
		interval int
		discard  bool
		expect   [][]int
	}{
		{[]int{1, 2, 3}, 1, 1, false, [][]int{{1}, {2}, {3}}},
		{[]int{1, 2, 3}, 1, 1, true, [][]int{{1}, {2}, {3}}},
		{[]int{1, 2, 3}, 2, 1, false, [][]int{{1, 2}, {2, 3}}},
		{[]int{1, 2, 3}, 2, 1, true, [][]int{{1, 2}, {2, 3}}},
		{[]int{1, 2, 3, 4}, 3, 2, false, [][]int{{1, 2, 3}, {3, 4}}},
		{[]int{1, 2, 3, 4}, 3, 2, true, [][]int{{1, 2, 3}}},
	}

	for _, example := range examples {
		name := fmt.Sprintf("testing for sequence %v chunk %v", example.sequence, example.chunk)
		t.Run(name, func(tt *testing.T) {
			if got := Chunk(example.sequence, example.chunk, example.interval, example.discard); !reflect.DeepEqual(example.expect, got) {
				tt.Errorf("Fail - expected %v but got %v", example.expect, got)
			}
		})
	}
}

func TestFrequencies(t *testing.T) {
	examples := []struct {
		given  []int
		expect map[int]int
	}{
		{[]int{3, 1, 2, -1}, map[int]int{3: 1, 1: 1, 2: 1, -1: 1}},
		{[]int{3, 3, 2, 2, 2, 1, 3}, map[int]int{3: 3, 1: 1, 2: 3}},
		{[]int{3, 3, 3}, map[int]int{3: 3}},
		{[]int{}, map[int]int{}},
	}

	for _, example := range examples {
		name := fmt.Sprintf("testing for input %v", example.given)

		t.Run(name, func(tt *testing.T) {
			if got := Frequencies(example.given); !reflect.DeepEqual(example.expect, got) {
				tt.Errorf("Fail - expected %v, got %v", example.expect, got)
			}
		})
	}
}

func TestGetMinMax(t *testing.T) {
	examples := []struct {
		given     []int
		expectMin int
		expectMax int
	}{
		{[]int{3, 1, 2, -1}, -1, 3},
		{[]int{3, 3, 2, 2, 2, 3}, 2, 3},
		{[]int{3, 3, 3}, 3, 3},
	}

	for _, example := range examples {
		name := fmt.Sprintf("testing for input %v", example.given)

		t.Run(name, func(tt *testing.T) {
			if min, max := GetMinMax(example.given); min != example.expectMin || max != example.expectMax {
				tt.Errorf("Fail - expected %v/%v, got %v/%v",
					example.expectMin,
					example.expectMax,
					min,
					max)
			}
		})
	}
}

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
			if got := Transpose(example.given); !reflect.DeepEqual(example.expect, got) {
				tt.Errorf("Fail - expected %v but got %v", example.expect, got)
			}
		})
	}
}

func TestMakeSet(t *testing.T) {
	examples := []struct {
		given  []string
		expect map[string]bool
	}{
		{[]string{}, map[string]bool{}},
		{[]string{"a"}, map[string]bool{"a": true}},
		{[]string{"a", "b"}, map[string]bool{"a": true, "b": true}},
		{[]string{"a", "a", "b"}, map[string]bool{"a": true, "b": true}},
		{[]string{"a", "a", "a"}, map[string]bool{"a": true}},
	}

	for _, example := range examples {
		name := fmt.Sprintf("testing for input %v", example.given)
		t.Run(name, func(tt *testing.T) {
			if got := MakeSet[string](example.given); !reflect.DeepEqual(example.expect, got) {
				tt.Errorf("Fail - expected %v but got %v", example.expect, got)
			}
		})
	}
}

func TestReverse(t *testing.T) {
	examples := []struct {
		seq []int
		rev []int
	}{
		{[]int{}, []int{}},
		{[]int{1}, []int{1}},
		{[]int{1, 2}, []int{2, 1}},
		{[]int{1, 2, 3}, []int{3, 2, 1}},
	}

	for _, example := range examples {
		name := fmt.Sprintf("testing for input %v", example)

		t.Run(name, func(tt *testing.T) {
			if rev := Reverse(example.seq); !reflect.DeepEqual(example.rev, rev) {
				tt.Errorf("Fail - expected %v but got %v", example.rev, rev)
			}
		})
	}
}

func TestReverseString(t *testing.T) {
	examples := []struct {
		seq string
		rev string
	}{
		{"", ""},
		{"a", "a"},
		{"aoc", "coa"},
	}

	for _, example := range examples {
		name := fmt.Sprintf("testing for input %v", example)

		t.Run(name, func(tt *testing.T) {
			if rev := ReverseString(example.seq); rev != example.rev {
				tt.Errorf("Fail - expected %v but got %v", example.rev, rev)
			}
		})
	}
}

func TestSum(t *testing.T) {
	examples := []struct {
		given  []int
		expect int
	}{
		{[]int{1}, 1},
		{[]int{1, 2}, 3},
		{[]int{1, 2, 3}, 6},
	}

	for _, example := range examples {
		name := fmt.Sprintf("testing for input %v", example.given)

		t.Run(name, func(tt *testing.T) {
			if total := Sum(example.given); total != example.expect {
				tt.Errorf("Fail - expected %v but got %v", example.expect, total)
			}
		})
	}
}

func TestProduct(t *testing.T) {
	examples := []struct {
		given  []int
		expect int
	}{
		{[]int{1}, 1},
		{[]int{1, 2}, 2},
		{[]int{1, 2, 3}, 6},
		{[]int{1, 2, 3, 4}, 24},
	}

	for _, example := range examples {
		name := fmt.Sprintf("testing for input %v", example.given)

		t.Run(name, func(tt *testing.T) {
			if total := Product(example.given); total != example.expect {
				tt.Errorf("Fail - expected %v but got %v", example.expect, total)
			}
		})
	}
}
