package day11

import (
	"fmt"
	"strings"
	"testing"
)

var input = []struct {
	Iterations []string
	StoneCount int
}{
	{
		Iterations: []string{
			"125 17",
			"253000 1 7",
			"253 0 2024 14168",
			"512072 1 20 24 28676032",
			"512 72 2024 2 0 2 4 2867 6032",
			"1036288 7 2 20 24 4048 1 4048 8096 28 67 60 32",
			"2097446912 14168 4048 2 0 2 4 40 48 2024 40 48 80 96 2 8 6 7 6 0 3 2",
		},
		StoneCount: 55312,
	},
}

func Test_part1(t *testing.T) {
	rocks := NewRocks(input[0].Iterations[0])

	for i, line := range input[0].Iterations[1:] {
		rocks.Iterate(i + 1)
		split := strings.Split(line, " ")
		if rocks.Count != len(split) {
			fmt.Printf("Expected %v got %v\n", len(split), rocks.Count)
			t.FailNow()
		}
	}
}

func Test_part2(t *testing.T) {
	expected := 4
	result := part2(input[0].Iterations[0])

	if result != expected {
		fmt.Printf("Expected %v got %v\n", expected, result)
		t.FailNow()
	}
}
