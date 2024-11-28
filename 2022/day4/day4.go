package day4

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/steffantucker/AdventOfCode/2022/utils"
)

func Run() {
	pairs := utils.GetStringList(2022, 4)
	p1 := part1(pairs)
	fmt.Printf("Part 1: %v\n", p1)
	p2 := part2(pairs)
	fmt.Printf("Part 2: %v\n", p2)
}

func part1(pairs []string) (overlaps int) {
	for _, areas := range pairs {
		area1, area2, _ := strings.Cut(areas, ",")
		a1start, a1end := areaNumbers(area1)
		a2start, a2end := areaNumbers(area2)
		if (a1start <= a2start && a1end >= a2end) || (a1start >= a2start && a1end <= a2end) {
			overlaps++
			continue
		}
	}
	return
}

func areaNumbers(area string) (int, int) {
	start, end, _ := strings.Cut(area, "-")
	startn, _ := strconv.Atoi(start)
	endn, _ := strconv.Atoi(end)
	return startn, endn
}

func part2(pairs []string) (overlaps int) {
	for _, areas := range pairs {
		area1, area2, _ := strings.Cut(areas, ",")
		a1start, a1end := areaNumbers(area1)
		a2start, a2end := areaNumbers(area2)
		if ((a1start >= a2start && a1start <= a2end) || (a1end >= a2start && a1end <= a2end)) ||
			((a2start >= a1start && a2start <= a1end) || (a2end >= a1start && a2end <= a1end)) {
			overlaps++
			continue
		}
	}
	return
}
