package day1

import (
	"fmt"
	"testing"
)

var lines = []string{
	"1abc2",
	"pqr3stu8vwx",
	"a1b2c3d4e5f",
	"treb7uchet",
}

var lines2 = []string{
	"two1nine",
	"eightwothree",
	"abcone2threexyz",
	"xtwone3four",
	"4nineeightseven2",
	"zoneight234",
	"7pqrstsixteen",
}

func Test_part1(t *testing.T) {
	expected := 142
	sum := part1(lines)

	if sum != expected {
		fmt.Printf("Expected %v got %v\n", expected, sum)
		t.FailNow()
	}
}

func Test_part2(t *testing.T) {
	expected := 281
	sum := part2(lines2)

	if sum != expected {
		fmt.Printf("Expected %v got %v\n", expected, sum)
		t.FailNow()
	}
}
