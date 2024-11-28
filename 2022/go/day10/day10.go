package day10

import (
	"fmt"
	"regexp"
	"strconv"

	"github.com/steffantucker/AdventOfCode/2022/go/utils"
)

func Run() {
	ops := utils.GetStringList(2022, 10)
	c := new(CPU)
	c.LoadOps(ops)
	c.Run()
	fmt.Printf("Part 1: %v\n", c.SignalStrength)
	fmt.Println("Part2:")
	for _, line := range c.CRT {
		fmt.Printf("%v\n", line)
	}
}

type CPU struct {
	X              int
	Cycles         int
	SignalStrength int
	CRT            []string

	opCounter  int
	addCounter int
	ops        []string
	vCounter   int
}

func (c *CPU) cycleCRT() {
	hPos := (c.Cycles - 1) % 40
	if hPos == 0 {
		c.vCounter++
	}
	if hPos-1 <= c.X && hPos+1 >= c.X {
		c.CRT[c.vCounter] += "#"
	} else {
		c.CRT[c.vCounter] += "."
	}
}

func (c *CPU) LoadOps(ops []string) {
	c.ops = ops
}

func (c *CPU) Run() {
	opReg := regexp.MustCompile(`(noop|addx) ?(-?\d*)`)
	c.Cycles = 1
	c.opCounter = 0
	c.addCounter = 0
	c.X = 1

	c.CRT = make([]string, 6)
	c.vCounter = -1
	for ; ; c.Cycles++ {
		if c.opCounter >= len(c.ops) {
			return
		}
		if c.Cycles == 20 || (c.Cycles-20)%40 == 0 {
			c.SignalStrength += c.Cycles * c.X
		}
		c.cycleCRT()
		operation := opReg.FindStringSubmatch(c.ops[c.opCounter])
		switch operation[1] {
		case "addx":
			c.addCounter++
			if c.addCounter == 2 {
				n, _ := strconv.Atoi(operation[2])
				c.X += n
				c.addCounter = 0
				c.opCounter++
			}
		case "noop":
			c.opCounter++
		}
	}
}

func part1(ops []string) int {
	c := new(CPU)
	c.LoadOps(ops)
	c.Run()
	return c.SignalStrength
}

func part2(ops []string) []string {
	c := new(CPU)
	c.LoadOps(ops)
	c.Run()
	return c.CRT
}
