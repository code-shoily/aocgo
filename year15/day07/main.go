// Package day07 - Solution for Advent of Code 2015/07
// Problem Link: http://adventofcode.com/2015/day/07
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

func solve(input string) (signal, signal) {
	aSignal := solvePart1(input)
	return aSignal, solvePart2(input, aSignal)
}

func solvePart1(input string) signal {
	evaluated, wires := parse(input, map[string]signal{})
	runDFS(evaluated, wires)
	return wires["a"].value
}

func solvePart2(input string, override signal) signal {
	evaluated, wires := parse(input, map[string]signal{"b": override})
	runDFS(evaluated, wires)
	return wires["a"].value
}

func parse(input string, override map[string]signal) (evaluated []Wire, wires map[string]*Wire) {
	wires = make(map[string]*Wire)

	for _, line := range strings.Split(input, "\n") {
		tokens := strings.Split(line, " -> ")
		op, a := tokens[1], tokens[0]
		wires[op] = parseWire(a, op)
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

func parseWire(op string, target string) *Wire {
	tokens := strings.Split(op, " ")
	wire := Wire{size: len(tokens), mapping: make(map[string]int), target: target}

	switch len(tokens) {
	case 1:
		wire.upgradeToAssignment(tokens[0])
	case 3:
		wire.upgradeToBinary(tokens[0], tokens[2], tokens[1])
	case 2:
		wire.upgradeToUnary(tokens[1], tokens[0])
	}

	return &wire
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
