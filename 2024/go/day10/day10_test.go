package day10

import (
	"fmt"
	"testing"
)

var input = []struct {
	Input      string
	P1Expected int
	P2Expected int
}{
	{
		Input: `0123
1234
8765
9876`,
		P1Expected: 1,
	},
	{
		Input: `...0...
...1...
...2...
6543456
7.....7
8.....8
9.....9`,
		P1Expected: 2,
		P2Expected: 2,
	},
	{
		Input: `..90..9
...1.98
...2..7
6543456
765.987
876....
987....`,
		P1Expected: 4,
		P2Expected: 13,
	},
	{
		Input: `10..9..
2...8..
3...7..
4567654
...8..3
...9..2
.....01`,
		P1Expected: 3,
	},
	{
		Input: `89010123
78121874
87430965
96549874
45678903
32019012
01329801
10456732`,
		P1Expected: 36,
		P2Expected: 81,
	},
	{
		Input: `012345
123456
234567
345678
4.6789
56789.`,
		P2Expected: 227,
	},
	{
		Input: `.....0.
..4321.
..5..2.
..6543.
..7..4.
..8765.
..9....`,
		P2Expected: 3,
	},
}

func Test_part1(t *testing.T) {
	for _, test := range input {
		if test.P1Expected == 0 {
			continue
		}
		result := part1(test.Input)

		if result != test.P1Expected {
			fmt.Printf("Expected %v got %v\n", test.P1Expected, result)
			t.FailNow()
		}
	}
}

func Test_part2(t *testing.T) {
	for _, test := range input {
		if test.P2Expected == 0 {
			continue
		}
		result := part2(test.Input)

		if result != test.P2Expected {
			fmt.Printf("Expected %v got %v\n", test.P2Expected, result)
			t.FailNow()
		}
	}
}
