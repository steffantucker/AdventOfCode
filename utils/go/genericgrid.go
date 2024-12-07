package utils

import "strings"

type GenericGrid[V comparable] interface {
	At(Coordinate) V
	Set(Coordinate, V)
	Find(V) Coordinate
	FindFunc(func(V) bool) Coordinate
	FindAll(V) []Coordinate
	FindFuncAll(func(V) bool) []Coordinate
}

type GenericMapGrid[V comparable] struct {
	Grid    map[Coordinate]V
	Bounds  Rectangle
	Default V
}

func NewGenericMapGrid[V comparable](def V) *GenericMapGrid[V] {
	return &GenericMapGrid[V]{
		Grid:    make(map[Coordinate]V),
		Bounds:  EmptyRectangle(),
		Default: def,
	}
}

func (m *GenericMapGrid[V]) FillFromString(in, rowSplit, colSplit string, newFunc func(string) V) {
	m.FillFromStringArray(strings.Split(in, rowSplit), colSplit, newFunc)
}

func (m *GenericMapGrid[V]) FillFromStringArray(in []string, colSplit string, newFunc func(string) V) {
	for y, row := range in {
		for x, val := range strings.Split(row, colSplit) {
			m.Set(NewCoordinate(x, y), newFunc(val))
		}
	}
}

func (m *GenericMapGrid[V]) At(c Coordinate) V {
	return m.Grid[c]
}

func (m *GenericMapGrid[V]) Set(c Coordinate, v V) {
	if c.X < m.Bounds.TopLeft.X {
		m.Bounds.TopLeft.X = c.X
	}
	if c.Y < m.Bounds.TopLeft.Y {
		m.Bounds.TopLeft.Y = c.Y
	}
	if c.X > m.Bounds.BottomRight.X {
		m.Bounds.BottomRight.X = c.X
	}
	if c.Y > m.Bounds.BottomRight.Y {
		m.Bounds.BottomRight.Y = c.Y
	}
	m.Grid[c] = v
}

func (m *GenericMapGrid[V]) Remove(c Coordinate) {
	delete(m.Grid, c)
}

func (m *GenericMapGrid[V]) Find(what V) (Coordinate, bool) {
	return m.FindFunc(func(v V) bool { return v == what })
}

func (m *GenericMapGrid[V]) FindAll(what V) []Coordinate {
	return m.FindFuncAll(func(v V) bool { return v == what })
}

func (m *GenericMapGrid[V]) FindFunc(how func(V) bool) (Coordinate, bool) {
	coords := m.FindFuncAll(how)
	if len(coords) == 0 {
		return Coordinate{X: 0, Y: 0}, false
	}
	return coords[0], true
}

func (m *GenericMapGrid[V]) FindFuncAll(how func(V) bool) []Coordinate {
	coords := []Coordinate{}
	for c, val := range m.Grid {
		if how(val) {
			coords = append(coords, c)
		}
	}
	return coords
}

func (m *GenericMapGrid[V]) IsInBounds(c Coordinate) bool {
	return m.Bounds.IsInRectangle(c)
}
