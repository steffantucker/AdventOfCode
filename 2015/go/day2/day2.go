package day2

import (
	"fmt"
	"slices"
	"strings"

	utils "github.com/steffantucker/AdventOfCode/utils/go"
)

func Run() {
	dimensions := utils.GetStringList(2015, 2)
	p1 := part1(dimensions)
	fmt.Printf("Part 1: %v\n", p1)
	p2 := part2(dimensions)
	fmt.Printf("Part 2: %v\n", p2)
}

func part1(dims []string) (total int) {
	for _, dim := range dims {
		d := strings.Split(dim, "x")
		l, w, h := utils.MustAtoi(d[0]), utils.MustAtoi(d[1]), utils.MustAtoi(strings.TrimSpace(d[2]))
		lw, wh, hl := l*w, w*h, h*l
		min := utils.SliceMin([]int{lw, wh, hl})
		total += 2*(l*w+w*h+h*l) + min
	}
	return
}

func part2(dims []string) (total int) {
	for _, dim := range dims {
		d := strings.Split(dim, "x")
		dint := []int{utils.MustAtoi(d[0]), utils.MustAtoi(d[1]), utils.MustAtoi(strings.TrimSpace(d[2]))}
		slices.Sort[[]int](dint)
		cubic := dint[0] * dint[1] * dint[2]
		total += 2*dint[0] + 2*dint[1] + cubic
	}
	return
}
