package day8

import (
	"fmt"

	"github.com/steffantucker/AdventOfCode/2022/utils"
)

type Key struct {
	I, J int
}

func Run() {
	trees := utils.GetNumberMatrix(2022, 8)
	p1 := part1(trees)
	fmt.Printf("Part 1: %v\n", p1)
	p2 := part2(trees)
	fmt.Printf("Part 2: %v\n", p2)
}

func part1(trees [][]int) (visible int) {
	visibleTrees := make(map[Key]bool)
	for i := range trees {
		startTallest, endTallest := trees[i][0], trees[i][len(trees[i])-1]
		height := len(trees) - 1
		for j := range trees[i] {
			width := len(trees[i]) - 1
			if i == 0 || j == 0 {
				visibleTrees[Key{I: i, J: j}] = true
			}
			if height-i == height {
				visibleTrees[Key{I: height, J: j}] = true
			}
			if width-j == width {
				visibleTrees[Key{I: i, J: width}] = true
				continue
			}
			if trees[i][j] > startTallest {
				startTallest = trees[i][j]
				visibleTrees[Key{I: i, J: j}] = true
			}
			if trees[i][width-j] > endTallest {
				endTallest = trees[i][width-j]
				visibleTrees[Key{I: i, J: width - j}] = true
			}
		}
	}
	for column := 1; column < len(trees[0])-1; column++ {
		startTallest, endTallest := trees[0][column], trees[len(trees)-1][column]
		for row := 1; row < len(trees)-1; row++ {
			height := len(trees) - 1
			if trees[row][column] > startTallest {
				startTallest = trees[row][column]
				visibleTrees[Key{I: row, J: column}] = true
			}
			if trees[height-row][column] > endTallest {
				endTallest = trees[height-row][column]
				visibleTrees[Key{I: height - row, J: column}] = true
			}
		}
	}
	return len(visibleTrees)
}

func scenicWest(i, j int, trees [][]int, c chan int) {
	scenic := 0
	treeHeight := trees[i][j]
	for column := j - 1; column >= 0; column-- {
		scenic++
		if trees[i][column] >= treeHeight {
			break
		}
	}
	c <- scenic
}
func scenicSouth(i, j int, trees [][]int, c chan int) {
	scenic := 0
	treeHeight := trees[i][j]
	for row := i + 1; row < len(trees); row++ {
		scenic++
		if trees[row][j] >= treeHeight {
			break
		}
	}
	c <- scenic
}
func scenicNorth(i, j int, trees [][]int, c chan int) {
	scenic := 0
	treeHeight := trees[i][j]
	for row := i - 1; row >= 0; row-- {
		scenic++
		if trees[row][j] >= treeHeight {
			break
		}
	}
	c <- scenic
}
func scenicEast(i, j int, trees [][]int, c chan int) {
	scenic := 0
	treeHeight := trees[i][j]
	for column := j + 1; column < len(trees[i]); column++ {
		scenic++
		if trees[i][column] >= treeHeight {
			break
		}
	}
	c <- scenic
}

func part2(trees [][]int) (mostScenic int) {
	c := make(chan int)
	for i := 1; i < len(trees)-1; i++ {
		for j := 1; j < len(trees[i])-1; j++ {
			scenicScore := 1
			go scenicNorth(i, j, trees, c)
			go scenicEast(i, j, trees, c)
			go scenicSouth(i, j, trees, c)
			go scenicWest(i, j, trees, c)
			for counter := 0; counter < 4; counter++ {
				score := <-c
				scenicScore *= score
			}
			if scenicScore > mostScenic {
				mostScenic = scenicScore
			}
		}
	}
	return
}
