package day13

import (
	"fmt"
	"testing"
)

var pairs = []string{
	`[1,1,3,1,1]
[1,1,5,1,1]`,
	`[[1],[2,3,4]]
[[1],4]`,
	`[9]
[[8,7,6]]`,
	`[[4,4],4,4]
[[4,4],4,4,4]`,
	`[7,7,7,7]
[7,7,7]`,
	`[]
[3]`,
	`[[[]]]
[[]]`,
	`[1,[2,[3,[4,[5,6,7]]]],8,9]
[1,[2,[3,[4,[5,6,0]]]],8,9]`,
}

func Test_part1(t *testing.T) {
	expected := 13
	max := part1(pairs)

	if max != expected {
		fmt.Printf("Expected %v got %v\n", expected, max)
		t.FailNow()
	}
}

func Test_part2(t *testing.T) {
	expected := 140
	max := part2(pairs)

	if max != expected {
		fmt.Printf("Expected %v got %v\n", expected, max)
		t.FailNow()
	}
}
