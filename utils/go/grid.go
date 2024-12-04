package utils

import (
	"strings"
)

type Grid interface {
	Set(Coordinate, rune)
	At(Coordinate) rune
	Bounds() (int, int, int, int)
	String() string
}

type ArrayGrid [][]rune

func (a *ArrayGrid) Set(c Coordinate, s rune) {
	if c.Y > len(*a) {
		*a = append(*a, make([][]rune, c.Y-len(*a)+1)...)
	}
	if c.X > len((*a)[c.Y]) {
		(*a)[c.Y] = append((*a)[c.Y], make([]rune, c.X-len((*a)[c.Y])+1)...)
	}
	(*a)[c.Y][c.X] = s
}

func (a ArrayGrid) At(c Coordinate) rune {
	return a[c.Y][c.X]
}

func (a ArrayGrid) Bounds() (minx, miny, maxx, maxy int) {
	for _, y := range a {
		if len(y)-1 > maxx {
			maxx = len(y) - 1
		}
	}
	return 0, 0, maxx, len(a)
}

func (a ArrayGrid) String() string {
	s := strings.Builder{}
	for _, row := range a {
		for _, v := range row {
			if v == 0 {
				s.WriteString(" ")
				continue
			}
			s.WriteRune(v)
		}
		s.WriteString("\n")
	}
	return s.String()
}
