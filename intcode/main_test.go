package intcode

import (
	"fmt"
	"reflect"
	"testing"
)

func TestProgram_InitializeProgram(t *testing.T) {
	examples := []struct {
		inputFeed     string
		startingValue int
		expect        []int
	}{
		{"1,2,3", 1, []int{1, 2, 3}},
		{"1", 2, []int{1}},
	}

	for _, example := range examples {
		name := fmt.Sprintf("testing for input %v", example.inputFeed)
		t.Run(name, func(tt *testing.T) {
			if program := InitializeProgram(example.inputFeed, example.startingValue); !reflect.DeepEqual(program.memory, example.expect) || program.startValue != example.startingValue {
				tt.Errorf("Fail - expected %v but got %v", example.expect, program.memory)
			}
		})
	}
}

func TestProgram_Step(t *testing.T) {
	examples := []struct {
		given  string
		expect []int
	}{
		{"1,0,0,0,99", []int{2, 0, 0, 0, 99}},
		{"2,3,0,3,99", []int{2, 3, 0, 6, 99}},
		{"2,4,4,5,99,0", []int{2, 4, 4, 5, 99, 9801}},
	}

	for _, example := range examples {
		program := InitializeProgram(example.given, 0)
		name := fmt.Sprintf("testing for input %v", example.given)
		t.Run(name, func(tt *testing.T) {
			program.Step()
			if !reflect.DeepEqual(example.expect, program.memory) {
				tt.Errorf("Fail - expected %v but got %v", example.expect, program.memory)
			}
		})
	}
}

func TestProgram_Run(t *testing.T) {
	examples := []struct {
		given  string
		expect []int
	}{
		{"1,0,0,0,99", []int{2, 0, 0, 0, 99}},
		{"2,3,0,3,99", []int{2, 3, 0, 6, 99}},
		{"2,4,4,5,99,0", []int{2, 4, 4, 5, 99, 9801}},
		{"1,1,1,4,99,5,6,0,99", []int{30, 1, 1, 4, 2, 5, 6, 0, 99}},
		{"1,9,10,3,2,3,11,0,99,30,40,50", []int{3500, 9, 10, 70, 2, 3, 11, 0, 99, 30, 40, 50}},
	}

	for _, example := range examples {
		program := InitializeProgram(example.given, 0)
		name := fmt.Sprintf("testing for input %v", example.given)
		t.Run(name, func(tt *testing.T) {
			program.Run()
			if !reflect.DeepEqual(example.expect, program.memory) {
				tt.Errorf("Fail - expected %v but got %v", example.expect, program.memory)
			}
		})
	}
}

func TestProgram_ProvideInitialParameters(t *testing.T) {
	examples := []struct {
		given  string
		noun   int
		verb   int
		expect []int
	}{
		{"1,0,0,0,99", 12, 2, []int{1, 12, 2, 0, 99}},
		{"2,3,0,3,99", 18, 9, []int{2, 18, 9, 3, 99}},
		{"2,4,4,5,99,0", 77, 0, []int{2, 77, 0, 5, 99, 0}},
	}

	for _, example := range examples {
		program := InitializeProgram(example.given, 0)
		name := fmt.Sprintf("testing for input %v", example.given)
		t.Run(name, func(tt *testing.T) {
			program.ProvideInitialParameters(example.noun, example.verb)
			if !reflect.DeepEqual(example.expect, program.memory) {
				tt.Errorf("Fail - expected %v but got %v", example.expect, program.memory)
			}
		})
	}
}

func TestProgram_Output(t *testing.T) {
	examples := []struct {
		given  string
		expect int
	}{
		{"1,0,0,0,99", 1},
		{"2,3,0,3,99", 2},
		{"2,4,4,5,99,0", 2},
		{"2808,4,4,5,99,0", 2808},
		{"30,1,1,4,2,5,6,0,99", 30},
	}

	for _, example := range examples {
		program := InitializeProgram(example.given, 0)
		name := fmt.Sprintf("testing for input %v", example.given)
		t.Run(name, func(tt *testing.T) {
			if got := program.Output(); example.expect != got {
				tt.Errorf("Fail - expected %v but got %v", example.expect, got)
			}
		})
	}
}

func TestParseParameterMode(t *testing.T) {
	type Example struct {
		instruction int
		opCode      int
		imMode1     bool
		imMode2     bool
	}

	examples := []Example{
		{1, 1, false, false},
		{2, 2, false, false},
		{101, 1, true, false},
		{1002, 2, false, true},
		{1101, 1, true, true},
	}

	assertAll := func(example Example, opCode int, pMode1 bool, pMode2 bool) bool {
		return opCode == example.opCode && pMode1 == example.imMode1 && pMode2 == example.imMode2
	}

	for _, example := range examples {
		name := fmt.Sprintf("testing for input %v", example.instruction)
		t.Run(name, func(tt *testing.T) {
			if opCode, pMode1, pMode2 := parseParameterMode(example.instruction); !assertAll(example, opCode, pMode1, pMode2) {
				tt.Errorf("Fail - expected %v/%v/%v but got %v/%v/%v",
					example.opCode,
					example.imMode1,
					example.imMode2,
					opCode,
					pMode1,
					pMode2)
			}
		})
	}
}
