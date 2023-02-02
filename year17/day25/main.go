// Package day25 - Solution for Advent of Code 2017/25
// Problem Link: http://adventofcode.com/2017/day/25
package day25

import (
	"container/list"
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

func solve(input string) (total int) {
	stm, iter := parse(input)
	for i := 0; i < iter; i++ {
		stm.run()
	}

	return stm.diagnosticChecksum()
}

func parse(input string) (*StateMachine, int) {
	var rawBlueprints []string
	for _, blueprint := range strings.Split(input, "\n\n") {
		rawBlueprints = append(rawBlueprints, blueprint)
	}

	header, blueprints := rawBlueprints[0], rawBlueprints[1:]
	state, n := parseHead(header)

	return newStateMachine(state, blueprints), n
}

func parseHead(head string) (state string, n int) {
	fmt.Sscanf(head,
		"Begin in state %s\nPerform a diagnostic checksum after %d steps", &state, &n)
	return trim(state), n
}

type StateMachine struct {
	state   string
	ruleMap map[string]*Blueprint
	ptr     *list.Element
	tape    *list.List
}

func (stm *StateMachine) run() {
	rule, value := stm.ruleMap[stm.state], stm.ptr.Value.(int)
	write, move, next := rule.write0, rule.move0, rule.next0

	if value == 1 {
		write, move, next = rule.write1, rule.move1, rule.next1
	}

	stm.ptr.Value = write

	if move == "right" {
		if stm.ptr.Next() == nil {
			stm.ptr = stm.tape.PushBack(0)
		} else {
			stm.ptr = stm.ptr.Next()
		}
	} else {
		if stm.ptr.Prev() == nil {
			stm.ptr = stm.tape.PushFront(0)
		} else {
			stm.ptr = stm.ptr.Prev()
		}
	}

	stm.state = next
}

func (stm *StateMachine) diagnosticChecksum() (checksum int) {
	for e := stm.tape.Front(); e != nil; e = e.Next() {
		checksum += e.Value.(int)
	}

	return
}

func newStateMachine(startState string, rules []string) *StateMachine {
	ruleMap := make(map[string]*Blueprint)
	parseBlueprints(rules, ruleMap)

	tape := list.New()
	ptr := tape.PushBack(0)

	return &StateMachine{startState, ruleMap, ptr, tape}
}

type Blueprint struct {
	write0 int
	move0  string
	next0  string
	write1 int
	move1  string
	next1  string
}

func parseBlueprints(blueprints []string, stateMap map[string]*Blueprint) {
	for _, description := range blueprints {
		state, blueprint := parseBlueprint(description)
		stateMap[state] = blueprint
	}
}

func parseBlueprint(rule string) (string, *Blueprint) {
	lines, state := strings.Split(rule, "\n"), ""

	fmt.Sscanf(lines[0], "In state %s", &state)
	write0, move0, next0 := blueprintData(lines[2], lines[3], lines[4])
	write1, move1, next1 := blueprintData(lines[6], lines[7], lines[8])

	return trim(state), &Blueprint{write0, move0, next0, write1, move1, next1}
}

func blueprintData(w, m, n string) (write int, move, next string) {
	fmt.Sscanf(w, "  - Write the value %d.", &write)
	fmt.Sscanf(m, "  - Move one slot to the %s", &move)
	fmt.Sscanf(n, "  - Continue with state %s", &next)

	return write, trim(move), trim(next)
}

func trim(s string) string {
	return strings.TrimFunc(s, func(r rune) bool {
		return r == '.' || r == ':'
	})
}
