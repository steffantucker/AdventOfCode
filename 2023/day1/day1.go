package day1

import (
	"fmt"
	"regexp"

	"github.com/steffantucker/AdventOfCode/2023/utils"
)

func Run() {
	lines := utils.GetStringList(2023, 1)
	p1 := part1(lines)
	fmt.Printf("Part 1: %v\n", p1)
	p2 := part2(lines)
	fmt.Printf("Part 2: %v\n", p2)
}

func part1(lines []string) (calSums int) {
	numRegex := regexp.MustCompile(`(\d{1})`)
	for _, line := range lines {
		nums := numRegex.FindAllString(line, -1)
		sum := utils.MustAtoi(nums[0] + nums[len(nums)-1])
		// fmt.Printf("first: %v\nlast: %v\nsum: %v\n", nums[0], nums[len(nums)-1], sum)
		calSums += sum
	}
	return
}

func part2(lines []string) (calSums int) {
	numDict := map[string]string{
		"1":     "1",
		"one":   "1",
		"2":     "2",
		"two":   "2",
		"3":     "3",
		"three": "3",
		"4":     "4",
		"four":  "4",
		"5":     "5",
		"five":  "5",
		"6":     "6",
		"six":   "6",
		"7":     "7",
		"seven": "7",
		"8":     "8",
		"eight": "8",
		"9":     "9",
		"nine":  "9",
	}
	fRegex := regexp.MustCompile(`(\d{1}|one|two|three|four|five|six|seven|eight|nine)`)
	for _, line := range lines {
		first := numDict[fRegex.FindString(line)]
		last := numDict[getLast(line)]
		if last == "" {
			panic("invalid number")
		}
		sum := utils.MustAtoi(first + last)
		fmt.Printf("str: %v\nnum: %v%v\nsum: %v\ncal: %v\n", line, first, last, sum, calSums)
		calSums += sum
	}
	return
}

func getLast(l string) string {
	regex := regexp.MustCompile(`(\d{1}|one|two|three|four|five|six|seven|eight|nine)`)
	for i := len(l) - 1; i >= 0; i-- {
		if regex.MatchString(l[i:]) {
			return regex.FindString(l[i:])
		}
	}
	return ""
}
