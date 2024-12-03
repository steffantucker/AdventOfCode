package dayN

import (
	"fmt"
	"testing"
)

var input = []string{}

func Test_part1(t *testing.T) {
	expected := 2
	result := part1(input)

	if result != expected {
		fmt.Printf("Expected %v got %v\n", expected, result)
		t.FailNow()
	}
}

func Test_part2(t *testing.T) {
	expected := 4
	result := part2(input)

	if result != expected {
		fmt.Printf("Expected %v got %v\n", expected, result)
		t.FailNow()
	}
}
