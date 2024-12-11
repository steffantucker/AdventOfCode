package day11

import (
	"fmt"
	"math"
	"strings"

	utils "github.com/steffantucker/AdventOfCode/utils/go"
)

type Iteration []int

type Rocks struct {
	Count int
	Rocks Iteration
	cache map[utils.Coordinate]int
}

func NewRocks(line string) *Rocks {
	rocks := make(Iteration, 0, len(line))
	for _, d := range strings.Split(line, " ") {
		rocks = append(rocks, utils.MustAtoi(d))
	}
	return &Rocks{
		Rocks: rocks,
		cache: map[utils.Coordinate]int{},
	}
}

func (r *Rocks) iterate(steps int, stone int) (count int) {
	if steps == 0 {
		return 1
	}
	hash := utils.NewCoordinate(steps, stone)
	if v, ok := r.cache[hash]; ok {
		return v
	}
	if stone == 0 {
		count = r.iterate(steps-1, 1)
	} else if length := int(math.Log10(float64(stone))) + 1; length%2 == 0 {
		pow := int(math.Pow10(length / 2))
		left := stone / pow
		right := stone % pow
		count = r.iterate(steps-1, left) + r.iterate(steps-1, right)
	} else {
		count = r.iterate(steps-1, stone*2024)
	}
	r.cache[hash] = count
	return count
}

func (r *Rocks) Iterate(times int) {
	r.Count = 0
	for _, stone := range r.Rocks {
		r.Count += r.iterate(times, stone)
	}
}

func Run() {
	input := utils.GetString(2024, 11)
	p1 := part1(input)
	fmt.Printf("Part 1: %v\n", p1)
	p2 := part2(input)
	fmt.Printf("Part 2: %v\n", p2)
}

func part1(input string) (result int) {
	rocks := NewRocks(input)
	rocks.Iterate(25)
	return rocks.Count
}

func part2(input string) (result int) {
	rocks := NewRocks(input)
	rocks.Iterate(75)
	return rocks.Count
}
