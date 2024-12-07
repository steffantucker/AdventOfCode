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

var test = `3072: 803 729 2 8 1
107745048: 3 5 760 8 8 7 3 2 4 38 1 7
59632245: 5 868 95 22 47
13787630077: 740 8 22 5 6 617 58 5
406202478: 929 525 437 52 4`

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
	result := solve(parse(strings.Split(test, "\n")), []string{"|", "*", "+"})

	if result != expected {
		fmt.Printf("Expected %v got %v\n", expected, result)
		t.FailNow()
	}
}
