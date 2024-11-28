package day3

import (
	"fmt"
	"testing"
)

var examples = []struct {
	In   string
	Out1 int
	Out2 int
}{
	{">", 2, 2},
	{"^v", 2, 3},
	{"^>v<", 4, 3},
	{"^v^v^v^v^v", 2, 11},
}

func Test_part1(t *testing.T) {
	for _, example := range examples {
		expected := example.Out1
		actual := part1(example.In)
		if expected != actual {
			fmt.Printf("Expected: %v, Actual: %v", expected, actual)
			t.Fail()
		}
	}
}

func Test_part2(t *testing.T) {
	for i, example := range examples {
		expected := example.Out2
		actual := part2(example.In)
		if expected != actual {
			fmt.Printf("%v Expected: %v, Actual: %v\n", i, expected, actual)
			t.Fail()
		}
	}
}
