package year24

import (
	"fmt"

	"github.com/code-shoily/aocgo/year24/day01"
)

func SolveForDay(day int) {
	switch day {
	case 1:
		day01.Run()
	default:
		fmt.Printf("2023/%d has not been solved yet", day)
	}
}
