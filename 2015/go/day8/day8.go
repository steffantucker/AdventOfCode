package day8

import (
	"fmt"
	"strconv"

	utils "github.com/steffantucker/AdventOfCode/utils/go"
)

func Run() {
	list := utils.GetStringList(2015, 8)
	p1 := part1(list)
	fmt.Printf("Part 1: %v\n", p1)
	p2 := part2(list)
	fmt.Printf("Part 2: %v\n", p2)
}

func part1(list []string) (sum int) {
	for i, line := range list {
		l, err := strconv.Unquote(line)
		if err != nil {
			fmt.Printf("%v: %v\n", i, err)
			continue
		}
		sum += len(line) - len(l)
	}
	return
}

func part2(list []string) (sum int) {
	for _, line := range list {
		sum += len(strconv.Quote(line)) - len(line)
	}
	return
}
