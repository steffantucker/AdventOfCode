package day1

import (
	"fmt"
	"math"
	"slices"

	utils "github.com/steffantucker/AdventOfCode/utils/go"
)

func Run() {
	lines := utils.GetNumberMatrix(2024, 1)
	p1 := part1(lines)
	fmt.Printf("Part 1: %v\n", p1)
	p2 := part2(lines)
	fmt.Printf("Part 2: %v\n", p2)
}

func part1(lines [][]int) (dist int) {
	left, right := []int{}, []int{}
	for _, row := range lines {
		left = append(left, row[0])
		right = append(right, row[1])
	}
	slices.Sort(left)
	slices.Sort(right)
	for i := 0; i < len(left); i++ {
		dist += int(math.Abs(float64(left[i] - right[i])))
	}
	return
}

func part2(lines [][]int) (simScore int) {
	left := []int{}
	right := make(map[int]int)
	for _, row := range lines {
		left = append(left, row[0])
		right[row[1]]++
	}
	for _, v := range left {
		simScore += v * right[v]
	}
	return
}
