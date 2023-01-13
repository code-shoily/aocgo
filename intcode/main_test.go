package intcode

import (
	"fmt"
	"reflect"
	"testing"
)

func TestProgram_InitializeProgram(t *testing.T) {
	examples := []struct {
		given  string
		expect []int
	}{
		{"1,2,3", []int{1, 2, 3}},
		{"1", []int{1}},
	}

	for _, example := range examples {
		name := fmt.Sprintf("testing for input %v", example.given)
		t.Run(name, func(tt *testing.T) {
			if program := InitializeProgram(example.given); !reflect.DeepEqual(program.memory, example.expect) {
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
		program := InitializeProgram(example.given)
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
		program := InitializeProgram(example.given)
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
		program := InitializeProgram(example.given)
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
		program := InitializeProgram(example.given)
		name := fmt.Sprintf("testing for input %v", example.given)
		t.Run(name, func(tt *testing.T) {
			if got := program.Output(); example.expect != got {
				tt.Errorf("Fail - expected %v but got %v", example.expect, got)
			}
		})
	}
}
