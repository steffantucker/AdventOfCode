package day10

import (
	"fmt"
	"maps"

	utils "github.com/steffantucker/AdventOfCode/utils/go"
)

func Run() {
	input := utils.GetString(2024, 10)
	p1 := part1(input)
	fmt.Printf("Part 1: %v\n", p1)
	p2 := part2(input)
	fmt.Printf("Part 2: %v\n", p2)
}

type Trails struct {
	Topo *utils.GenericMapGrid[int]
}

func (t *Trails) score(c utils.Coordinate, next int) map[utils.Coordinate]bool {
	trails := make(map[utils.Coordinate]bool)
	dirs := t.Topo.FindAround(c, next, true)
	if len(dirs) == 0 {
		return trails
	}
	if next == 9 {
		for _, coord := range dirs {
			trails[coord] = true
		}
		return trails
	}
	for _, dir := range dirs {
		maps.Copy(trails, t.score(dir, next+1))
	}
	return trails
}

func (t *Trails) rating(c utils.Coordinate, next int) (rating int) {
	dirs := t.Topo.FindAround(c, next, true)
	if len(dirs) == 0 {
		return
	}
	if next == 9 {
		return len(dirs)
	}
	for _, dir := range dirs {
		rating += t.rating(dir, next+1)
	}
	return
}

func part1(input string) (result int) {
	topo := Trails{
		Topo: utils.NewGenericMapGrid(-1),
	}
	topo.Topo.FillFromString(input, "\n", "", func(s string) int {
		if s == "." {
			return -1
		}
		return utils.MustAtoi(s)
	})

	for _, zero := range topo.Topo.FindAll(0) {
		result += len(topo.score(zero, 1))
	}
	return
}

func part2(input string) (result int) {
	topo := Trails{
		Topo: utils.NewGenericMapGrid(-1),
	}
	topo.Topo.FillFromString(input, "\n", "", func(s string) int {
		if s == "." {
			return -1
		}
		return utils.MustAtoi(s)
	})

	for _, zero := range topo.Topo.FindAll(0) {
		result += topo.rating(zero, 1)
	}
	return
}
