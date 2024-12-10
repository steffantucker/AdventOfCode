package day9

import (
	"fmt"
	"testing"
)

// p1:1928
// p2:2858
var input = "2333133121414131402"

var p2 = []struct {
	Input    string
	Expected int
}{
	{
		Input:    "2333133121414131402",
		Expected: 2858,
	},
	{
		Input:    "12345",
		Expected: 132,
	},
	{
		Input:    "14113",
		Expected: 16,
	},
	{
		Input:    "10101010101010101010101",
		Expected: 506,
	},
}

func Test_part1(t *testing.T) {
	expected := 1928
	result := part1(input)

	if result != expected {
		fmt.Printf("Expected %v got %v\n", expected, result)
		t.FailNow()
	}
}

func Test_part2(t *testing.T) {
	for _, test := range p2 {
		result := part2(test.Input)
		if result != test.Expected {
			fmt.Printf("Expected %v got %v\n", test.Expected, result)
			t.FailNow()
		}
	}
}
