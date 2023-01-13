package year22

import (
	"fmt"
	"github.com/code-shoily/aocgo/year22/day01"
	"github.com/code-shoily/aocgo/year22/day02"
	"github.com/code-shoily/aocgo/year22/day03"
)

func SolveForDay(day int) {
	switch day {
	case 1:
		day01.Run()
	case 2:
		day02.Run()
	case 3:
		day03.Run()
	default:
		fmt.Printf("2022/%d has not been solved yet", day)
	}
}
