package day3

import (
	"fmt"
	"regexp"
	"slices"

	utils "github.com/steffantucker/AdventOfCode/utils/go"
)

func Run() {
	input := utils.GetString(2024, 3)
	p1 := part1(input)
	fmt.Printf("Part 1: %v\n", p1)
	p2 := part2(input)
	fmt.Printf("Part 2: %v\n", p2)
}

func part1(input string) (total int) {
	mulRegex := regexp.MustCompile(`mul\((\d{1,3})\,(\d{1,3})\)`)
	matches := mulRegex.FindAllSubmatch([]byte(input), -1)
	for _, match := range matches {
		total += utils.MustAtoi(string(match[1])) * utils.MustAtoi(string(match[2]))
	}
	return
}

func part2(input string) (total int) {
	mulMatches := regexp.MustCompile(`mul\((\d{1,3})\,(\d{1,3})\)`).FindAllStringSubmatchIndex(input, -1)
	doMatches := regexp.MustCompile(`do\(\)`).FindAllStringSubmatchIndex(input, -1)
	dontMatches := regexp.MustCompile(`don't\(\)`).FindAllStringSubmatchIndex(input, -1)
	var instructions []instr
	instructions = append(instructions, parseMulMatches(mulMatches, input)...)
	instructions = append(instructions, parseDoDontMatches(doMatches, DO)...)
	instructions = append(instructions, parseDoDontMatches(dontMatches, DONT)...)
	slices.SortFunc(instructions, func(a, b instr) int { return a.location - b.location })
	do := true
	for _, in := range instructions {
		switch in.instr {
		case DO:
			do = true
		case DONT:
			do = false
		case MUL:
			if do {
				total += in.data
			}
		}
	}

	return
}

func parseMulMatches(matches [][]int, s string) []instr {
	instructions := make([]instr, len(matches)/6)
	for _, match := range matches {
		inst := instr{
			location: match[0],
			instr:    MUL,
		}
		n1 := utils.MustAtoi(s[match[2]:match[3]])
		n2 := utils.MustAtoi(s[match[4]:match[5]])
		inst.data = n1 * n2
		instructions = append(instructions, inst)
	}
	return instructions
}

func parseDoDontMatches(matches [][]int, t instrType) []instr {
	instructions := make([]instr, len(matches)/2)
	for _, match := range matches {
		inst := instr{
			location: match[0],
			instr:    t,
		}
		instructions = append(instructions, inst)
	}
	return instructions
}

type instrType int

const (
	MUL instrType = iota
	DO
	DONT
)

type instr struct {
	location int
	instr    instrType
	data     int
}
