package main

import (
	"errors"
	"fmt"
	"github.com/code-shoily/aocgo/gen"
	"github.com/code-shoily/aocgo/year15"
	"github.com/code-shoily/aocgo/year16"
	"github.com/code-shoily/aocgo/year17"
	"github.com/code-shoily/aocgo/year18"
	"github.com/code-shoily/aocgo/year19"
	"github.com/code-shoily/aocgo/year20"
	"github.com/code-shoily/aocgo/year21"
	"github.com/code-shoily/aocgo/year22"
	"os"
	"strconv"
)

const currentYear = 2022

func validate(year, day int) error {
	if year < 2015 || year > currentYear {
		return errors.New("invalid year entered")
	}
	if day < 1 || day > 25 {
		return errors.New("invalid date entered")
	}
	return nil
}

func solve(year, day int) {
	switch year {
	case 2015:
		year15.SolveForDay(day)
	case 2016:
		year16.SolveForDay(day)
	case 2017:
		year17.SolveForDay(day)
	case 2018:
		year18.SolveForDay(day)
	case 2019:
		year19.SolveForDay(day)
	case 2020:
		year20.SolveForDay(day)
	case 2021:
		year21.SolveForDay(day)
	case 2022:
		year22.SolveForDay(day)
	}
}

func extractFromArgs(args []string) (cmd string, year int, day int) {
	if yearInt, error := strconv.Atoi(args[1]); error == nil {
		year = yearInt
	}

	if dayInt, error := strconv.Atoi(args[2]); error == nil {
		day = dayInt
	}

	return args[0], year, day
}

func main() {
	args := os.Args
	cmd, year, day := "gen", 2018, 4

	if len(args) == 4 {
		cmd, year, day = extractFromArgs(args[1:])
	}

	if error := validate(year, day); error != nil {
		fmt.Println(error.Error())
	} else {
		switch cmd {
		case "gen":
			gen.GenerateSources(year, day)
		case "solve":
			solve(year, day)
		}
	}
}
