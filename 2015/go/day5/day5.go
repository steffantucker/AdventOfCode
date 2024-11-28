package day5

import (
	"fmt"
	"regexp"
	"strings"

	utils "github.com/steffantucker/AdventOfCode/utils/go"
)

func Run() {
	list := utils.GetStringList(2015, 5)
	p1 := part1(list)
	fmt.Printf("Part 1: %v\n", p1)
	p2 := part2(list)
	fmt.Printf("Part 2: %v\n", p2)
}

func part1(list []string) (nicecount int) {
	for _, non := range list {
		if isNice(non) {
			nicecount++
		}
	}
	return
}

func isNice(s string) bool {
	if strings.Contains(s, "ab") || strings.Contains(s, "cd") || strings.Contains(s, "pq") || strings.Contains(s, "xy") {
		return false
	}
	re := regexp.MustCompile(`[aeiou]`)
	matches := re.FindAll([]byte(s), -1)
	if len(matches) >= 3 {
		var last rune
		for _, c := range s {
			if c == last {
				return true
			}
			last = c
		}
	}
	return false
}

func part2(list []string) (nicecount int) {
	for _, l := range list {
		if isNiceBetter(l) {
			nicecount++
		}
	}
	return
}

func isNiceBetter(s string) bool {
	if !hasPairs(s) {
		return false
	}
	if !hasSeperated(s) {
		return false
	}
	return true
}

func hasSeperated(s string) bool {
	runes := []rune(s)
	for i := 0; i < len(runes)-2; i += 1 {
		if runes[i] == runes[i+2] {
			return true
		}
	}
	return false
}

func hasPairs(s string) bool {
	runes := []rune(s)
	for i := 0; i < len(runes)-3; i += 1 {
		pair := string(runes[i]) + string(runes[i+1])
		for j := i + 2; j < len(runes)-1; j += 1 {
			pair2 := string(runes[j]) + string(runes[j+1])
			if pair == pair2 {
				return true
			}
		}
	}
	return false
}
