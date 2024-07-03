package day3

import (
	"fmt"
	"strings"

	utils "github.com/steffantucker/AdventOfCode/utils/go"
)

func Run() {
	packs := utils.GetStringList(2022, 3)
	p1 := part1(packs)
	fmt.Printf("Part 1: %v\n", p1)
	p2 := part2(packs)
	fmt.Printf("Part 2: %v\n", p2)
}

func part1(packs []string) (score int) {
	for _, pack := range packs {
		pocket1 := pack[:len(pack)/2]
		pocket2 := pack[len(pack)/2:]
		for _, item := range pocket1 {
			if strings.ContainsRune(pocket2, item) {
				score += priority(item)
				break
			}
		}
	}
	return
}

func priority(item rune) int {
	if item >= 97 {
		return int(item % 96)
	}
	return int(item%64) + 26
}

func part2(packs []string) (score int) {
	for i := 0; i < len(packs); i += 3 {
		for _, item := range packs[i] {
			if strings.ContainsRune(packs[i+1], item) && strings.ContainsRune(packs[i+2], item) {
				score += priority(item)
				break
			}
		}
	}
	return
}
