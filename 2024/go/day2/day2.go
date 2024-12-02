package day2

import (
	"fmt"
	"math"

	utils "github.com/steffantucker/AdventOfCode/utils/go"
)

func Run() {
	pairs := utils.GetNumberMatrix(2024, 2)
	p1 := part1(pairs)
	fmt.Printf("Part 1: %v\n", p1)
	p2 := part2(pairs)
	fmt.Printf("Part 2: %v\n", p2)
}

func part1(pairs [][]int) (safeReports int) {
	for _, report := range pairs {
		if isExcessivelySafe(report) {
			safeReports++
		}
	}
	return
}

func isExcessivelySafe(report []int) bool {
	isDecreasing := true
	if report[0] < report[1] {
		isDecreasing = false
	}
	if report[0] == report[1] {
		return false
	}
	for i := 0; i < len(report)-1; i++ {
		diff := report[i] - report[i+1]
		if diff < 0 && isDecreasing {
			return false
		}
		if diff > 0 && !isDecreasing {
			return false
		}
		absDiff := math.Abs(float64(diff))
		if absDiff > 3 || absDiff < 1 {
			return false
		}
	}
	return true
}

func part2(reports [][]int) (safeCount int) {
	for _, report := range reports {
		if isPermissivelySafe(report) {
			safeCount++
		}
	}
	return
}

func isPermissivelySafe(report []int) bool {
	if isExcessivelySafe(report) {
		return true
	}
	for i := 0; i < len(report); i++ {
		// for some reason if the report slice isn't copied, then it gets modified
		// causing subsequent checks to run on bad data
		// seems to happen in the isExcessivelySafe function, but the function
		// doesn't modify the argument
		cp := make([]int, len(report))
		copy(cp, report)
		partial1 := cp[:i]
		partial2 := cp[i+1:]
		partialReport := append(partial1, partial2...)
		if isExcessivelySafe(partialReport) {
			return true
		}
	}
	return false
}
