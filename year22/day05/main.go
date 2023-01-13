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

type ContainerMap = map[int][]string
type MoveInstruction struct {
	quantity int
	from     int
	to       int
}

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
		toMove := make([]string, movement.quantity)
		for idx, container := range state[movement.from][:movement.quantity] {
			toMove[movement.quantity-idx-1] = container
		}

		state[movement.from] = state[movement.from][movement.quantity:]

		for i := len(toMove) - 1; i >= 0; i-- {
			state[movement.to] = append([]string{toMove[i]}, state[movement.to]...)
		}
	}

	return topContainers(state)
}

func solvePart2(input string) string {
	state, moves := parse(input)

	for _, movement := range moves {
		toMove := make([]string, movement.quantity)
		copy(toMove, state[movement.from][:movement.quantity])
		state[movement.from] = state[movement.from][movement.quantity:]
		state[movement.to] = append(toMove, state[movement.to]...)
	}

	return topContainers(state)
}

func parse(input string) (ContainerMap, []MoveInstruction) {
	sections := strings.Split(input, "\n\n")
	return parseContainerState(sections[0]), parseMovements(sections[1])
}

func getContainerId(token string) string {
	return token[1:2]
}

func parseContainerState(input string) ContainerMap {
	lines := strings.Split(input, "\n")
	width := len(lines[0])
	containerGrid := make([][]string, 0, len(lines))

	for _, line := range lines[:len(lines)-1] {
		var row []string

		for i := 0; i < width; i += 4 {
			if container := line[i : i+3]; container != "   " {
				row = append(row, getContainerId(container))
			} else {
				row = append(row, "_")
			}
		}

		containerGrid = append(containerGrid, row)
	}

	return toContainerMap(containerGrid)
}

func toContainerMap(grid [][]string) ContainerMap {
	containerState := make(ContainerMap)

	for _, column := range grid {
		idx := 1
		for _, container := range column {
			if container != "_" {
				containerState[idx] = append(containerState[idx], container)
			}
			idx++
		}
	}

	return containerState
}

func parseMovements(input string) []MoveInstruction {
	lines := strings.Split(input, "\n")
	movements := make([]MoveInstruction, 0, len(lines))

	for _, line := range lines {
		var quantity, from, to int
		fmt.Sscanf(line, "move %d from %d to %d", &quantity, &from, &to)
		movements = append(movements, MoveInstruction{quantity, from, to})
	}

	return movements
}

func topContainers(stack ContainerMap) string {
	output := make([]string, len(stack))
	for k, v := range stack {
		output[k-1] = v[0]
	}

	return strings.Join(output, "")
}
