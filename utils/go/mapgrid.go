package utils

import (
	"strings"
)

type MapGrid struct {
	Grid    map[Coordinate]string
	Bounds  Rectangle
	Default string
}

type Rectangle struct {
	TopLeft, BottomRight Coordinate
}

func EmptyRectangle() Rectangle {
	return Rectangle{
		TopLeft:     NewCoordinate(0, 0),
		BottomRight: NewCoordinate(0, 0),
	}
}

func (r Rectangle) IsInRectangle(c Coordinate) bool {
	x := c.X >= r.TopLeft.X && c.X <= r.BottomRight.X
	y := c.Y >= r.TopLeft.Y && c.Y <= r.BottomRight.Y
	return x && y
}

func NewMapGrid() MapGrid {
	return NewMapGridWithDefault(" ")
}

func NewMapGridWithDefault(def string) MapGrid {
	return MapGrid{
		Grid: make(map[Coordinate]string),
		Bounds: Rectangle{
			TopLeft: Coordinate{X: 0, Y: 0},
		},
		Default: def,
	}
}

func (m *MapGrid) FillFromString(grid string) {
	for y, row := range strings.Split(grid, "\n") {
		for x, val := range strings.Split(row, "") {
			m.Set(c(x, y), val)
		}
	}
}

func (m MapGrid) find(what string, all bool) []Coordinate {
	coordinates := []Coordinate{}
	for y := m.Bounds.TopLeft.Y; y <= m.Bounds.BottomRight.Y; y++ {
		for x := m.Bounds.TopLeft.X; x <= m.Bounds.BottomRight.X; x++ {
			c := NewCoordinate(x, y)
			if val, ok := m.Grid[c]; ok && val == what {
				coordinates = append(coordinates, c)
				if !all {
					return coordinates
				}
			}
		}
	}
	return coordinates
}

func (m MapGrid) Find(what string) Coordinate {
	return m.find(what, false)[0]
}

func (m MapGrid) FindAll(what string) []Coordinate {
	return m.find(what, true)
}

func (m MapGrid) FindAround(c Coordinate, what string, ortho bool) map[Direction]Coordinate {
	directionMap := OctilinearCoords
	if ortho {
		directionMap = OrthogonalCoords
	}
	directions := make(map[Direction]Coordinate)
	for direction, coord := range directionMap {
		look := Sum(c, coord)
		if val, ok := m.Grid[look]; ok && val == what {
			directions[direction] = look
		}
	}
	return directions
}

func (m *MapGrid) Set(c Coordinate, r string) {
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
	m.Grid[c] = r
}

func (m MapGrid) At(c Coordinate) string {
	if r, ok := m.Grid[c]; !ok {
		return m.Default
	} else {
		return r
	}
}

func (m MapGrid) GetBounds() (bounds Rectangle) {
	return m.Bounds
}

func (m MapGrid) IsInBounds(c Coordinate) bool {
	x := c.X >= m.Bounds.TopLeft.X && c.X <= m.Bounds.BottomRight.X
	y := c.Y >= m.Bounds.TopLeft.Y && c.Y <= m.Bounds.BottomRight.Y

	return x && y
}

func (m MapGrid) String() string {
	s := strings.Builder{}
	bounds := m.GetBounds()
	for y := bounds.TopLeft.Y; y <= bounds.BottomRight.Y; y++ {
		for x := bounds.TopLeft.X; x <= bounds.BottomRight.Y; x++ {
			if r, ok := m.Grid[c(x, y)]; !ok {
				s.WriteString(m.Default)
			} else {
				s.WriteString(r)
			}
		}
		s.WriteString("\n")
	}
	return s.String()
}
