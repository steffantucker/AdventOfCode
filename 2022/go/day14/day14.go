package day14

import (
	"fmt"
	"strings"

	"github.com/steffantucker/AdventOfCode/2022/utils"
)

func Run() {
	pairs := utils.GetStringList(2022, 14)
	p1 := part1(pairs)
	fmt.Printf("Part 1: %v\n", p1)
	p2 := part2(pairs)
	fmt.Printf("Part 2: %v\n", p2)
}

func part1(lines []string) (sand int) {
	cavern := buildCavern(lines)
	return
}

func buildCavern(lines []string) utils.Grid {
	cavern := utils.ArrayGrid{}
	for _, line := range lines {
		for _, cs := range strings.Split(line, " -> ") {

		}
	}
	return cavern
}

func part2(pairs []string) (overlaps int) {
	return
}
