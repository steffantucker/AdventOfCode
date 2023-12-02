package day14

import (
	"fmt"
	"testing"
)

var pairs = []string{
	"498,4 -> 498,6 -> 496,6",
	"503,4 -> 502,4 -> 502,9 -> 494,9",
}

func Test_part1(t *testing.T) {
	expected := 24
	max := part1(pairs)

	if max != expected {
		fmt.Printf("Expected %v got %v\n", expected, max)
		t.FailNow()
	}
}

func Test_part2(t *testing.T) {
	expected := 4
	max := part2(pairs)

	if max != expected {
		fmt.Printf("Expected %v got %v\n", expected, max)
		t.FailNow()
	}
}
