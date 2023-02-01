package year16

import (
	"fmt"
	"github.com/code-shoily/aocgo/year16/day01"
	"github.com/code-shoily/aocgo/year16/day02"
	"github.com/code-shoily/aocgo/year16/day03"
	"github.com/code-shoily/aocgo/year16/day04"
	"github.com/code-shoily/aocgo/year16/day05"
	"github.com/code-shoily/aocgo/year16/day06"
	"github.com/code-shoily/aocgo/year16/day07"
	"github.com/code-shoily/aocgo/year16/day12"
	"github.com/code-shoily/aocgo/year16/day13"
	"github.com/code-shoily/aocgo/year16/day16"
	"github.com/code-shoily/aocgo/year16/day23"
)

func SolveForDay(day int) {
	switch day {
	case 1:
		day01.Run()
	case 2:
		day02.Run()
	case 3:
		day03.Run()
	case 4:
		day04.Run()
	case 5:
		day05.Run()
	case 6:
		day06.Run()
	case 7:
		day07.Run()
	case 12:
		day12.Run()
	case 13:
		day13.Run()
	case 16:
		day16.Run()
	case 23:
		day23.Run()
	default:
		fmt.Printf("2016/%d has not been solved yet", day)
	}
}
