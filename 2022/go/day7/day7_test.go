package day7

import (
	"fmt"
	"testing"
)

var termOutput = []string{
	"$ cd /",
	"$ ls",
	"dir a",
	"14848514 b.txt",
	"8504156 c.dat",
	"dir d",
	"$ cd a",
	"$ ls",
	"dir e",
	"29116 f",
	"2557 g",
	"62596 h.lst",
	"$ cd e",
	"$ ls",
	"584 i",
	"$ cd ..",
	"$ cd ..",
	"$ cd d",
	"$ ls",
	"4060174 j",
	"8033020 d.log",
	"5626152 d.ext",
	"7214296 k",
}

func Test_part1(t *testing.T) {
	root := buildFileStructure(termOutput)
	expected := 95437
	max := part1(root, 100000)

	if max != expected {
		fmt.Printf("Expected %v got %v\n", expected, max)
		t.FailNow()
	}
}

func Test_part2(t *testing.T) {
	root := buildFileStructure(termOutput)
	expected := 24933642
	max := part2(root)

	if max != expected {
		fmt.Printf("Expected %v got %v\n", expected, max)
		t.FailNow()
	}
}
