package day2

import (
	"fmt"
	"log"

	"github.com/steffantucker/AdventOfCode/2022/go/utils"
)

var assumedScoreMap = map[string]int{
	"B X": 1, // paper rock
	"C Y": 2, // scissors paper
	"A Z": 3, // rock scissors

	"A X": 4, // rock rock
	"B Y": 5, // paper paper
	"C Z": 6, // scissors scissors

	"C X": 7, // scissors rock
	"A Y": 8, // rock paper
	"B Z": 9, // paper scissors
}

var trueScoreMap = map[string]int{
	"B X": 1,
	"C X": 2,
	"A X": 3,
	"A Y": 4,
	"B Y": 5,
	"C Y": 6,
	"C Z": 7,
	"A Z": 8,
	"B Z": 9,
}

func Run() {
	guide := utils.GetStringList(2022, 2)
	p1 := part1(guide)
	fmt.Printf("Part 1: %v\n", p1)
	p2 := part2(guide)
	fmt.Printf("Part 2: %v\n", p2)
}

func part1(guide []string) (score int) {
	for _, round := range guide {
		s, ok := assumedScoreMap[round]
		if !ok {
			log.Fatalf("Unexpected pair: %q\n", round)
		}
		score += s
	}
	return
}

func part2(guide []string) (score int) {
	for _, round := range guide {
		s, ok := trueScoreMap[round]
		if !ok {
			log.Fatalf("Unexpected pair: %q\n", round)
		}
		score += s
	}
	return
}
