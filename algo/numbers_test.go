package algo

import (
	"fmt"
	"reflect"
	"testing"
)

func TestIntToDigits(t *testing.T) {
	examples := []struct {
		given  int
		expect []int
	}{
		{1, []int{1}},
		{12, []int{1, 2}},
		{1000, []int{1, 0, 0, 0}},
	}

	for _, example := range examples {
		if got := IntToDigits(example.given); !reflect.DeepEqual(got, example.expect) {
			name := fmt.Sprintf("testing for input %v", example.given)
			t.Run(name, func(tt *testing.T) {
				tt.Errorf("Fail - expected %v but got %v", example.expect, got)
			})
		}
	}
}
