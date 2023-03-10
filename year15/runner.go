package year15

import (
	"fmt"
	"github.com/code-shoily/aocgo/year15/day01"
	"github.com/code-shoily/aocgo/year15/day02"
	"github.com/code-shoily/aocgo/year15/day03"
	"github.com/code-shoily/aocgo/year15/day04"
	"github.com/code-shoily/aocgo/year15/day05"
	"github.com/code-shoily/aocgo/year15/day06"
	"github.com/code-shoily/aocgo/year15/day07"
	"github.com/code-shoily/aocgo/year15/day08"
	"github.com/code-shoily/aocgo/year15/day09"
	"github.com/code-shoily/aocgo/year15/day10"
	"github.com/code-shoily/aocgo/year15/day11"
	"github.com/code-shoily/aocgo/year15/day12"
	"github.com/code-shoily/aocgo/year15/day13"
	"github.com/code-shoily/aocgo/year15/day14"
	"github.com/code-shoily/aocgo/year15/day15"
	"github.com/code-shoily/aocgo/year15/day17"
	"github.com/code-shoily/aocgo/year15/day20"
	"github.com/code-shoily/aocgo/year15/day24"
	"github.com/code-shoily/aocgo/year15/day25"
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
	case 8:
		day08.Run()
	case 7:
		day07.Run()
	case 9:
		day09.Run()
	case 10:
		day10.Run()
	case 11:
		day11.Run()
	case 12:
		day12.Run()
	case 13:
		day13.Run()
	case 14:
		day14.Run()
	case 15:
		day15.Run()
	case 17:
		day17.Run()
	case 20:
		day20.Run()
	case 24:
		day24.Run()
	case 25:
		day25.Run()
	default:
		fmt.Printf("2015/%d has not been solved yet", day)
	}
}
