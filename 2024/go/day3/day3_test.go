package day3

import (
	"fmt"
	"testing"
)

type TestData struct {
	input string
	part1 int
	part2 int
}

var testCases = []TestData{
	{
		"mul(44,46)",
		2024,
		2024,
	},
	{
		"mul(4*",
		0,
		0,
	},
	{
		"mul(6,9!",
		0,
		0,
	},
	{
		"mul ( 2 , 4 )",
		0,
		0,
	},
	{
		"xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))",
		161,
		161,
	},
	{
		"xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))",
		161,
		48,
	},
}

func Test_part1(t *testing.T) {
	for _, testCase := range testCases {
		result := part1(testCase.input)

		if result != testCase.part1 {
			fmt.Printf("Expected %v got %v\n", testCase.part1, result)
			t.FailNow()
		}
	}
}

func Test_part2(t *testing.T) {
	for _, testCase := range testCases {
		result := part2(testCase.input)

		if result != testCase.part2 {
			fmt.Printf("Expected %v got %v\n", testCase.part2, result)
			t.FailNow()
		}
	}
}
