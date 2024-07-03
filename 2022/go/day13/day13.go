package day13

import (
	"encoding/json"
	"fmt"
	"reflect"
	"slices"
	"strings"

	utils "github.com/steffantucker/AdventOfCode/utils/go"
)

func Run() {
	pairs := utils.GetParagraphs(2022, 13)
	p1 := part1(pairs)
	fmt.Printf("Part 1: %v\n", p1)
	p2 := part2(pairs)
	fmt.Printf("Part 2: %v\n", p2)
}

func compare(a, b any) (res int) {
	tleft, tright := reflect.TypeOf(a).String(), reflect.TypeOf(b).String()
	if tleft == tright && tleft == "float64" {
		if a.(float64) < b.(float64) {
			return -1
		} else if a.(float64) > b.(float64) {
			return 1
		}
		return 0
	}
	left := a.([]any)
	right := b.([]any)
	for i := 0; i < len(left); i++ {
		// run out of right
		if i >= len(right) {
			return 1
		}
		tleft, tright = reflect.TypeOf(left[i]).String(), reflect.TypeOf(right[i]).String()
		// first element of both is number
		if tleft == tright {
			// first element both array
			res = compare(left[i], right[i])
			if res != 0 {
				return
			}
		} else {
			// left is a number
			if tleft == "float64" {
				res = compare([]any{left[i]}, right[i])
				if res != 0 {
					return
				}
			} else {
				// right is a number
				res = compare(left[i], []any{right[i]})
				if res != 0 {
					return
				}
			}
		}
	}
	if len(left) < len(right) {
		return -1
	}
	return 0
}

func part1(pairs []string) (sum int) {
	var left, right []any

	for i, pair := range pairs {
		p := strings.Split(pair, "\n")
		json.Unmarshal([]byte(p[0]), &left)
		json.Unmarshal([]byte(p[1]), &right)
		if res := compare(left, right); res > 0 {
			sum += i + 1
		}
	}
	return
}

func part2(pairs []string) (key int) {
	div1 := []any{[]any{2.0}}
	div2 := []any{[]any{6.0}}
	packets := []any{div1, div2}
	for _, pair := range pairs {
		var left, right []any
		p := strings.Split(pair, "\n")
		json.Unmarshal([]byte(p[0]), &left)
		json.Unmarshal([]byte(p[1]), &right)
		packets = append(packets, left)
		packets = append(packets, right)
	}

	slices.SortFunc(packets, compare)

	key1, key2 := 0, 0
	for i, packet := range packets {
		if reflect.DeepEqual(packet, div1) {
			key1 = i + 1
		}
		if reflect.DeepEqual(packet, div2) {
			key2 = i + 1
		}
	}
	return key1 * key2
}
