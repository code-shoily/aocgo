package day02

import (
	"reflect"
	"testing"
)

func TestParseLine(t *testing.T) {
	expected := []string{"5-16", "b", "bbbbhbbbbpbxbbbcb"}
	got := parseLine("5-16 b: bbbbhbbbbpbxbbbcb")
	if !reflect.DeepEqual(expected, got) {
		t.Errorf("Fail - expected %v, got %v", expected, got)
	}
}

func TestNewPassword(t *testing.T) {
	expected := PasswordData{5, 16, "b", "bbbbhbbbbpbxbbbcb"}
	got := newPasswordData("5-16 b: bbbbhbbbbpbxbbbcb")
	if expected != got {
		t.Errorf("Fail - expected %#v, got %#v", expected, got)
	}
}

func TestSolve(t *testing.T) {
	solve1, solve2 := solve(input)
	expected1, expected2 := 607, 321

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
