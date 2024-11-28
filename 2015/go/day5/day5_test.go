package day5

import (
	"fmt"
	"testing"
)

var example = []struct {
	In  string
	Out bool
}{
	{"ugknbfddgicrmopn", true},
	{"aaa", true},
	{"jchzalrnumimnmhp", false},
	{"haegwjzuvuyypxyu", false},
	{"dvszwmarrgswjxmb", false},
}

var example2 = []struct {
	In  string
	Out bool
}{
	{"qjhvhtzxzqqjkmpb", true},
	{"xxyxx", true},
	{"uurcxstgmygtbstg", false},
	{"ieodomkazucvgmuy", false},
	{"suerykeptdsutidb", false},
	{"xckozymymezzarpy", true},
}

func Test_isNice(t *testing.T) {
	for i, str := range example {
		expected := str.Out
		actual := isNice(str.In)
		if expected != actual {
			fmt.Printf("%v expected: %v actual:%v\n", i, expected, actual)
			t.Fail()
		}
	}
}

func Test_isNiceBetter(t *testing.T) {
	for i, str := range example2 {
		expected := str.Out
		actual := isNiceBetter(str.In)
		if expected != actual {
			fmt.Printf("%v expected: %v actual:%v\n", i, expected, actual)
			t.Fail()
		}
	}
}
