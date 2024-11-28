package day5

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/steffantucker/AdventOfCode/2022/utils"
)

type Stacks map[int][]rune

func Run() {
	instructions := utils.GetStringList(2022, 5)
	p1 := run(instructions, part1)
	fmt.Printf("Part 1: %v\n", p1)
	p2 := run(instructions, part2)
	fmt.Printf("Part 2: %v\n", p2)
}

func run(instructions []string, moveFunc func(int, int, int, Stacks) Stacks) string {
	stacks := makeStacks(instructions)
	moves := getMoves(instructions)
	for _, move := range moves {
		count, from, to := getMoveInfo(move)
		stacks = moveFunc(count, from, to, stacks)
	}
	output := ""
	for i := 1; ; i++ {
		crate, ok := stacks[i]
		if !ok {
			break
		}
		output += string(crate[len(crate)-1])
	}
	return output
}

func part1(count, from, to int, stacks Stacks) Stacks {
	for i := 0; i < count; i++ {
		l := len(stacks[from]) - 1
		x := ' '
		x, stacks[from] = stacks[from][l], stacks[from][:l]
		stacks[to] = append(stacks[to], x)
	}
	return stacks
}

func part2(count, from, to int, stacks Stacks) Stacks {
	l := len(stacks[from]) - count
	var x []rune
	x, stacks[from] = stacks[from][l:], stacks[from][:l]
	stacks[to] = append(stacks[to], x...)
	return stacks
}

func makeStacks(layers []string) Stacks {
	stacks := make(Stacks)
	for _, layer := range layers {
		if strings.Contains(layer, "1") {
			break
		}
		counter := 0
		for i, a := range layer {
			if i%4 == 1 {
				counter++
				if a != ' ' {
					stacks[counter] = append([]rune{a}, stacks[counter]...)
				}
			}
		}
	}
	return stacks
}

func getMoves(strings []string) []string {
	for i, a := range strings {
		if a == "" {
			return strings[i+1:]
		}
	}
	return []string{}
}

func getMoveInfo(move string) (int, int, int) {
	r := regexp.MustCompile(`move (\d+) from (\d+) to (\d+)`)
	matches := r.FindStringSubmatch(move)
	count, _ := strconv.Atoi(matches[1])
	from, _ := strconv.Atoi(matches[2])
	to, _ := strconv.Atoi(matches[3])
	return count, from, to
}
