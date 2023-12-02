package utils

import (
	"fmt"
	"math"
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

type MapGrid map[Coordinate]rune

func (m *MapGrid) Set(c Coordinate, r rune) {
	(*m)[c] = r
}

func (m MapGrid) At(c Coordinate) rune {
	if r, ok := m[c]; !ok {
		return rune(0)
	} else {
		return r
	}
}

func (m MapGrid) Bounds() (minx, miny, maxx, maxy int) {
	maxx, maxy = math.MinInt, math.MinInt
	minx, miny = math.MaxInt, math.MaxInt
	for k := range m {
		if k.X > maxx {
			maxx = k.X
		}
		if k.Y > maxy {
			maxy = k.Y
		}
		if k.X < minx {
			minx = k.X
		}
		if k.Y < miny {
			miny = k.Y
		}
	}
	return
}

func (m MapGrid) String() string {
	s := strings.Builder{}
	minx, miny, maxx, maxy := m.Bounds()
	for y := miny; y <= maxy; y++ {
		for x := minx; x <= maxx; x++ {
			if r, ok := m[c(x, y)]; !ok {
				s.WriteString(" ")
			} else {
				s.WriteRune(r)
			}
		}
		s.WriteString("\n")
	}
	return s.String()
}

type Coordinate struct {
	X, Y int
}

func c(x, y int) Coordinate {
	return Coordinate{X: x, Y: y}
}

func CoordFromString(s string) (c Coordinate) {
	cs := strings.Split(strings.TrimSpace(s), ",")
	if len(cs) > 2 {
		panic(fmt.Sprintf("Coordinates formatted wrong, provided: %q", s))
	}
	c.X = MustAtoi(strings.TrimSpace(cs[0]))
	c.Y = MustAtoi(strings.TrimSpace(cs[1]))
	return
}

// func North East South West
