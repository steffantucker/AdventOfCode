package day6

import (
	"fmt"
	"strings"

	"github.com/steffantucker/AdventOfCode/2022/utils"
)

func Run() {
	datastream, _ := utils.GetInput(2022, 6)
	p1 := part(datastream, 4)
	fmt.Printf("Part 1: %v\n", p1)
	p2 := part(datastream, 14)
	fmt.Printf("Part 2: %v\n", p2)
}

func part(datastream string, length int) (marker int) {
	bs := strings.Split(datastream, "")
	for i := length; i < len(bs); i++ {
		if isMarker(bs[i-length : i]) {
			return i
		}
	}
	return -1
}

func isMarker(group []string) bool {
	for i := 0; i < len(group); i++ {
		for j := i + 1; j < len(group); j++ {
			if group[i] == group[j] {
				return false
			}
		}
	}
	return true
}
