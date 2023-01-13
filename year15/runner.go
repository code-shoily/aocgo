package year15

import (
	"fmt"
	"github.com/code-shoily/aocgo/year15/day01"
	"github.com/code-shoily/aocgo/year15/day02"
)

func SolveForDay(day int) {
	switch day {
	case 1:
		day01.Run()
	case 2:
		day02.Run()
	default:
		fmt.Printf("2015/%d has not been solved yet", day)
	}
}
