package seq

import (
	"fmt"
	"reflect"
	"testing"
)

func TestPermutations(t *testing.T) {
	var emptyPermutation [][]int

	examples := []struct {
		given  []int
		expect [][]int
	}{
		{
			[]int{},
			emptyPermutation,
		},
		{
			[]int{1},
			[][]int{{1}},
		},
		{
			[]int{1, 2},
			[][]int{{1, 2}, {2, 1}},
		},
		{
			[]int{1, 2, 3},
			[][]int{{1, 2, 3}, {2, 1, 3}, {3, 2, 1}, {2, 3, 1}, {3, 1, 2}, {1, 3, 2}},
		},
	}

	for _, example := range examples {
		name := fmt.Sprintf("testing for input %v", example.given)
		t.Run(name, func(tt *testing.T) {
			if got := Permutations(example.given); !reflect.DeepEqual(example.expect, got) {
				tt.Errorf("Fail - expected %v got %v", example.expect, got)
			}
		})
	}
}

func TestStringPermutations(t *testing.T) {
	examples := []struct {
		given  string
		expect []string
	}{
		{"a", []string{"a"}},
		{"ab", []string{"ab", "ba"}},
		{"abc", []string{"abc", "bac", "cba", "bca", "cab", "acb"}},
	}

	for _, example := range examples {
		name := fmt.Sprintf("testing for input %v", example.given)
		t.Run(name, func(tt *testing.T) {
			if got := StringPermutations(example.given); !reflect.DeepEqual(got, example.expect) {
				tt.Errorf("Fail - expected %v but got %v", example.expect, got)
			}
		})
	}
}
