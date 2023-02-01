package monorail

import (
	"strconv"
	"strings"
)

type Instruction struct {
	argCount    int
	Cmd         string
	RegisterMap []bool
	Registers   []string
	Values      []int
}

func (instruction *Instruction) Toggle() {
	if instruction.argCount == 1 {
		if instruction.Cmd == "inc" {
			instruction.Cmd = "dec"
		} else {
			instruction.Cmd = "inc"
		}
	} else {
		if instruction.Cmd == "jnz" {
			instruction.Cmd = "cpy"
		} else {
			instruction.Cmd = "jnz"
		}
	}
}

func ParseInstruction(line string) *Instruction {
	tokens := strings.Split(line, " ")
	var (
		registerMap []bool
		registers   []string
		values      []int
	)
	for _, token := range tokens[1:] {
		value, isRegister := GetValue(token)
		registerMap = append(registerMap, isRegister) // TODO: This is registerMap

		if isRegister {
			registers = append(registers, token)
		} else {
			values = append(values, value)
		}
	}

	return &Instruction{len(tokens) - 1, tokens[0], registerMap, registers, values}

}

func GetValue(token string) (value int, isRegister bool) {
	if value, err := strconv.Atoi(token); err == nil {
		return value, false
	}
	return -1, true
}

type Computer struct {
	Registers    map[string]int
	Ptr          int
	Instructions []*Instruction
}

func (c *Computer) Run() {
	for c.Ptr < len(c.Instructions) {
		i := c.Instructions[c.Ptr]
		switch i.Cmd {
		case "inc":
			c.Ptr += c.incrementBy(i.Registers[0], 1)
		case "dec":
			c.Ptr += c.incrementBy(i.Registers[0], -1)
		case "cpy":
			c.Ptr += c.copy(i)
		case "tgl":
			c.Ptr += c.toggle(i)
		case "jnz":
			c.Ptr += c.jnz(i)
		}
	}
}

func (c *Computer) incrementBy(register string, value int) int {
	c.Registers[register] += value
	return 1
}

func (c *Computer) copy(i *Instruction) int {
	if i.RegisterMap[0] {
		copyValue := c.Registers[i.Registers[0]]
		c.Registers[i.Registers[1]] = copyValue
	} else {
		copyValue := i.Values[0]
		c.Registers[i.Registers[0]] = copyValue
	}
	return 1
}

func (c *Computer) jnz(i *Instruction) int {
	switch a, b := i.RegisterMap[0], i.RegisterMap[1]; {
	case a && b:
		if c.Registers[i.Registers[0]] != 0 {
			return c.Registers[i.Registers[1]]
		}
	case a:
		if c.Registers[i.Registers[0]] != 0 {
			return i.Values[0]
		}
	case b:
		if i.Values[0] != 0 {
			return c.Registers[i.Registers[0]]
		}

	default:
		if i.Values[0] != 0 {
			return i.Values[1]
		}
	}

	return 1
}

func (c *Computer) toggle(i *Instruction) int {
	var toggleStep int
	if i.RegisterMap[0] {
		toggleStep = c.Ptr + c.Registers[i.Registers[0]]
	} else {
		toggleStep = c.Ptr + i.Values[0]
	}

	if toggleStep < len(c.Instructions) {
		c.Instructions[toggleStep].Toggle()
	}

	return 1
}

func NewComputer(instructions []*Instruction, bootstrap map[string]int) *Computer {
	return &Computer{bootstrap, 0, instructions}
}
