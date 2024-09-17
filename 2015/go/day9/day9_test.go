package day9

import (
	"fmt"
	"testing"
)

var example []string = []string{
	"London to Dublin = 464",
	"London to Belfast = 518",
	"Dublin to Belfast = 141",
}

func Test_part1(t *testing.T) {
	expected := 605
	actual := part1(example)
	if expected != actual {
		fmt.Printf("expected: %v actual: %v", expected, actual)
		t.FailNow()
	}
}
