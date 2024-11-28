package day2

import (
	"fmt"
	"testing"
)

var example = []string{
	"2x3x4",
}
var example2 = []string{
	"1x1x10",
}

func Test_part1(t *testing.T) {
	expected := 58
	actual := part1(example)
	if expected != actual {
		fmt.Printf("Expected: %v, but actual: %v\n", expected, actual)
		t.Fail()
	}

	expected2 := 43
	actual2 := part1(example2)
	if expected2 != actual2 {
		fmt.Printf("Expected: %v, but actual: %v\n", expected2, actual2)
		t.FailNow()
	}
}
func Test_part2(t *testing.T) {
	expected := 34
	actual := part2(example)
	if expected != actual {
		fmt.Printf("Expected: %v, but actual: %v\n", expected, actual)
		t.Fail()
	}

	expected2 := 14
	actual2 := part2(example2)
	if expected2 != actual2 {
		fmt.Printf("Expected: %v, but actual: %v\n", expected2, actual2)
		t.FailNow()
	}
}
