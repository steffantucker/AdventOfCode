package utils

import (
	"fmt"
	"strings"
)

type Direction int

const (
	NONE Direction = iota
	UPLEFT
	UP
	UPRIGHT
	LEFT
	RIGHT
	DOWNLEFT
	DOWN
	DOWNRIGHT
	CENTER
)

var DirectionCoords = map[Direction]Coordinate{
	UPLEFT:    {X: -1, Y: -1},
	UP:        {X: 0, Y: -1},
	UPRIGHT:   {X: 1, Y: -1},
	LEFT:      {X: -1, Y: 0},
	RIGHT:     {X: 1, Y: 0},
	DOWNLEFT:  {X: -1, Y: 1},
	DOWN:      {X: 0, Y: 1},
	DOWNRIGHT: {X: 1, Y: 1},
}

var ReverseDirections = map[Direction]Direction{
	UPLEFT:    DOWNRIGHT,
	UP:        DOWN,
	UPRIGHT:   DOWNLEFT,
	LEFT:      RIGHT,
	RIGHT:     LEFT,
	DOWNLEFT:  UPRIGHT,
	DOWN:      UP,
	DOWNRIGHT: UPLEFT,
}

type Coordinate struct {
	X, Y int
}

func NewCoordinate(x, y int) Coordinate {
	return c(x, y)
}

func Sum(a Coordinate, b Coordinate) Coordinate {
	return c(a.X+b.X, a.Y+b.Y)
}

func SumDirection(a Coordinate, d Direction) Coordinate {
	return Sum(a, DirectionCoords[d])
}

func c(x, y int) Coordinate {
	return Coordinate{X: x, Y: y}
}

func (c Coordinate) Add(x Coordinate) {
	c.X += x.X
	c.Y += x.Y
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
