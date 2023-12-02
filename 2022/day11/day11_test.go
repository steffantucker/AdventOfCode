package day11

import (
	"fmt"
	"testing"
)

var pairs = []string{
	`Monkey 0:
  Starting items: 79, 98
  Operation: new = old * 19
  Test: divisible by 23
    If true: throw to monkey 2
    If false: throw to monkey 3`,
	`Monkey 1:
  Starting items: 54, 65, 75, 74
  Operation: new = old + 6
  Test: divisible by 19
    If true: throw to monkey 2
    If false: throw to monkey 0`,
	`Monkey 2:
  Starting items: 79, 60, 97
  Operation: new = old * old
  Test: divisible by 13
    If true: throw to monkey 1
    If false: throw to monkey 3`,
	`Monkey 3:
  Starting items: 74
  Operation: new = old + 3
  Test: divisible by 17
    If true: throw to monkey 0
    If false: throw to monkey 1`,
}

func Test_part1(t *testing.T) {
	expected := 10605
	bussiness := MakeMonkeyBusiness(pairs)
	level := part1(bussiness)

	if level != expected {
		fmt.Printf("Expected %v got %v\n", expected, level)
		t.FailNow()
	}
}

func Test_part2(t *testing.T) {
	expected := 2713310158
	bussiness := MakeMonkeyBusiness(pairs)
	level := part2(bussiness)

	if level != expected {
		fmt.Printf("Expected %v got %v\n", expected, level)
		t.FailNow()
	}
}
