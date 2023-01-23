// Package day09 - Solution for Advent of Code 2017/09
// Problem Link: http://adventofcode.com/2017/day/09
package day09

import (
	_ "embed"
	"fmt"
)

//go:embed input.txt
var input []byte

// Run prints out the result of the solution.
func Run() {
	fmt.Println(solve(input))
}

func solve(input []byte) (int, int) {
	return solvePart1(input), solvePart2(input)
}

func solvePart1(stream []byte) (totalScore int) {
	var (
		currentGroup            int
		garbageMode, cancelMode bool
	)
	for _, char := range stream {
		switch {
		case cancelMode:
			cancelMode = false
		case char == '!' && !cancelMode:
			cancelMode = true
		case char == '<' && !cancelMode:
			garbageMode = true
		case char == '>' && !cancelMode:
			garbageMode = false
		case char == '{' && !garbageMode && !cancelMode:
			currentGroup++
			totalScore += currentGroup
		case char == '}' && !garbageMode && !cancelMode:
			currentGroup--
		}
	}
	return totalScore
}

func solvePart2(stream []byte) (totalGarbage int) {
	var garbageMode, cancelMode bool
	for _, char := range stream {
		switch {
		case cancelMode:
			cancelMode = false
		case char == '!' && !cancelMode:
			cancelMode = true
		case char == '<' && !garbageMode && !cancelMode:
			garbageMode = true
		case char == '>' && !cancelMode:
			garbageMode = false
		case !garbageMode:
		default:
			totalGarbage++
		}
	}
	return totalGarbage
}
