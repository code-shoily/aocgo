package intcode

import (
	"fmt"
	"github.com/code-shoily/aocgo/utils"
	"strconv"
)

type Program struct {
	memory     []int
	current    int
	done       bool
	startValue int
}

// Step commits one iteration of a program chunk and increases the stack
func (p *Program) Step() {
	if opCode := p.getOpCode(); opCode != 99 {
		switch opCode {
		case 1, 2, 102, 1002, 1102, 101, 1001, 1101:
			doBinaryOp(p, opCode)
		case 3:
			doStore(p)
		case 4, 104:
			doOutput(p, opCode)
		case 5, 6, 7, 8, 105, 106, 107, 108, 1005, 1006, 1007, 1008, 1105, 1106, 1107, 1108:
			doBinaryOp(p, opCode)
		default:
			panic(fmt.Errorf("OpCode %d isn't implemented yet", opCode))
		}
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

func (p *Program) getOpCode() int {
	return p.memory[p.current]
}

// InitializeProgram initializes a program with int codes.
func InitializeProgram(input string, startValue int) *Program {
	return &Program{memory: utils.SplitByInts(input, ","), startValue: startValue}
}

// OPERATIONS
func doStore(p *Program) {
	address := p.memory[p.current+1]
	p.memory[address] = p.startValue
	p.current = p.current + 2
}

func doOutput(p *Program, opCodeMeta int) {
	_, imMode, _ := parseParameterMode(opCodeMeta)
	outputSource := p.memory[p.current+1]

	if !imMode {
		outputSource = p.memory[outputSource]
	}

	p.memory[0] = outputSource
	p.current = p.current + 2
}

func doBinaryOp(p *Program, opCodeMeta int) {
	opCode, imMode1, imMode2 := parseParameterMode(opCodeMeta)
	chunk := p.memory[p.current+1 : p.current+4]
	op1, op2, out := chunk[0], chunk[1], chunk[2]

	if !imMode1 {
		op1 = p.memory[op1]
	}
	if !imMode2 {
		op2 = p.memory[op2]
	}

	var jump bool

	switch opCode {
	case 1:
		p.memory[out] = op1 + op2
	case 2:
		p.memory[out] = op1 * op2
	case 5:
		jump = true
		if op1 != 0 {
			p.current = op2
		} else {
			p.current += 3
		}
	case 6:
		jump = true
		if op1 == 0 {
			p.current = op2
		} else {
			p.current += 3
		}
		return
	case 7:
		if op1 < op2 {
			p.memory[out] = 1
		} else {
			p.memory[out] = 0
		}
	case 8:
		if op1 == op2 {
			p.memory[out] = 1
		} else {
			p.memory[out] = 0
		}
	}

	if !jump {
		p.current += 4
	}
}

func parseParameterMode(instruction int) (int, bool, bool) {
	line := fmt.Sprintf("%04s", strconv.Itoa(instruction))
	imMode1 := line[1] == '1'
	imMode2 := line[0] == '1'
	if opCode, err := strconv.Atoi(line[2:]); err == nil {
		return opCode, imMode1, imMode2
	}
	panic("Invalid instruction")
}
