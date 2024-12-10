package day4

import (
	"fmt"

	utils "github.com/steffantucker/AdventOfCode/utils/go"
)

func Run() {
	input := utils.GetString(2024, 4)
	p1 := part1(input)
	fmt.Printf("Part 1: %v\n", p1)
	p2 := part2(input)
	fmt.Printf("Part 2: %v\n", p2)
}

func part1(input string) (result int) {
	search := newp1(input)
	return search.FindXMASES()
}

type p1 struct {
	Grid utils.MapGrid
	Xs   []utils.Coordinate
}

func newp1(input string) p1 {
	p := p1{
		Grid: utils.NewMapGrid(),
	}
	p.Grid.FillFromString(input)
	p.Xs = p.Grid.FindAll("X")
	return p
}

func (p p1) FindXMASES() (count int) {
	for _, xcoord := range p.Xs {
		count += p.HasXMAS(xcoord)
	}
	return
}

func (p p1) IsInDirection(what string, c utils.Coordinate, dir utils.Direction) bool {
	return p.Grid.At(utils.SumDirection(c, dir)) == what
}

func (p p1) HasXMAS(x utils.Coordinate) (count int) {
	ms := p.Grid.FindAround(x, "M", false)
	for dir, mcoord := range ms {
		if p.IsInDirection("A", mcoord, dir) {
			acoord := utils.SumDirection(mcoord, dir)
			if p.IsInDirection("S", acoord, dir) {
				count++
			}
		}
	}
	return
}

func part2(input string) (result int) {
	p := newp2(input)
	return p.CountXMASES()
}

type p2 struct {
	Grid utils.MapGrid
	As   []utils.Coordinate
}

func newp2(input string) p2 {
	p := p2{Grid: utils.NewMapGrid()}
	p.Grid.FillFromString(input)
	p.As = p.Grid.FindAll("A")
	return p
}

func (p p2) CountXMASES() (count int) {
	for _, a := range p.As {
		mdirs := p.Grid.FindAround(a, "M", false)
		delete(mdirs, utils.UP)
		delete(mdirs, utils.DOWN)
		delete(mdirs, utils.LEFT)
		delete(mdirs, utils.RIGHT)
		if len(mdirs) != 2 {
			continue
		}
		mascount := 0
		for mdir := range mdirs {
			if p.IsInDirection("S", a, utils.ReverseDirections[mdir]) {
				mascount++
			}
		}
		if mascount == 2 {
			count++
		}
	}
	return
}

func (p p2) IsInDirection(what string, c utils.Coordinate, dir utils.Direction) bool {
	return p.Grid.At(utils.SumDirection(c, dir)) == what
}
