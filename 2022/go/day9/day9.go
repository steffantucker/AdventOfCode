package day9

import (
	"fmt"
	"math"
	"regexp"
	"strconv"

	"github.com/steffantucker/AdventOfCode/2022/go/utils"
)

func Run() {
	pairs := utils.GetStringList(2022, 9)
	p1 := part1(pairs)
	fmt.Printf("Part 1: %v\n", p1)
	p2 := part2(pairs)
	fmt.Printf("Part 2: %v\n", p2)
}

type Knot struct {
	x, y int
}

func part1(instructions []string) int {
	instReg := regexp.MustCompile(`(R|L|U|D) (\d+)`)
	head := Knot{x: 0, y: 0}
	tail := Knot{x: 0, y: 0}
	locations := map[Knot]bool{tail: true}
	for _, inst := range instructions {
		reged := instReg.FindStringSubmatch(inst)
		direction := reged[1]
		steps, _ := strconv.Atoi(reged[2])
		for ; steps > 0; steps-- {
			switch direction {
			case "U":
				head.y++
			case "R":
				head.x++
			case "D":
				head.y--
			case "L":
				head.x--
			}
			xdiff := math.Abs(float64(head.x - tail.x))
			ydiff := math.Abs(float64(head.y - tail.y))
			if xdiff > 1 || ydiff > 1 {
				if head.x < tail.x {
					tail.x -= 1
				}
				if head.x > tail.x {
					tail.x += 1
				}
				if head.y < tail.y {
					tail.y -= 1
				}
				if head.y > tail.y {
					tail.y += 1
				}
				locations[tail] = true
			}
		}
	}
	return len(locations)
}

func part2(instructions []string) (overlaps int) {

	instReg := regexp.MustCompile(`(R|L|U|D) (\d+)`)
	rope := []Knot{{x: 0, y: 0}, {x: 0, y: 0}, {x: 0, y: 0}, {x: 0, y: 0}, {x: 0, y: 0}, {x: 0, y: 0}, {x: 0, y: 0}, {x: 0, y: 0}, {x: 0, y: 0}, {x: 0, y: 0}}
	locations := map[Knot]bool{rope[0]: true}
	for _, inst := range instructions {
		reged := instReg.FindStringSubmatch(inst)
		direction := reged[1]
		steps, _ := strconv.Atoi(reged[2])
		for ; steps > 0; steps-- {
			switch direction {
			case "U":
				rope[0].y++
			case "R":
				rope[0].x++
			case "D":
				rope[0].y--
			case "L":
				rope[0].x--
			}
			for i := 1; i < len(rope); i++ {
				rope[i] = moveTail(rope[i-1], rope[i])
			}
			locations[rope[9]] = true
		}
	}
	return len(locations)
}

func moveTail(head, tail Knot) Knot {
	xdiff := math.Abs(float64(head.x - tail.x))
	ydiff := math.Abs(float64(head.y - tail.y))
	if xdiff > 1 || ydiff > 1 {
		if head.x < tail.x {
			tail.x -= 1
		}
		if head.x > tail.x {
			tail.x += 1
		}
		if head.y < tail.y {
			tail.y -= 1
		}
		if head.y > tail.y {
			tail.y += 1
		}
	}
	return tail
}
