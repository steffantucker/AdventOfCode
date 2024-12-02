package day2

import (
	"fmt"
	"testing"
)

var input = [][]int{
	{7, 6, 4, 2, 1}, // safe
	{1, 2, 7, 8, 9}, // unsafe
	{9, 7, 6, 2, 1}, // unsafe
	{1, 3, 2, 4, 5}, // safe
	{8, 6, 4, 4, 1}, // safe
	{1, 3, 6, 7, 9}, // safe
	{1, 1, 2},       // safe
	{2, 6, 1},       // safe
	{1, 2, 5, 9},    // safe
}

func Test_part1(t *testing.T) {
	expected := 2
	max := part1(input)

	if max != expected {
		fmt.Printf("Expected %v got %v\n", expected, max)
		t.FailNow()
	}
}

func Test_part2(t *testing.T) {
	expected := 7
	max := part2(input)

	if max != expected {
		fmt.Printf("Expected %v got %v\n", expected, max)
		t.FailNow()
	}
}
