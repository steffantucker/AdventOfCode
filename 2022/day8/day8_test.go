package day8

import (
	"fmt"
	"testing"
)

var trees = [][]int{
	{3, 0, 3, 7, 3},
	{2, 5, 5, 1, 2},
	{6, 5, 3, 3, 2},
	{3, 3, 5, 4, 9},
	{3, 5, 3, 9, 0},
}

func Test_part1(t *testing.T) {
	expected := 21
	visible := part1(trees)

	if visible != expected {
		fmt.Printf("Expected %v got %v\n", expected, visible)
		t.FailNow()
	}
}

func Test_part2(t *testing.T) {
	expected := 8
	max := part2(trees)

	if max != expected {
		fmt.Printf("Expected %v got %v\n", expected, max)
		t.FailNow()
	}
}
