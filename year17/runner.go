package year17

import (
	"fmt"
	"github.com/code-shoily/aocgo/year17/day01"
	"github.com/code-shoily/aocgo/year17/day02"
	"github.com/code-shoily/aocgo/year17/day05"
)

func SolveForDay(day int) {
	switch day {
	case 1:
		day01.Run()
	case 2:
		day02.Run()
	case 5:
		day05.Run()
	default:
		fmt.Printf("2017/%d has not been solved yet", day)
	}
}
