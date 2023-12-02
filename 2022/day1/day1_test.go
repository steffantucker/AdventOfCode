package day1

import (
	"fmt"
	"testing"
)

var snacks = [][]int{
	{
		1000,
		2000,
		3000,
	}, {
		4000,
	}, {
		5000,
		6000,
	}, {
		7000,
		8000,
		9000,
	}, {
		10000,
	},
}

func Test_part1(t *testing.T) {
	expected := 24000
	max := part1(snacks)

	if max != expected {
		fmt.Printf("Expected %v got %v\n", expected, max)
		t.FailNow()
	}
}

func Test_part2(t *testing.T) {
	expected := 45000
	max := part2(snacks)

	if max != expected {
		fmt.Printf("Expected %v got %v\n", expected, max)
		t.FailNow()
	}
}
