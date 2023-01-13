package utils

import (
	"fmt"
	"reflect"
	"testing"
)

func TestSplitEmptyLines(t *testing.T) {
	if got := SplitLines(""); len(got) != 0 {
		t.Errorf("Fail - expected empty slice but got %v", len(got))
	}
}

func TestSplitLines(t *testing.T) {
	examples := []struct {
		given  string
		expect []string
	}{
		{"a", []string{"a"}},
		{"a\nb", []string{"a", "b"}},
	}

	for _, example := range examples {
		name := fmt.Sprintf("testing for input %v", example.given)
		t.Run(name, func(tt *testing.T) {
			if got := SplitLines(example.given); !reflect.DeepEqual(got, example.expect) {
				tt.Errorf("Fail - expected %v, got %v", got, example.expect)
			}
		})
	}
}

func TestSplitEmptyIntLines(t *testing.T) {
	if got := SplitIntLines(""); len(got) != 0 {
		t.Errorf("Fail - expected empty slice but got %v", len(got))
	}
}

func TestSplitIntLines(t *testing.T) {
	examples := []struct {
		given  string
		expect []int
	}{
		{"1", []int{1}},
		{"1\n2", []int{1, 2}},
	}

	for _, example := range examples {
		name := fmt.Sprintf("testing for input %v", example.given)
		t.Run(name, func(tt *testing.T) {
			if got := SplitIntLines(example.given); !reflect.DeepEqual(got, example.expect) {
				t.Errorf("Fail - expected %v, got %v", got, example.expect)
			}
		})
	}
}

func TestSplitEmptyBy(t *testing.T) {
	if got := SplitBy("", ", "); len(got) != 0 {
		t.Errorf("Fail - expected empty slice but got %v", len(got))
	}
}

func TestSplitBy(t *testing.T) {
	examples := []struct {
		given  string
		by     string
		expect []string
	}{
		{"a, b, c", ", ", []string{"a", "b", "c"}},
		{"a|b", "|", []string{"a", "b"}},
	}

	for _, example := range examples {
		name := fmt.Sprintf("testing for input %v", example.given)
		t.Run(name, func(tt *testing.T) {
			if got := SplitBy(example.given, example.by); !reflect.DeepEqual(got, example.expect) {
				tt.Errorf("Fail - expected %v, got %v", got, example.expect)
			}
		})
	}
}

func TestSplitEmptyByInts(t *testing.T) {
	if got := SplitByInts("", ", "); len(got) != 0 {
		t.Errorf("Fail - expected empty slice but got %v", len(got))
	}
}

func TestSplitByInts(t *testing.T) {
	examples := []struct {
		given  string
		by     string
		expect []int
	}{
		{"11, 12, 14", ", ", []int{11, 12, 14}},
		{"0|1", "|", []int{0, 1}},
	}

	for _, example := range examples {
		name := fmt.Sprintf("testing for input %v", example.given)
		t.Run(name, func(tt *testing.T) {
			if got := SplitByInts(example.given, example.by); !reflect.DeepEqual(got, example.expect) {
				tt.Errorf("Fail - expected %v, got %v", got, example.expect)
			}
		})
	}
}
