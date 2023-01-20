package day07

import (
	"fmt"
	"reflect"
	"testing"
)

func TestHasABBA(t *testing.T) {
	examples := []struct {
		given  string
		expect bool
	}{
		{"aaaa", false},
		{"terabbamnopqrst", true},
		{"abyamnopqrst", false},
		{"ioxxojasdfghzxcvbn", true},
	}

	for _, example := range examples {
		name := fmt.Sprintf("testing for input %v", example.given)
		t.Run(name, func(tt *testing.T) {
			if got := hasABBA(example.given); got != example.expect {
				tt.Errorf("Fail - expected %v but got %v", example.expect, got)
			}
		})
	}
}

func TestGetABAtoBAB(t *testing.T) {
	examples := []struct {
		given  string
		expect map[string]bool
	}{
		{"abaxyz", map[string]bool{"bab": true}},
		{"aaakekeke", map[string]bool{"kek": true, "eke": true}},
		{"zazbzbzbcdb", map[string]bool{"aza": true, "zbz": true, "bzb": true}},
	}

	for _, example := range examples {
		name := fmt.Sprintf("testing for input %v", example.given)
		t.Run(name, func(tt *testing.T) {
			if got := getABAtoBAB(example.given); !reflect.DeepEqual(example.expect, got) {
				tt.Errorf("Fail - expected %v but got %v", example.expect, got)
			}
		})
	}
}

func TestIPv7_SupportsTLS(t *testing.T) {
	examples := []struct {
		packets   []string
		hypernets []string
		expect    bool
	}{
		{[]string{"abba", "qrst"}, []string{"mnop"}, true},
		{[]string{"aaaa", "tyui"}, []string{"qwer"}, false},
		{[]string{"abcd", "xyyx"}, []string{"bddb"}, false},
		{[]string{"ioxxoj", "zxcvbn"}, []string{"asdfgh"}, true},
	}

	for _, example := range examples {
		given := IPv7{
			example.packets,
			example.hypernets,
		}
		name := fmt.Sprintf("testing for input %v", given)
		t.Run(name, func(tt *testing.T) {
			if got := given.supportsTLS(); got != example.expect {
				tt.Errorf("Fail - expected %v but got %v", example.expect, got)
			}
		})
	}
}

func TestIPv7_SupportsSSL(t *testing.T) {
	examples := []struct {
		packets   []string
		hypernets []string
		expect    bool
	}{
		{[]string{"aba", "xyz"}, []string{"bab"}, true},
		{[]string{"xyx", "xyx"}, []string{"xyx"}, false},
		{[]string{"aaa", "eke"}, []string{"kek"}, true},
		{[]string{"zazbz", "cdb"}, []string{"bzb"}, true},
	}

	for _, example := range examples {
		given := IPv7{
			example.packets,
			example.hypernets,
		}
		name := fmt.Sprintf("testing for input %v", given)
		t.Run(name, func(tt *testing.T) {
			if got := given.supportsSSL(); got != example.expect {
				tt.Errorf("Fail - expected %v but got %v", example.expect, got)
			}
		})
	}
}

func TestSolve(t *testing.T) {
	solve1, solve2 := solve(input)
	expected1, expected2 := 105, 258

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
