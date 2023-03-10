package year17

import (
	"fmt"
	"github.com/code-shoily/aocgo/year17/day01"
	"github.com/code-shoily/aocgo/year17/day02"
	"github.com/code-shoily/aocgo/year17/day04"
	"github.com/code-shoily/aocgo/year17/day05"
	"github.com/code-shoily/aocgo/year17/day06"
	"github.com/code-shoily/aocgo/year17/day08"
	"github.com/code-shoily/aocgo/year17/day09"
	"github.com/code-shoily/aocgo/year17/day25"
)

func SolveForDay(day int) {
	switch day {
	case 1:
		day01.Run()
	case 2:
		day02.Run()
	case 4:
		day04.Run()
	case 5:
		day05.Run()
	case 6:
		day06.Run()
	case 8:
		day08.Run()
	case 9:
		day09.Run()
	case 25:
		day25.Run()
	default:
		fmt.Printf("2017/%d has not been solved yet", day)
	}
}
