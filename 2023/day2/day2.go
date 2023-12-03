package day2

import (
	"fmt"
	"regexp"
	"slices"

	"github.com/steffantucker/AdventOfCode/2023/utils"
)

func Run() {
	lines := utils.GetStringList(2023, 2)
	p1 := part1(lines)
	fmt.Printf("Part 1: %v\n", p1)
	p2 := part2(lines)
	fmt.Printf("Part 2: %v\n", p2)
}

func part1(lines []string) (sum int) {
	wantedGame := map[string]int{
		"red":   12,
		"green": 13,
		"blue":  14,
	}
	idRegex := regexp.MustCompile(`Game (\d*):`)
	redRegex := regexp.MustCompile(`(\d*?) red`)
	greenRegex := regexp.MustCompile(`(\d*?) green`)
	blueRegex := regexp.MustCompile(`(\d*?) blue`)
	for _, line := range lines {
		id := idRegex.FindStringSubmatch(line)[1]
		redMax := colorMax(redRegex.FindAllString(line, -1))
		greenMax := colorMax(greenRegex.FindAllString(line, -1))
		blueMax := colorMax(blueRegex.FindAllString(line, -1))
		if redMax <= wantedGame["red"] && greenMax <= wantedGame["green"] && blueMax <= wantedGame["blue"] {
			sum += utils.MustAtoi(id)
		}
	}
	return
}

func colorMax(v []string) (max int) {
	numRegex := regexp.MustCompile(`(\d*)`)
	intSlice := []int{}
	for _, value := range v {
		intSlice = append(intSlice, utils.MustAtoi(numRegex.FindString(value)))
	}
	return slices.Max(intSlice)
}

func part2(lines []string) (sum int) {
	redRegex := regexp.MustCompile(`(\d*?) red`)
	greenRegex := regexp.MustCompile(`(\d*?) green`)
	blueRegex := regexp.MustCompile(`(\d*?) blue`)
	for _, line := range lines {
		redMax := colorMax(redRegex.FindAllString(line, -1))
		greenMax := colorMax(greenRegex.FindAllString(line, -1))
		blueMax := colorMax(blueRegex.FindAllString(line, -1))
		sum += redMax * greenMax * blueMax
	}
	return
}
