package day1

import (
	"fmt"
	"log"
	"strings"

	utils "github.com/steffantucker/AdventOfCode/utils/go"
)

func Run() {
	input, err := utils.GetInput(2015, 1)
	if err != nil {
		log.Fatalf("Failed getting input %#v\n", err)
	}

	p1 := part1(input)
	fmt.Printf("Part 1: %v\n", p1)
	p2 := part2(input)
	fmt.Printf("Part 2: %v\n", p2)
}

func part1(parens string) int {
	left, right := 0, 0
	for _, p := range strings.Split(parens, "") {
		switch p {
		case "(":
			left++
		case ")":
			right++
		}
	}
	return left - right
}

func part2(parens string) int {
	left, right := 0, 0
	for i, p := range strings.Split(parens, "") {
		switch p {
		case "(":
			left++
		case ")":
			right++
		}
		if (left - right) == -1 {
			return i + 1
		}
	}
	return 0
}
