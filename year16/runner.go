package year16

import (
	"fmt"
	"github.com/code-shoily/aocgo/year16/day01"
	"github.com/code-shoily/aocgo/year16/day02"
	"github.com/code-shoily/aocgo/year16/day03"
	"github.com/code-shoily/aocgo/year16/day05"
	"github.com/code-shoily/aocgo/year16/day06"
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
		fmt.Printf("2016/%d has not been solved yet", day)
	}
}
