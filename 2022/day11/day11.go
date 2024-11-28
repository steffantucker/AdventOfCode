package day11

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/steffantucker/AdventOfCode/2022/utils"
)

func Run() {
	business := MakeMonkeyBusiness(utils.GetParagraphs(2022, 11))
	p1 := part1(business)
	fmt.Printf("Part 1: %v\n", p1)
	business = MakeMonkeyBusiness(utils.GetParagraphs(2022, 11))
	p2 := part2(business)
	fmt.Printf("Part 2: %v\n", p2)
}

func MakeMonkeyBusiness(list []string) MonkeyBusiness {
	business := MonkeyBusiness{}
	business.product = 1
	business.monkeys = make([]Monkey, 0, len(list))
	for _, m := range list {
		business.monkeys = append(business.monkeys, NewMonkey(m))
	}
	for _, m := range business.monkeys {
		business.product *= m.testNum
	}
	return business
}

func NewMonkey(m string) Monkey {
	monkey := Monkey{}
	var id, t, f, opNum, test int
	var list, opSymbol string
	fmt.Sscanf(strings.NewReplacer("* old", "^ 2", ", ", ",").Replace(m), `Monkey %d:
  Starting items: %s
  Operation: new = old %s %d
  Test: divisible by %d
    If true: throw to monkey %d
    If false: throw to monkey %d")`, &id, &list, &opSymbol, &opNum, &test, &t, &f)

	monkey.id = id
	monkey.trueThrow = t
	monkey.falseThrow = f
	monkey.testNum = test
	json.Unmarshal([]byte("["+list+"]"), &monkey.items)
	monkey.operation = map[string]func(int) int{
		"+": func(o int) int { return o + opNum },
		"-": func(o int) int { return o - opNum },
		"*": func(o int) int { return o * opNum },
		"/": func(o int) int { return o / opNum },
		"^": func(o int) int { return o * o },
	}[opSymbol]
	monkey.test = func(worry int) int {
		if worry%test == 0 {
			return t
		}
		return f
	}
	return monkey
}

type MonkeyBusiness struct {
	monkeys []Monkey
	product int
}

func (m *MonkeyBusiness) run(count int, worryFunc func(int) int) {
	for i := 0; i < count; i++ {
		for j, monkey := range m.monkeys {
			for _, item := range monkey.items {
				worry := worryFunc(monkey.operation(item))
				m.monkeys[monkey.test(worry)].items = append(m.monkeys[monkey.test(worry)].items, worry)
				m.monkeys[j].inspectionCount++
			}
			m.monkeys[j].items = monkey.items[:0]
		}
	}
}

func (m MonkeyBusiness) level() int {
	first, second := 0, 0
	for _, monkey := range m.monkeys {
		if monkey.inspectionCount > first {
			first, second = monkey.inspectionCount, first
		} else if monkey.inspectionCount > second {
			second = monkey.inspectionCount
		}
	}
	return first * second
}

type Monkey struct {
	id              int
	items           []int
	trueThrow       int
	falseThrow      int
	inspectionCount int
	testNum         int

	operation func(int) int
	test      func(int) int
}

func part1(bussiness MonkeyBusiness) int {
	bussiness.run(20, func(i int) int { return i / 3 })
	return bussiness.level()
}

func part2(business MonkeyBusiness) int {
	business.run(10000, func(i int) int { return i % business.product })
	return business.level()
}
