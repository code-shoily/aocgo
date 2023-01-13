package algo

import (
	"fmt"
	"testing"
)

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
