package day6

import (
	"fmt"
	"log"
	"regexp"

	utils "github.com/steffantucker/AdventOfCode/utils/go"
)

type Grid struct {
	Grid map[string]bool
}

func NewGrid() Grid {
	return Grid{
		Grid: make(map[string]bool),
	}
}

func (g *Grid) Box(tx, ty, bx, by int, inst string) {
	for y := ty; y <= by; y++ {
		for x := tx; x <= bx; x++ {
			xy := fmt.Sprintf("%v,%v", x, y)
			switch inst {
			case "turn on":
				g.Grid[xy] = true
			case "turn off":
				g.Grid[xy] = false
			case "toggle":
				g.Grid[xy] = !g.Grid[xy]
			}
		}
	}
}

func (g *Grid) CountLit() (sum int) {
	for _, c := range g.Grid {
		if c {
			sum++
		}
	}
	return
}

type BrightnessGrid struct {
	Grid map[string]int
}

func NewBrightnessGrid() BrightnessGrid {
	return BrightnessGrid{
		Grid: make(map[string]int),
	}
}

func (g *BrightnessGrid) Box(tx, ty, bx, by int, inst string) {
	for y := ty; y <= by; y++ {
		for x := tx; x <= bx; x++ {
			xy := fmt.Sprintf("%v,%v", x, y)
			switch inst {
			case "turn on":
				g.Grid[xy] += 1
			case "turn off":
				g.Grid[xy] -= 1
			case "toggle":
				g.Grid[xy] += 2
			}
			if g.Grid[xy] < 0 {
				g.Grid[xy] = 0
			}
		}
	}
}

func (g *BrightnessGrid) CountLit() (sum int) {
	for _, c := range g.Grid {
		sum += c
	}
	return
}

func Run() {
	list := utils.GetStringList(2015, 6)
	p1 := part1(list)
	fmt.Printf("Part 1: %v\n", p1)
	p2 := part2(list)
	fmt.Printf("Part 2: %v\n", p2)
}

func part1(list []string) int {
	lights := NewGrid()
	instregex := regexp.MustCompile(`(turn on|turn off|toggle) (\d{1,3}),(\d{1,3}) through (\d{1,3}),(\d{1,3})`)
	for _, l := range list {
		m := instregex.FindStringSubmatch(l)
		if len(m) == 0 {
			log.Fatalf("%v no matches", l)
		}
		lights.Box(utils.MustAtoi(m[2]), utils.MustAtoi(m[3]), utils.MustAtoi(m[4]), utils.MustAtoi(m[5]), m[1])
	}
	return lights.CountLit()
}

func part2(list []string) int {
	lights := NewBrightnessGrid()
	instregex := regexp.MustCompile(`(turn on|turn off|toggle) (\d{1,3}),(\d{1,3}) through (\d{1,3}),(\d{1,3})`)
	for _, l := range list {
		m := instregex.FindStringSubmatch(l)
		if len(m) == 0 {
			log.Fatalf("%v no matches", l)
		}
		lights.Box(utils.MustAtoi(m[2]), utils.MustAtoi(m[3]), utils.MustAtoi(m[4]), utils.MustAtoi(m[5]), m[1])
	}
	return lights.CountLit()
}
