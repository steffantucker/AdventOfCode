package day1

import (
	"fmt"

	"github.com/steffantucker/AdventOfCode/2022/utils"
)

func Run() {
	snacks := utils.GetGroupedNumberList(2022, 1)
	p1 := part1(snacks)
	fmt.Printf("Part 1: %v\n", p1)
	p2 := part2(snacks)
	fmt.Printf("Part 2: %v\n", p2)
}

func part1(snacks [][]int) int {
	calorieSums := calorieSums(snacks)
	return utils.SliceMax(calorieSums)
}

func part2(snacks [][]int) int {
	first, second, third := 0, 0, 0
	calorieSums := calorieSums(snacks)
	for _, sum := range calorieSums {
		if sum >= first {
			third = second
			second = first
			first = sum
		} else if sum >= second {
			third = second
			second = sum
		} else if sum >= third {
			third = sum
		}
	}
	fmt.Printf("Part 2:\n First: %v\nSecond: %v\n Third: %v\n", first, second, third)
	return first + second + third
}

func calorieSums(snacks [][]int) []int {
	calorieSums := make([]int, 0, 300)
	for _, group := range snacks {
		sum := 0
		for _, calorie := range group {
			sum += calorie
		}
		calorieSums = append(calorieSums, sum)
	}
	return calorieSums
}
