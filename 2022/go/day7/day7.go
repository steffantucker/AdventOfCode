package day7

import (
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"

	"github.com/steffantucker/AdventOfCode/2022/go/utils"
)

type ComFile struct {
	name        string
	isDirectory bool
	contents    map[string]*ComFile
	size        int
	parent      *ComFile
}

func Run() {
	termOutput := utils.GetStringList(2022, 7)
	root := buildFileStructure(termOutput)
	p1 := part1(root, 100000)
	fmt.Printf("Part 1: %v\n", p1)
	p2 := part2(root)
	fmt.Printf("Part 2: %v\n", p2)
}

func part1(root *ComFile, filterSize int) (sizeSum int) {
	if !root.isDirectory {
		return 0
	}
	for _, file := range root.contents {
		if file.isDirectory {
			sizeSum += part1(file, 100000)
		}
	}
	if root.size <= filterSize {
		return sizeSum + root.size
	}
	return
}

func buildFileStructure(termOutput []string) *ComFile {
	root := ComFile{name: "/", isDirectory: true}
	currentDir := &root
	commandReg := regexp.MustCompile(`\$ (cd|ls) ?(\S*)`)
	for i, line := range termOutput {
		cmd := commandReg.FindStringSubmatch(line)
		if cmd == nil {
			continue
		}
		switch cmd[1] {
		case "cd":
			switch cmd[2] {
			case "..":
				currentDir = currentDir.parent
			case "/":
				currentDir = &root
			default:
				var ok bool
				currentDir, ok = currentDir.contents[cmd[2]]
				if !ok {
					log.Fatalf("Directory not found: %q\n", cmd[2])
				}
			}
		case "ls":
			currentDir.addContents(termOutput[i+1:])
		}
	}
	root.calculateSize()
	return &root
}

func (c *ComFile) calculateSize() int {
	if c.isDirectory {
		for _, content := range c.contents {
			c.size += content.calculateSize()
		}
	}
	return c.size
}

func (c *ComFile) addContents(lines []string) {
	dirReg := regexp.MustCompile(`dir (\S+)`)
	fileReg := regexp.MustCompile(`(\d+) (\S+)`)
	c.contents = make(map[string]*ComFile)

	for _, line := range lines {
		if strings.Contains(line, "$") {
			break
		}
		if name := dirReg.FindStringSubmatch(line); name != nil {
			c.contents[name[1]] = &ComFile{name: name[1], isDirectory: true, parent: c}
		}
		if file := fileReg.FindStringSubmatch(line); file != nil {
			s, _ := strconv.Atoi(file[1])
			c.contents[file[2]] = &ComFile{name: file[2], parent: c, size: s}
		}
	}
}

func part2(root *ComFile) (folderSize int) {
	maxSize := 70000000
	updateSize := 30000000
	neededSize := updateSize - (maxSize - root.size)
	sizeList := sizeFilterList(root, neededSize)
	return utils.SliceMin(sizeList)
}

func sizeFilterList(root *ComFile, filter int) (filterList []int) {
	if root.isDirectory && root.size >= filter {
		filterList = append(filterList, root.size)
	}
	for _, f := range root.contents {
		if f.isDirectory {
			filterList = append(filterList, sizeFilterList(f, filter)...)
		}
	}
	return filterList
}
