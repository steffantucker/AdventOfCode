package day2

import (
	"fmt"
	"testing"
)

var guide = []string{
	"A Y",
	"B X",
	"C Z",
}

func Test_part1(t *testing.T) {
	expected := 15
	max := part1(guide)

	if max != expected {
		fmt.Printf("Expected %v got %v\n", expected, max)
		t.FailNow()
	}
}

func Test_part2(t *testing.T) {
	expected := 12
	max := part2(guide)

	if max != expected {
		fmt.Printf("Expected %v got %v\n", expected, max)
		t.FailNow()
	}
}
