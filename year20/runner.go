package year20

import (
	"fmt"
	"github.com/code-shoily/aocgo/year20/day01"
	"github.com/code-shoily/aocgo/year20/day02"
	"github.com/code-shoily/aocgo/year20/day03"
	"github.com/code-shoily/aocgo/year20/day04"
	"github.com/code-shoily/aocgo/year20/day05"
	"github.com/code-shoily/aocgo/year20/day06"
	"github.com/code-shoily/aocgo/year20/day07"
	"github.com/code-shoily/aocgo/year20/day08"
	"github.com/code-shoily/aocgo/year20/day09"
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
	case 8:
		day08.Run()
	case 9:
		day09.Run()
	default:
		fmt.Printf("2020/%d has not been solved yet", day)
	}
}
