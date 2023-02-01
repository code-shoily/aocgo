package day25

import (
	"fmt"
	"testing"
)

var conversionTable = []struct {
	snafu   string
	decimal int
}{
	{"1", 1},
	{"2", 2},
	{"1=", 3},
	{"1-", 4},
	{"10", 5},
	{"11", 6},
	{"12", 7},
	{"2=", 8},
	{"2-", 9},
	{"20", 10},
	{"1=0", 15},
	{"1-0", 20},
	{"1=11-2", 2022},
	{"1-0---0", 12345},
	{"1121-1110-1=0", 314159265},
}

func TestSnafuToDecimal(t *testing.T) {
	for _, example := range conversionTable {
		name := fmt.Sprintf("testing for input %v", example.snafu)
		t.Run(name, func(tt *testing.T) {
			if decimal := snafuToDecimal(example.snafu); decimal != example.decimal {
				t.Errorf("Fail - expected %v but got %v", example.decimal, decimal)
			}
		})
	}
}

func TestDecimalToSnafu(t *testing.T) {
	for _, example := range conversionTable {
		name := fmt.Sprintf("testing for input %v", example.snafu)
		t.Run(name, func(tt *testing.T) {
			if snafu := decimalToSnafu(example.decimal); snafu != example.snafu {
				t.Errorf("Fail - expected %v but got %v", example.snafu, snafu)
			}
		})
	}
}

func TestSolve(t *testing.T) {
	solution := solve(input)
	expected := "2-==10===-12=2-1=-=0"

	if solution != expected {
		t.Errorf("Fail - part 1. Expected %v, got %v", expected, solution)
	}
}

func BenchmarkSolve(b *testing.B) {
	for i := 0; i < b.N; i++ {
		solve(input)
	}
}
