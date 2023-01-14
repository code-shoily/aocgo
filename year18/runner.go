package year18

import (
	"fmt"
	"github.com/code-shoily/aocgo/year18/day01"
	"github.com/code-shoily/aocgo/year18/day02"
	"github.com/code-shoily/aocgo/year18/day03"
	"github.com/code-shoily/aocgo/year18/day04"
	"github.com/code-shoily/aocgo/year18/day05"
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
	default:
		fmt.Printf("2018/%d has not been solved yet", day)
	}
}
