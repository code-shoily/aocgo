// Package day06 - Solution for Advent of Code 2020/06
// Problem Link: http://adventofcode.com/2020/day/06
package day06

import (
	_ "embed"
	"fmt"
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

func solvePart1(responses []GroupResponse) (total int) {
	for _, response := range responses {
		total += len(response.questionTally)
	}

	return total
}

func solvePart2(responses []GroupResponse) (total int) {
	for _, response := range responses {
		for _, tally := range response.questionTally {
			if tally == response.respondents {
				total++
			}
		}
	}

	return total
}

type GroupResponse struct {
	respondents   int
	questionTally map[rune]int
}

func parse(responses string) (data []GroupResponse) {
	for _, groupResponse := range strings.Split(responses, "\n\n") {
		data = append(data, makeGroupResponse(groupResponse))
	}

	return data
}

func makeGroupResponse(groupResponse string) GroupResponse {
	responses := strings.Split(groupResponse, "\n")
	respondents := len(responses)
	questionTally := make(map[rune]int)

	for _, response := range responses {
		for _, question := range response {
			questionTally[question]++
		}
	}

	return GroupResponse{respondents, questionTally}
}
