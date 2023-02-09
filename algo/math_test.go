package algo

import (
	"fmt"
	"reflect"
	"testing"
)

func TestFactorial(t *testing.T) {
	given := []int{1, 2, 3, 4, 12}
	expect := [][]int{
		{1},
		{1, 2},
		{1, 3},
		{1, 4, 2},
		{1, 12, 2, 6, 3, 4},
	}

	for i := 0; i < len(given); i++ {
		name := fmt.Sprintf("Testing for input %v", given[i])
		t.Run(name, func(tt *testing.T) {
			if got := Factors(given[i]); !reflect.DeepEqual(expect[i], got) {
				t.Errorf("Fail - expected %v but got %v", expect[i], got)
			}
		})
	}
}
