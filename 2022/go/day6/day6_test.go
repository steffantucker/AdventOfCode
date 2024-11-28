package day6

import (
	"fmt"
	"testing"
)

var datastreams = []string{
	"mjqjpqmgbljsphdztnvjfqwrcgsmlb",
	"bvwbjplbgvbhsrlpgdmjqwftvncz",
	"nppdvjthqldpwncqszvftbrmjlhg",
	"nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg",
	"zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw",
}

var expects = []int{
	7,
	5,
	6,
	10,
	11,
}

var messageExpects = []int{
	19,
	23,
	23,
	29,
	26,
}

func Test_part1(t *testing.T) {
	for i, datastream := range datastreams {
		marker := part(datastream, 4)

		if marker != expects[i] {
			fmt.Printf("Expected %v got %v\n", expects[i], marker)
			t.FailNow()
		}
	}
}

func Test_part2(t *testing.T) {
	for i, datastream := range datastreams {
		marker := part(datastream, 14)

		if marker != messageExpects[i] {
			fmt.Printf("Expected %v got %v\n", messageExpects[i], marker)
			t.FailNow()
		}
	}
}
