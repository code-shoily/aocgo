package day07

import (
	"strconv"
	"strings"
)

type signal uint16

type Wire struct {
	size      int
	target    string
	operation string
	op1Val    signal
	op2Val    signal
	value     signal
	mapping   map[string]int
	children  []string
}

func (wire *Wire) upgradeToAssignment(op1 string) {
	if value, err := strconv.Atoi(op1); err == nil {
		wire.value = signal(value)
		return
	}
	wire.operation = "ASSIGN"
	wire.mapping[op1] = 1
}

func (wire *Wire) upgradeToBinary(op1, op2, op string) {
	wire.upgradeToUnary(op1, op)

	if op2Val, err := strconv.Atoi(op2); err == nil {
		wire.op2Val = signal(op2Val)
		return
	}
	wire.mapping[op2] = 2
}

func (wire *Wire) upgradeToUnary(op1, op string) {
	wire.operation = op

	if op1Val, err := strconv.Atoi(op1); err == nil {
		wire.op1Val = signal(op1Val)
		return
	}
	wire.mapping[op1] = 1
}

func (wire *Wire) evaluate() (evaluated bool) {
	if len(wire.mapping) == 0 {
		switch wire.size {
		case 1:
			wire.value = wire.op1Val
		case 2:
			wire.value = ^wire.op1Val
		case 3:
			wire.binaryOp()
		}
		return true
	}
	return
}

func (wire *Wire) binaryOp() {
	switch strings.TrimSpace(wire.operation) {
	case "AND":
		wire.value = wire.op1Val & wire.op2Val
	case "OR":
		wire.value = wire.op1Val | wire.op2Val
	case "LSHIFT":
		wire.value = wire.op1Val << wire.op2Val
	case "RSHIFT":
		wire.value = wire.op1Val >> wire.op2Val
	}
}

func (wire *Wire) receive(opVar string, value signal) {
	switch wire.mapping[opVar] {
	case 1:
		wire.op1Val = value
	case 2:
		wire.op2Val = value
	}
	delete(wire.mapping, opVar)
}
