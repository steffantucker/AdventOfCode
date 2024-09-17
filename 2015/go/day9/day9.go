package day9

import (
	"fmt"
	"strings"

	utils "github.com/steffantucker/AdventOfCode/utils/go"
)

type City struct {
	Name string
	Dist int
}

func Run() {
	list := utils.GetStringList(2015, 9)
	p1 := part1(list)
	fmt.Printf("Part 1: %v\n", p1)
	// p2 := part2(list)
	// fmt.Printf("Part 2: %v\n", p2)
}

func Walk(cities map[string][]City)

func part1(list []string) int {
	cities := make(map[string][]City)
	for _, l := range list {
		from, to, dist := parseLine(strings.TrimSpace(l))
		cities[from] = append(cities[from], City{Name: to, Dist: dist})
	}
	return 0
}

func parseLine(line string) (string, string, int) {
	l := strings.Split(line, " ")
	from, to := l[0], l[2]
	dist := utils.MustAtoi(l[4])
	return from, to, dist
}
