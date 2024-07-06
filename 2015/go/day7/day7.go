package day7

import (
	"fmt"
	"slices"
	"strings"

	utils "github.com/steffantucker/AdventOfCode/utils/go"
)

func Run() {
	instr := utils.GetStringList(2015, 7)
	p1 := part1(instr)
	fmt.Printf("Part 1: %v\n", p1)
	p2 := part2(instr)
	fmt.Printf("Part 2: %v\n", p2)
}

type Circuit struct {
	Wires       map[string]uint16
	Unconnected []string
}

func NewCircuit() Circuit {
	return Circuit{
		Wires:       map[string]uint16{},
		Unconnected: []string{},
	}
}

func (c *Circuit) GetWire(wire string) uint16 {
	return c.Wires[wire]
}

func (c *Circuit) SetWire(wire string, value uint16) {
	c.Wires[wire] = value
}

// signal
func (c *Circuit) processSignal(instr string) {
	i := strings.Split(instr, " ")
	if _, ok := c.Wires[i[2]]; ok {
		return
	}
	var signal uint16
	_, err := fmt.Sscan(i[0], &signal)
	if err == nil {
		c.Wires[i[2]] = signal
		return
	}
	if v, ok := c.Wires[i[0]]; ok {
		c.Wires[i[2]] = v
		return
	}
	c.Unconnected = append(c.Unconnected, instr)
}

// not
func (c *Circuit) processNot(instr string) {
	i := strings.Split(instr, " ")
	if val, ok := c.Wires[i[1]]; ok {
		c.Wires[i[3]] = ^val
	} else {
		c.Unconnected = append(c.Unconnected, instr)
	}
}

// or
func (c *Circuit) processOr(instr string) {
	i := strings.Split(instr, " ")
	a, aok := c.Wires[i[0]]
	b, bok := c.Wires[i[2]]
	if aok && bok {
		c.Wires[i[4]] = a | b
	} else {
		c.Unconnected = append(c.Unconnected, instr)
	}
}

// and
func (c *Circuit) processAnd(instr string) {
	i := strings.Split(instr, " ")
	a, aok := c.Wires[i[0]]
	b, bok := c.Wires[i[2]]
	if i[0] == "1" && bok {
		c.Wires[i[4]] = 1 & b
	} else if aok && bok {
		c.Wires[i[4]] = a & b
	} else {
		c.Unconnected = append(c.Unconnected, instr)
	}
}

// rshift
func (c *Circuit) processRshift(instr string) {
	i := strings.Split(instr, " ")
	if v, ok := c.Wires[i[0]]; ok {
		var shiftvalue uint16
		fmt.Sscan(i[2], &shiftvalue)
		c.Wires[i[4]] = v >> shiftvalue
	} else {
		c.Unconnected = append(c.Unconnected, instr)
	}
}

// lshift
func (c *Circuit) processLshift(instr string) {
	i := strings.Split(instr, " ")
	if v, ok := c.Wires[i[0]]; ok {
		var shiftvalue uint16
		fmt.Sscan(i[2], &shiftvalue)
		c.Wires[i[4]] = v << shiftvalue
	} else {
		c.Unconnected = append(c.Unconnected, instr)
	}
}

func (c *Circuit) processOp(i string) {
	if strings.Contains(i, "NOT") {
		c.processNot(i)
	} else if strings.Contains(i, "AND") {
		c.processAnd(i)
	} else if strings.Contains(i, "OR") {
		c.processOr(i)
	} else if strings.Contains(i, "RSHIFT") {
		c.processRshift(i)
	} else if strings.Contains(i, "LSHIFT") {
		c.processLshift(i)
	} else {
		c.processSignal(i)
	}
}

func (c *Circuit) Run(instr []string) {
	for _, i := range instr {
		c.processOp(strings.TrimSpace(i))
	}
	for len(c.Unconnected) > 0 {
		op := c.Unconnected[0]
		c.Unconnected = c.Unconnected[1:]
		c.processOp(op)
	}
}

func part1(instr []string) uint16 {
	circuit := NewCircuit()
	slices.SortFunc[[]string](instr, func(a, b string) int { return len(a) - len(b) })
	circuit.Run(instr)
	return circuit.GetWire("a")
}

func part2(instr []string) uint16 {
	circuit := NewCircuit()
	slices.SortFunc[[]string](instr, func(a, b string) int { return len(a) - len(b) })
	circuit.Run(instr)
	a := circuit.GetWire("a")
	circuit = NewCircuit()
	circuit.SetWire("b", a)
	circuit.Run(instr)
	return circuit.GetWire("a")
}
