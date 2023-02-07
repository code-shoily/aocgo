package algo

import (
	"fmt"
	"reflect"
	"testing"
)

func TestTwoSumSorted(t *testing.T) {
	examples := []struct {
		given  []int
		target int
		a      int
		b      int
		ok     bool
	}{
		{[]int{5, 8, 10, 10}, 20, 10, 10, true},
		{[]int{1, 2, 3, 4}, 6, 2, 4, true},
		{[]int{3, 8, 5, 16}, 16, 0, 0, false},
	}

	for _, example := range examples {
		name := fmt.Sprintf("Testing for input %v", example.given)
		t.Run(name, func(t *testing.T) {
			a, b, ok := TwoSumSorted(example.given, example.target)
			if a != example.a || b != example.b || ok != example.ok {
				t.Errorf("Fail - expected %v for %d and %d but got %v for %d and %d",
					example.ok, example.a, example.b, ok, a, b)
			}
		})
	}
}

func TestThreeSumSorted(t *testing.T) {
	examples := []struct {
		given  []int
		target int
		a      int
		b      int
		c      int
		ok     bool
	}{
		{[]int{5, 5, 8, 10, 10}, 20, 5, 5, 10, true},
		{[]int{1, 2, 3, 4}, 6, 1, 2, 3, true},
		{[]int{3, 8, 5, 16}, 16, 3, 8, 5, true},
		{[]int{3, 8, 5, 16}, 17, 0, 0, 0, false},
	}

	for _, example := range examples {
		name := fmt.Sprintf("Testing for input %v", example.given)
		t.Run(name, func(t *testing.T) {
			a, b, c, ok := ThreeSumSorted(example.given, example.target)
			if a != example.a || b != example.b || c != example.c || ok != example.ok {
				t.Errorf("Fail - expected %v for %d, %d and %d but got %v for %d, %d and %d",
					example.ok, example.a, example.b, example.c, ok, a, b, c)
			}
		})
	}
}

func TestTwoSum(t *testing.T) {
	examples := []struct {
		given  []int
		target int
		a      int
		b      int
		ok     bool
	}{
		{[]int{5, 8, 10, 10}, 20, 10, 10, true},
		{[]int{1, 2, 3, 4}, 6, 2, 4, true},
		{[]int{3, 8, 5, 16}, 16, 0, 0, false},
	}

	for _, example := range examples {
		name := fmt.Sprintf("Testing for input %v", example.given)
		t.Run(name, func(t *testing.T) {
			a, b, ok := TwoSum(example.given, example.target)
			if a != example.a || b != example.b || ok != example.ok {
				t.Errorf("Fail - expected %v for %d and %d but got %v for %d and %d",
					example.ok, example.a, example.b, ok, a, b)
			}
		})
	}
}

func TestThreeSum(t *testing.T) {
	examples := []struct {
		given  []int
		target int
		a      int
		b      int
		c      int
		ok     bool
	}{
		{[]int{5, 5, 8, 10, 10}, 20, 5, 5, 10, true},
		{[]int{1, 2, 3, 4}, 6, 1, 2, 3, true},
		{[]int{3, 8, 5, 16}, 16, 3, 5, 8, true},
		{[]int{3, 8, 5, 16}, 17, 0, 0, 0, false},
	}

	for _, example := range examples {
		name := fmt.Sprintf("Testing for input %v", example.given)
		t.Run(name, func(t *testing.T) {
			a, b, c, ok := ThreeSum(example.given, example.target)
			if a != example.a || b != example.b || c != example.c || ok != example.ok {
				t.Errorf("Fail - expected %v for %d, %d and %d but got %v for %d, %d and %d",
					example.ok, example.a, example.b, example.c, ok, a, b, c)
			}
		})
	}
}

func TestSubArraySum(t *testing.T) {
	examples := []struct {
		given  []int
		target int
		a      int
		b      int
		ok     bool
	}{
		{[]int{5}, 5, 0, 0, true},
		{[]int{5, 5}, 10, 0, 1, true},
		{[]int{5, -5}, 0, 0, 1, true},
		{[]int{-3, 8, 5, 13}, 10, 0, 2, true},
		{[]int{5, 8, 10, 10, 2, 5, 2, 4, 5}, 13, 0, 1, true},
		{[]int{5, 8, 10, 10, 2, 5, 2, 4, 5}, 22, 2, 4, true},
		{[]int{5, 8, 10, 10, 2, 5, 2, 4, 5}, 50, 0, 0, false},
		{[]int{3, 18, 5, 13}, 16, 0, 0, false},
		{[]int{5, 2, 5}, 10, 0, 0, false},
	}

	for _, example := range examples {
		name := fmt.Sprintf("Testing for input %v", example.given)
		t.Run(name, func(t *testing.T) {
			a, b, ok := SubArraySum(example.given, example.target)
			if a != example.a || b != example.b || ok != example.ok {
				t.Errorf("Fail - expected %v for %d and %d but got %v for %d and %d",
					example.ok, example.a, example.b, ok, a, b)
			}
		})
	}
}

func TestSubsetSum(t *testing.T) {
	examples := []struct {
		given  []int
		target int
		expect [][]int
	}{
		{[]int{1}, 1, [][]int{{1}}},
		{[]int{1, 2}, 10, [][]int(nil)},
		{[]int{1, 2, 3}, 3, [][]int{{1, 2}, {3}}},
		{[]int{20, 15, 10, 5, 5}, 25, [][]int{{20, 5}, {20, 5}, {15, 10}, {15, 5, 5}}},
	}

	for _, example := range examples {
		name := fmt.Sprintf("Testing for input %v", example.given)
		t.Run(name, func(t *testing.T) {

			if got := SubsetSum(example.given, example.target); !reflect.DeepEqual(example.expect, got) {
				t.Errorf("Fail - expected %v but got %v", example.expect, got)
			}
		})
	}
}
