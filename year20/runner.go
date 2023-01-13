package year20

import (
	"fmt"
	"github.com/code-shoily/aocgo/year20/day01"
)

func SolveForDay(day int) {
	switch day {
	case 1:
		day01.Run()
	default:
		fmt.Printf("2020/%d has not been solved yet", day)
	}
}
