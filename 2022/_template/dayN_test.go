package dayN

import (
	"fmt"
	"testing"
)

var pairs = []string{}

func Test_part1(t *testing.T) {
	expected := 2
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
