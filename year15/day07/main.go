// Package day07 - Solution for Advent of Code 2015/07
// Problem Link: http://adventofcode.com/2015/day/07
// FIXME: Use graphs.Graph data structure and algorithm
package day07

import (
	_ "embed"
	"fmt"
	"github.com/code-shoily/aocgo/algo"
	"strings"
)

//go:embed input.txt
var input string

// Run prints out the result of the solution.
func Run() {
	fmt.Println(solve(input))
}

func solve(input string) (part1 signal, part2 signal) {
	part1 = connect(input, map[string]signal{})
	part2 = connect(input, map[string]signal{"b": part1})
	return
}

func connect(input string, override map[string]signal) signal {
	evaluated, wires := parse(input, override)
	runDFS(evaluated, wires)
	return wires["a"].value
}

func parse(input string, override map[string]signal) (evaluated []Wire, wires map[string]*Wire) {
	wires = make(map[string]*Wire)

	for _, line := range strings.Split(input, "\n") {
		tokens := strings.Split(line, " -> ")
		op, a := tokens[1], tokens[0]
		wires[op] = newWire(a, op)
	}

	updateChildren(wires)

	for wire, overriddenSignal := range override {
		wires[wire].value = overriddenSignal
	}

	for _, evaluatedWire := range wires {
		if len(evaluatedWire.mapping) == 0 {
			evaluated = append(evaluated, *evaluatedWire)
		}
	}

	return
}

func updateChildren(wires map[string]*Wire) {
	for name, wire := range wires {
		for parent := range wire.mapping {
			parentWire := wires[parent]
			parentWire.children = append(parentWire.children, name)
		}
	}
}

func runDFS(evaluated []Wire, wires map[string]*Wire) {
	stack := algo.Stack[Wire]{}

	for _, wire := range evaluated {
		stack.Push(wire)
	}

	for !stack.IsEmpty() {
		current, _ := stack.Pop()

		for _, child := range current.children {
			childWire := wires[child]
			childWire.receive(current.target, current.value)

			if childWire.evaluate() {
				stack.Push(*childWire)
			}
		}
	}
}
