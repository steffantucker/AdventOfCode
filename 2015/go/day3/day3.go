package day3

import (
	"fmt"
	"strings"

	utils "github.com/steffantucker/AdventOfCode/utils/go"
)

func Run() {
	route, _ := utils.GetInput(2015, 3)
	p1 := part1(route)
	fmt.Printf("Part 1: %v\n", p1)
	p2 := part2(route)
	fmt.Printf("Part 2: %v\n", p2)
}

func part1(dirs string) int {
	x, y := 0, 0
	visited := map[string]int{"0,0": 1}
	for _, d := range strings.Split(dirs, "") {
		switch d {
		case "^":
			y++
		case ">":
			x++
		case "<":
			x--
		case "v":
			y--
		}
		visited[fmt.Sprintf("%v,%v", x, y)] += 1
	}
	return len(visited)
}

func part2(dirs string) int {
	santas := []struct{ X, Y int }{{0, 0}, {0, 0}}
	visited := map[string]int{"0,0": 2}
	for i, d := range strings.Split(dirs, "") {
		switch d {
		case "^":
			santas[i%2].Y++
		case ">":
			santas[i%2].X++
		case "<":
			santas[i%2].X--
		case "v":
			santas[i%2].Y--
		}
		visited[fmt.Sprintf("%v,%v", santas[i%2].X, santas[i%2].Y)] += 1
	}
	return len(visited)
}
