// Package day12 - Solution for Advent of Code 2015/12
// Problem Link: http://adventofcode.com/2015/day/12
package day12

import (
	_ "embed"
	"encoding/json"
	"fmt"
	"regexp"
	"strconv"
)

//go:embed input.txt
var input string

// Run prints out the result of the solution.
func Run() {
	fmt.Println(solve(input))
}

func solve(input string) (int, int) {
	return solvePart1(input), solvePart2(input)
}

func solvePart1(input string) (total int) {
	tokens := regexp.MustCompile(`(-?\d+)`).FindAllString(input, -1)
	for _, token := range tokens {
		if number, err := strconv.Atoi(token); err == nil {
			total += number
		}
	}
	return total
}

func solvePart2(input string) (total int) {
	if !regexp.MustCompile("[{}]").MatchString(input) ||
		!regexp.MustCompile("red").MatchString(input) {
		return solvePart1(input)
	}

	var (
		unmarshalledObject map[string]any
		unmarshalledList   []any
	)

	err := json.Unmarshal([]byte(input), &unmarshalledObject)

	if err != nil {
		err := json.Unmarshal([]byte(input), &unmarshalledList)
		if err != nil {
			panic(err)
		}

		for _, v := range unmarshalledList {
			str, err := json.Marshal(v)
			if err != nil {
				panic(err)
			}
			total += solvePart2(string(str))
		}

		return
	}

	for _, v := range unmarshalledObject {
		str, ok := v.(string)
		if ok && str == "red" {
			return 0
		}
	}

	for _, v := range unmarshalledObject {
		str, err := json.Marshal(v)
		if err != nil {
			panic(err)
		}
		total += solvePart2(string(str))
	}

	return
}
