package day1

import (
	"fmt"
	"testing"
)

var lines = [][]int{
	{3, 4},
	{4, 3},
	{2, 5},
	{1, 3},
	{3, 9},
	{3, 3},
}

func Test_part1(t *testing.T) {
	expected := 11
	sum := part1(lines)

	if sum != expected {
		fmt.Printf("Expected %v got %v\n", expected, sum)
		t.FailNow()
	}
}

func Test_part2(t *testing.T) {
	expected := 31
	sum := part2(lines)

	if sum != expected {
		fmt.Printf("Expected %v got %v\n", expected, sum)
		t.FailNow()
	}
}
