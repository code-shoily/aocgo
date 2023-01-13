package year15

import (
	"fmt"
	"github.com/code-shoily/aocgo/year15/day01"
	"github.com/code-shoily/aocgo/year15/day02"
	"github.com/code-shoily/aocgo/year15/day03"
	"github.com/code-shoily/aocgo/year15/day05"
	"github.com/code-shoily/aocgo/year15/day06"
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
	case 6:
		day06.Run()
	default:
		fmt.Printf("2015/%d has not been solved yet", day)
	}
}
