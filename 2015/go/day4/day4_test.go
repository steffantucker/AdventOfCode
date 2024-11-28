package day4

import (
	"fmt"
	"testing"
)

var examples = []struct {
	In  string
	Out int
}{
	{"abcdef", 609043}, {"pqrstuv", 1048970},
}

func Test_part1(t *testing.T) {
	for i, example := range examples {
		expected := example.Out
		actual := part1(example.In)
		if expected != actual {
			fmt.Printf("%v Expected: %v Actual %v\n", i, expected, actual)
			t.Fail()
		}
	}
}
