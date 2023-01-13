// Package day05 - Solution for Advent of Code 2022/05
// Problem Link: http://adventofcode.com/2022/day/05
package day05

import (
	_ "embed"
	"fmt"
	"strings"
)

//go:embed input.txt
var input string

type ContainerState = map[int][]string
type Move struct {
	quantity int
	from     int
	to       int
}

var empty = "   "
var emptyMarker = "_"

// Run prints out the result of the solution.
func Run() {
	fmt.Println(solve(input))
}

func solve(input string) (string, string) {
	return solvePart1(input), solvePart2(input)
}

func solvePart1(input string) string {
	state, moves := parse(input)

	for _, movement := range moves {
		containers := make([]string, movement.quantity)
		for idx, container := range state[movement.from][:movement.quantity] {
			containers[movement.quantity-idx-1] = container
		}

		state[movement.from] = state[movement.from][movement.quantity:]
		state[movement.to] = append(containers, state[movement.to]...)
	}

	return topContainers(state)
}

func solvePart2(input string) string {
	state, moves := parse(input)

	for _, movement := range moves {
		containers := make([]string, movement.quantity)
		copy(containers, state[movement.from][:movement.quantity])
		state[movement.from] = state[movement.from][movement.quantity:]
		state[movement.to] = append(containers, state[movement.to]...)
	}

	return topContainers(state)
}

func parse(input string) (ContainerState, []Move) {
	sections := strings.Split(input, "\n\n")
	return parseContainerState(sections[0]), parseMoves(sections[1])
}

func getContainerId(token string) string {
	return token[1:2]
}

func parseContainerState(input string) ContainerState {
	lines := strings.Split(input, "\n")
	width := len(lines[0])
	containerGrid := make([][]string, 0, len(lines))

	for _, line := range lines[:len(lines)-1] {
		var level []string

		for i := 0; i < width; i += 4 {
			if container := line[i : i+3]; container != empty {
				level = append(level, getContainerId(container))
			} else {
				level = append(level, emptyMarker)
			}
		}

		containerGrid = append(containerGrid, level)
	}

	return asContainerState(containerGrid)
}

func asContainerState(grid [][]string) ContainerState {
	containerState := make(ContainerState)

	for _, column := range grid {
		idx := 1
		for _, container := range column {
			if container != emptyMarker {
				containerState[idx] = append(containerState[idx], container)
			}
			idx++
		}
	}

	return containerState
}

func parseMoves(input string) []Move {
	lines := strings.Split(input, "\n")
	moves := make([]Move, 0, len(lines))

	for _, line := range lines {
		var quantity, from, to int
		fmt.Sscanf(line, "move %d from %d to %d", &quantity, &from, &to)
		moves = append(moves, Move{quantity, from, to})
	}

	return moves
}

func topContainers(state ContainerState) string {
	output := make([]string, len(state))
	for k, v := range state {
		output[k-1] = v[0]
	}

	return strings.Join(output, "")
}
