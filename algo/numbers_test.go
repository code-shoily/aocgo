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
		name := fmt.Sprintf("testing for input %v", example.given)
		t.Run(name, func(tt *testing.T) {
			if got := IntToDigits(example.given); !reflect.DeepEqual(got, example.expect) {
				tt.Errorf("Fail - expected %v but got %v", example.expect, got)
			}
		})
	}
}

func TestIsEven(t *testing.T) {
	examples := []struct {
		given  int
		expect bool
	}{
		{0, true},
		{1, false},
		{12, true},
		{101, false},
	}

	for _, example := range examples {

		name := fmt.Sprintf("testing for input %v", example.given)
		t.Run(name, func(tt *testing.T) {
			if got := IsEven(example.given); got != example.expect {
				tt.Errorf("Fail - expected %v but got %v", example.expect, got)
			}
		})
	}
}

func TestIsOdd(t *testing.T) {
	examples := []struct {
		given  int
		expect bool
	}{
		{0, false},
		{1, true},
		{12, false},
		{101, true},
	}

	for _, example := range examples {

		name := fmt.Sprintf("testing for input %v", example.given)
		t.Run(name, func(tt *testing.T) {
			if got := IsOdd(example.given); got != example.expect {
				tt.Errorf("Fail - expected %v but got %v", example.expect, got)
			}
		})
	}
}

func TestCountSetBits(t *testing.T) {
	examples := []struct {
		given  int
		expect int
	}{
		{0, 0},
		{1, 1},
		{2, 1},
		{3, 2},
		{4, 1},
		{5, 2},
		{6, 2},
	}

	for _, example := range examples {

		name := fmt.Sprintf("testing for input %v", example.given)
		t.Run(name, func(tt *testing.T) {
			if got := CountSetBits(example.given); got != example.expect {
				tt.Errorf("Fail - expected %v but got %v", example.expect, got)
			}
		})
	}
}

func TestSum(t *testing.T) {
	examples := []struct {
		given  []int
		expect int
	}{
		{[]int{}, 0},
		{[]int{1}, 1},
		{[]int{1, 2, -3}, 0},
		{[]int{1, 2, 3}, 6},
	}

	for _, example := range examples {
		name := fmt.Sprintf("testing for input %v", example.given)
		t.Run(name, func(tt *testing.T) {
			if got := Sum(example.given); got != example.expect {
				tt.Errorf("Fail - expected %v but got %v", example.expect, got)
			}
		})
	}
}

func TestMean(t *testing.T) {
	examples := []struct {
		given  []int
		expect float64
	}{
		{[]int{1}, 1},
		{[]int{-1}, -1},
		{[]int{1, 2, -3}, 0},
		{[]int{1, 2, 3}, 2},
	}

	for _, example := range examples {
		name := fmt.Sprintf("testing for input %v", example.given)
		t.Run(name, func(tt *testing.T) {
			if got := Mean(example.given); got != example.expect {
				tt.Errorf("Fail - expected %v but got %v", example.expect, got)
			}
		})
	}
}

func TestMedian(t *testing.T) {
	examples := []struct {
		given  []int
		expect float64
	}{
		{[]int{1}, 1},
		{[]int{-1}, -1},
		{[]int{1, 2, 3}, 2},
		{[]int{1, 22, -3, 4, 6}, 4},
		{[]int{16, 1, 2, 0, 4, 2, 7, 1, 2, 14}, 2},
	}

	for _, example := range examples {
		name := fmt.Sprintf("testing for input %v", example.given)
		t.Run(name, func(tt *testing.T) {
			if got := Median(example.given); got != example.expect {
				tt.Errorf("Fail - expected %v but got %v", example.expect, got)
			}
		})
	}
}
