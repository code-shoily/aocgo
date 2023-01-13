package intcode

import "github.com/code-shoily/aocgo/utils"

type Program struct {
	memory  []int
	current int
	done    bool
}

// Step commits one iteration of a program chunk and increases the stack
func (p *Program) Step() {
	if p.memory[p.current] != 99 {
		chunk := p.memory[p.current : p.current+4]
		cmd, op1, op2, out := chunk[0], chunk[1], chunk[2], chunk[3]
		switch cmd {
		case 1:
			p.memory[out] = p.memory[op1] + p.memory[op2]
		case 2:
			p.memory[out] = p.memory[op1] * p.memory[op2]
		}

		p.current = p.current + 4
	} else {
		p.done = true
	}
}

// Run executes the program until 99 is met
func (p *Program) Run() {
	for !p.done {
		p.Step()
	}
}

// Output shows the output of the program
func (p *Program) Output() int {
	return p.memory[0]
}

// ProvideInitialParameters corrects the program at positions 1 and 2
func (p *Program) ProvideInitialParameters(noun int, verb int) {
	p.memory[1] = noun
	p.memory[2] = verb
}

// InitializeProgram initializes a program with int codes.
func InitializeProgram(input string) *Program {
	return &Program{memory: utils.SplitByInts(input, ",")}
}
