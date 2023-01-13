package day05

import (
	"fmt"
	"reflect"
	"strings"
	"testing"
)

func TestParse(t *testing.T) {
	given := "abcd\nefgh\nijkl"
	expect := [][]string{
		{"a", "b", "c", "d"},
		{"e", "f", "g", "h"},
		{"i", "j", "k", "l"},
	}
	if got := parse(given); !reflect.DeepEqual(got, expect) {
		t.Errorf("Fail - expected %v but got %v", expect, got)
	}
}

func TestInPairs(t *testing.T) {
	examples := []struct {
		given  []string
		expect [][2]string
	}{
		{
			[]string{"a"},
			[][2]string{},
		},
		{
			[]string{"a", "b"},
			[][2]string{{"a", "b"}},
		},
		{
			[]string{"a", "b", "c"},
			[][2]string{{"a", "b"}, {"b", "c"}},
		},
		{
			[]string{"a", "b", "c", "d"},
			[][2]string{{"a", "b"}, {"b", "c"}, {"c", "d"}},
		},
	}

	for _, example := range examples {
		name := fmt.Sprintf("testing for input %v", example.given)
		t.Run(name, func(tt *testing.T) {
			if got := inPairs(example.given); !reflect.DeepEqual(got, example.expect) {
				tt.Errorf("Fail - expected %v but got %v", example.given, got)
			}
		})
	}
}

func TestInTriples(t *testing.T) {
	examples := []struct {
		given  []string
		expect [][3]string
	}{
		{
			[]string{"a"},
			[][3]string{},
		},
		{
			[]string{"a", "b"},
			[][3]string{},
		},
		{
			[]string{"a", "b", "c"},
			[][3]string{{"a", "b", "c"}},
		},
		{
			[]string{"a", "b", "c", "d"},
			[][3]string{{"a", "b", "c"}, {"b", "c", "d"}},
		},
	}

	for _, example := range examples {
		name := fmt.Sprintf("testing for input %v", example.given)
		t.Run(name, func(tt *testing.T) {
			if got := inTriples(example.given); !reflect.DeepEqual(got, example.expect) {
				tt.Errorf("Fail - expected %v but got %v", example.given, got)
			}
		})
	}
}

func TestHasSandwichedLetters(t *testing.T) {
	examples := []struct {
		given  []string
		expect bool
	}{
		{
			strings.Split("xxyxx", ""),
			true,
		},
		{
			strings.Split("xyx", ""),
			true,
		},
		{
			strings.Split("abcdefeghi", ""),
			true,
		},
		{
			strings.Split("aaa", ""),
			true,
		},
	}

	for _, example := range examples {
		name := fmt.Sprintf("testing for input %v", example.given)
		t.Run(name, func(tt *testing.T) {
			if got := hasSandwichedLetters(example.given); !reflect.DeepEqual(got, example.expect) {
				tt.Errorf("Fail - expected %v but got %v", example.given, got)
			}
		})
	}
}

func TestNonOverlappingPair(t *testing.T) {
	examples := []struct {
		given  []string
		expect bool
	}{
		{
			strings.Split("xxyxx", ""),
			true,
		},
		{
			strings.Split("ieodomkazucvgmuy", ""),
			false,
		},
		{
			strings.Split("uurcxstgmygtbstg", ""),
			true,
		},
		{
			strings.Split("xyxy", ""),
			true,
		},
		{
			strings.Split("aabcdefgaa", ""),
			true,
		},
		{
			strings.Split("aaa", ""),
			false,
		},
		{
			strings.Split("qjhvhtzxzqqjkmpb", ""),
			true,
		},
	}

	for _, example := range examples {
		name := fmt.Sprintf("testing for input %v", example.given)
		t.Run(name, func(tt *testing.T) {
			if got := hasNonOverlappingPair(example.given); !reflect.DeepEqual(got, example.expect) {
				tt.Errorf("Fail - expected %v but got %v", example.given, got)
			}
		})
	}
}

func TestVowelCount(t *testing.T) {
	examples := []struct {
		given  []string
		expect int
	}{
		{
			[]string{"a"},
			1,
		},
		{
			[]string{"x", "y"},
			0,
		},
		{
			[]string{"a", "w", "o", "l"},
			2,
		},
		{
			[]string{"a", "e", "i", "o", "u"},
			5,
		},
	}

	for _, example := range examples {
		name := fmt.Sprintf("testing for input %v", example.given)
		t.Run(name, func(tt *testing.T) {
			if got := vowelCount(example.given); got != example.expect {
				tt.Errorf("Fail - expected %v but got %v", example.given, got)
			}
		})
	}
}

func TestIsNiceV1(t *testing.T) {
	examples := []struct {
		given  []string
		expect bool
	}{
		{
			strings.Split("ugknbfddgicrmopn", ""),
			true,
		},
		{
			strings.Split("aaa", ""),
			true,
		},
		{
			strings.Split("jchzalrnumimnmhp", ""),
			false,
		},
		{
			strings.Split("haegwjzuvuyypxyu", ""),
			false,
		},
		{
			strings.Split("dvszwmarrgswjxmb", ""),
			false,
		},
	}

	for _, example := range examples {
		name := fmt.Sprintf("testing for input %v", example.given)
		t.Run(name, func(tt *testing.T) {
			if got := isNiceV1(example.given); got != example.expect {
				tt.Errorf("Fail - expected %v but got %v", example.given, got)
			}
		})
	}
}

func TestIsNiceV2(t *testing.T) {
	examples := []struct {
		given  []string
		expect bool
	}{
		{
			strings.Split("qjhvhtzxzqqjkmpb", ""),
			true,
		},
		{
			strings.Split("xxyxx", ""),
			true,
		},
		{
			strings.Split("uurcxstgmygtbstg", ""),
			false,
		},
		{
			strings.Split("ieodomkazucvgmuy", ""),
			false,
		},
	}

	for _, example := range examples {
		name := fmt.Sprintf("testing for input %v", example.given)
		t.Run(name, func(tt *testing.T) {
			if got := isNiceV2(example.given); got != example.expect {
				tt.Errorf("Fail - expected %v but got %v", example.given, got)
			}
		})
	}
}

func TestSolve(t *testing.T) {
	solve1, solve2 := solve(input)
	expected1, expected2 := 255, 55

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

func BenchmarkSolve(b *testing.B) {
	for i := 0; i < b.N; i++ {
		solve(input)
	}
}
