package day7

import (
	"fmt"
	"strings"
	"testing"
)

var input = `190: 10 19
3267: 81 40 27
83: 17 5
156: 15 6
7290: 6 8 6 15
161011: 16 10 13
192: 17 8 14
21037: 9 7 18 13
292: 11 6 16 20`

func Test_part1(t *testing.T) {
	expected := 3749
	result := solve(parse(strings.Split(input, "\n")), []string{"*", "+"})

	if result != expected {
		fmt.Printf("Expected %v got %v\n", expected, result)
		t.FailNow()
	}
}

func Test_part2(t *testing.T) {
	expected := 11387
	result := solve(parse(strings.Split(input, "\n")), []string{"|", "*", "+"})

	if result != expected {
		fmt.Printf("Expected %v got %v\n", expected, result)
		t.FailNow()
	}
}
