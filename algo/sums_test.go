package algo

import (
	"fmt"
	"testing"
)

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
		{[]int{3, 8, 5, 16}, 16, -1, -1, false},
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
		{[]int{3, 8, 5, 16}, 16, 3, 8, 5, true},
		{[]int{3, 8, 5, 16}, 17, -1, -1, -1, false},
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
