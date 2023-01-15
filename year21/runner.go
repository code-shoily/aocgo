package year21

import (
	"fmt"
	"github.com/code-shoily/aocgo/year21/day01"
	"github.com/code-shoily/aocgo/year21/day02"
	"github.com/code-shoily/aocgo/year21/day03"
	"github.com/code-shoily/aocgo/year21/day05"
)

func SolveForDay(day int) {
	switch day {
	case 1:
		day01.Run()
	case 2:
		day02.Run()
	case 3:
		day03.Run()
	case 5:
		day05.Run()
	default:
		fmt.Printf("2021/%d has not been solved yet", day)
	}
}
