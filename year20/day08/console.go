package day08

import (
	"strconv"
	"strings"
)

type instruction struct {
	cmd   string
	value int
}

func (i *instruction) swap() {
	switch i.cmd {
	case "jmp":
		i.cmd = "nop"
	case "nop":
		i.cmd = "jmp"
	}
}

func newInstruction(line string) instruction {
	tokens := strings.Split(line, " ")

	if value, err := strconv.Atoi(tokens[1]); err == nil {
		return instruction{tokens[0], value}
	}

	panic("[logical error] unreachable zone")
}

type console struct {
	instructions       []instruction
	instructionLog     map[int]bool
	instructionPointer int
	accumulator        int
	loop               bool
	completed          bool
}

func (c *console) run() {
	for !c.loop && !c.completed {
		c.step()
	}
}

func (c *console) step() {
	if _, loop := c.instructionLog[c.instructionPointer]; loop {
		c.loop = true
	} else if c.instructionPointer < -1 || c.instructionPointer >= len(c.instructions) {
		c.completed = true
	} else {
		c.instructionLog[c.instructionPointer] = true
		currentInstruction := c.instructions[c.instructionPointer]

		switch currentInstruction.cmd {
		case "nop":
			c.instructionPointer++
		case "acc":
			c.accumulator += currentInstruction.value
			c.instructionPointer++
		case "jmp":
			c.instructionPointer += currentInstruction.value
		}
	}
}

func newConsole(instructions []instruction) *console {
	return &console{instructions: instructions, instructionLog: make(map[int]bool)}
}
