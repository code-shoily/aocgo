// Package day25 - Solution for Advent of Code 2022/25
// Problem Link: http://adventofcode.com/2022/day/25
package day25

import (
	"container/list"
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

// Run prints out the result of the solution.
func Run() {
	fmt.Println(solve(input))
}

func solve(input string) string {
	var totalDecimal int
	for _, snafu := range parse(input) {
		totalDecimal += snafuToDecimal(snafu)
	}
	return decimalToSnafu(totalDecimal)
}

func snafuToDecimal(snafu string) (decimal int) {
	size := len(snafu)
	for idx, value := range snafu {
		multiplier := intPow(5, size-1-idx)
		switch value {
		case '-':
			decimal += -1 * multiplier
		case '=':
			decimal += -2 * multiplier
		default:
			decimal += int(value-'0') * multiplier
		}
	}
	return
}

func decimalToSnafu(decimal int) (snafu string) {
	return convertDecimalToSnafu(decimal, list.New())
}

func convertDecimalToSnafu(decimal int, snafuList *list.List) string {
	if decimal == 0 {
		return listToString(snafuList)
	}

	var nextNumber int
	var snafuChar string

	switch value := decimal % 5; value {
	case 3:
		nextNumber, snafuChar = (decimal+2)/5, "="
	case 4:
		nextNumber, snafuChar = (decimal+1)/5, "-"
	default:
		nextNumber, snafuChar = decimal/5, strconv.Itoa(value)
	}

	snafuList.PushFront(snafuChar)
	return convertDecimalToSnafu(nextNumber, snafuList)
}

func parse(input string) []string {
	return strings.Split(input, "\n")
}

func intPow(base int, exp int) int {
	if exp == 0 {
		return 1
	}
	return base * intPow(base, exp-1)
}

func listToString(list *list.List) string {
	var slice []string
	for e := list.Front(); e != nil; e = e.Next() {
		slice = append(slice, e.Value.(string))
	}

	return strings.Join(slice, "")
}
