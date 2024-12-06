package day5

import (
	"fmt"
	"slices"
	"strings"

	utils "github.com/steffantucker/AdventOfCode/utils/go"
)

func Run() {
	input := utils.GetParagraphs(2024, 5)
	p1 := part1(input)
	fmt.Printf("Part 1: %v\n", p1)
	p2 := part2(input)
	fmt.Printf("Part 2: %v\n", p2)
}

func part1(input []string) (result int) {
	rules := parseRules(input[0])
	pages := strings.Split(input[1], "\n")
	for _, page := range pages {
		if middle, ok := rules.IsCompliant(page); ok {
			result += middle
		}
	}

	return
}

type Rules struct {
	BeforeRules map[string][]string
	AfterRules  map[string][]string
}

func (r Rules) IsCompliant(pageList string) (int, bool) {
	pages := strings.Split(pageList, ",")
	middle := utils.MustAtoi(pages[len(pages)/2])
	for i := 0; i < len(pages)-1; i++ {
		before := !verify(r.BeforeRules[pages[i]], pages[i+1:])
		after := !verify(r.AfterRules[pages[i]], pages[:i])
		if before || after {
			return 0, false
		}
	}

	return middle, true
}

func verify(pageRules []string, list []string) bool {
	return verifyWithIndex(pageRules, list) == -1
}

func verifyWithIndex(pageRules []string, list []string) int {
	if len(pageRules) == 0 || len(list) == 0 {
		return -1
	}
	for _, rule := range pageRules {
		if i := slices.Index(list, rule); i != -1 {
			return i
		}
	}
	return -1
}

func (r Rules) Fix(in string) int {
	pages := strings.Split(in, ",")
	for i := 0; i < len(pages)-1; i++ {
		currentPage := pages[i]
		before := pages[:i]
		if beforeProblem := verifyWithIndex(r.AfterRules[pages[i]], before); beforeProblem != -1 {
			problem := before[beforeProblem]
			before = slices.Delete(before, beforeProblem, beforeProblem+1)
			pages = slices.Concat(before, []string{currentPage, problem}, pages[i+1:])
			i -= 2
			continue
		}
		after := pages[i+1:]
		if afterProblem := verifyWithIndex(r.BeforeRules[pages[i]], after); afterProblem != -1 {
			problem := after[afterProblem]
			after = slices.Delete(after, afterProblem, afterProblem+1)
			pages = slices.Concat(pages[:i], []string{problem, currentPage}, after)
			i -= 1
		}
	}
	middle := utils.MustAtoi(pages[len(pages)/2])
	if _, ok := r.IsCompliant(strings.Join(pages, ",")); !ok {
		panic(fmt.Sprintf("fixed: %v but not compliant", middle))
	}
	return middle
}

func parseRules(rulesString string) Rules {
	afterRules := make(map[string][]string)
	beforeRules := make(map[string][]string)
	for _, rule := range strings.Split(rulesString, "\n") {
		parts := strings.Split(rule, "|")
		afterRules[parts[0]] = append(afterRules[parts[0]], parts[1])
		beforeRules[parts[1]] = append(beforeRules[parts[1]], parts[0])
	}
	return Rules{
		AfterRules:  afterRules,
		BeforeRules: beforeRules,
	}
}

// bad:
// 11162
func part2(input []string) (result int) {
	rules := parseRules(input[0])
	pages := strings.Split(input[1], "\n")
	for _, page := range pages {
		if _, ok := rules.IsCompliant(page); !ok {
			result += rules.Fix(page)
		}
	}

	return
}
