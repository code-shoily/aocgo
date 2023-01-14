// Package day05 - Solution for Advent of Code 2020/05
// Problem Link: http://adventofcode.com/2020/day/05
package day05

import (
	_ "embed"
	"fmt"
	"github.com/code-shoily/aocgo/utils"
	"sort"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

// Run prints out the result of the solution.
func Run() {
	fmt.Println(solve(input))
}

func solve(input string) (int, int) {
	data := parse(input)
	return solvePart1(data), solvePart2(data)
}

func solvePart1(boardingPasses [][]string) (maxSeatID int) {
	for _, boardingPass := range boardingPasses {
		if seatID := getSeatID(boardingPass[0], boardingPass[1]); seatID > maxSeatID {
			maxSeatID = seatID
		}
	}

	return maxSeatID
}

func solvePart2(boardingPasses [][]string) int {
	seats := make([]int, 0, len(boardingPasses))
	for _, boardingPass := range boardingPasses {
		seatID := getSeatID(boardingPass[0], boardingPass[1])
		seats = append(seats, seatID)
	}

	sort.Ints(seats)

	for i := 0; i < len(seats)-1; i++ {
		if seats[i+1]-seats[i] == 2 {
			return seats[i] + 1
		}
	}
	panic("someone took my seat")
}

func parse(input string) (data [][]string) {
	for _, boardingPass := range utils.SplitLines(input) {
		splitAt := len(boardingPass) - 3
		rowCol := []string{
			toBinary(boardingPass[:splitAt]),
			toBinary(boardingPass[splitAt:]),
		}
		data = append(data, rowCol)
	}

	return data
}

func toBinary(s string) string {
	replacer := strings.NewReplacer("F", "0", "B", "1", "R", "1", "L", "0")
	return replacer.Replace(s)
}

func getSeatID(rowBinary, colBinary string) int {
	row, _ := strconv.ParseInt(rowBinary, 2, 64)
	col, _ := strconv.ParseInt(colBinary, 2, 64)

	return int(row*8 + col)
}
