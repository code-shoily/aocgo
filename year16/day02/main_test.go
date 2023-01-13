package day02

import "testing"

func TestSolve(t *testing.T) {
	solve1, solve2 := solve(input)
	expect1, expect2 := "76792", "A7AC3"

	if solve1 != expect1 {
		t.Errorf("Fail - part 1 is failing, got %s, expected %s", solve1, expect1)
	}

	if solve2 != expect2 {
		t.Errorf("Fail - part 2 is failing, got %s, expected %s", solve2, expect2)
	}

}

func BenchmarkSolve(b *testing.B) {
	for i := 0; i < b.N; i++ {
		solve(input)
	}
}
