package day7

import (
	"fmt"
	"strings"

	utils "github.com/steffantucker/AdventOfCode/utils/go"
)

type Calibration struct {
	Want  int
	Haves []int
}

func Run() {
	input := strings.Split(utils.GetString(2024, 7), "\n")
	equations := []Calibration{}
	for _, line := range input {
		if len(line) == 0 {
			break
		}
		parts := strings.Split(line, ": ")
		want := utils.MustAtoi(parts[0])
		haves := []int{}
		for _, h := range strings.Split(parts[1], " ") {
			haves = append(haves, utils.MustAtoi(h))
		}
		equations = append(equations, Calibration{Want: want, Haves: haves})
	}
	p1 := solve(equations, []string{"*", "+"})
	fmt.Printf("Part 1: %v\n", p1)
	p2 := solve(equations, []string{"|", "*", "+"})
	// low   73050804055819
	// wrong 636198093037897
	fmt.Printf("Part 2: %v\n", p2)
}

func parse(input []string) (equations []Calibration) {
	for _, line := range input {
		if len(line) == 0 {
			break
		}
		parts := strings.Split(line, ": ")
		want := utils.MustAtoi(parts[0])
		haves := []int{}
		for _, h := range strings.Split(parts[1], " ") {
			haves = append(haves, utils.MustAtoi(h))
		}
		equations = append(equations, Calibration{Want: want, Haves: haves})
	}
	return
}

func solve(equations []Calibration, symbols []string) (result int) {
	for _, eq := range equations {
		if isSolvable(eq.Want, eq.Haves[1:], eq.Haves[0], symbols) {
			result += eq.Want
		} else {
			fmt.Printf("%d: %d\n", eq.Want, eq.Haves)
		}
	}
	return
}

func isSolvable(want int, haves []int, subtotal int, symbols []string) bool {
	for _, symbol := range symbols {
		itertotal := subtotal
		switch symbol {
		case "*":
			itertotal *= haves[0]
		case "+":
			itertotal += haves[0]
		case "|":
			itertotal = utils.MustAtoi(fmt.Sprintf("%d%d", itertotal, haves[0]))
		}
		if itertotal == want && len(haves) == 1 {
			return true
		}
		if len(haves) > 1 && itertotal < want {
			if found := isSolvable(want, haves[1:], itertotal, symbols); found {
				return true
			}
		}
	}
	return false
}
