package day12

import (
	"fmt"
	"strings"

	"github.com/steffantucker/AdventOfCode/2022/go/utils"
)

func Run() {
	land := utils.GetStringList(2022, 12)
	p1 := part1(land)
	fmt.Printf("Part 1: %v\n", p1)
	p2 := part2(land)
	fmt.Printf("Part 2: %v\n", p2)
}

type Mapwalker struct {
	m                          [][]rune
	rowQueue, columnQueue      []int
	sr, sc                     int
	moves                      int
	layerNodes, nextLayerNodes int
	rows, columns              int
	visited                    [][]bool
	ended                      bool
	start, end                 string
}

func NewMapWalker(m []string, start string) Mapwalker {
	mw := Mapwalker{}
	mw.rowQueue, mw.columnQueue = make([]int, 0, 10), make([]int, 0, 10)
	mw.layerNodes = 1
	mw.rows, mw.columns = len(m), len(m[0])
	mw.start = start
	for r, row := range m {
		if strings.Contains(row, start) {
			mw.sr = r
			mw.sc = strings.Index(row, start)
		}
		mw.m = append(mw.m, []rune(row))
	}
	for i := 0; i < mw.rows; i++ {
		mw.visited = append(mw.visited, make([]bool, mw.columns))
	}
	return mw
}

func (m *Mapwalker) solve(end string) {
	m.end = end
	m.rowQueue = append(m.rowQueue, m.sr)
	m.columnQueue = append(m.columnQueue, m.sc)
	m.visited[m.sr][m.sc] = true
	for len(m.rowQueue) > 0 {
		r := m.rowQueue[len(m.rowQueue)-1]
		c := m.columnQueue[len(m.columnQueue)-1]
		m.rowQueue = m.rowQueue[:len(m.rowQueue)-1]
		m.columnQueue = m.columnQueue[:len(m.columnQueue)-1]
		if string(m.m[r][c]) == end {
			m.ended = true
			break
		}
		m.neighbours(r, c)
		m.layerNodes--
		if m.layerNodes == 0 {
			//m.print()
			m.layerNodes = m.nextLayerNodes
			m.nextLayerNodes = 0
			m.moves++
		}
	}
}

func (m Mapwalker) print() {
	for r, row := range m.m {
		for c, i := range row {
			if m.visited[r][c] {
				fmt.Print("O")
			} else {
				fmt.Print(string(i))
			}
		}
		fmt.Println()
	}
	fmt.Println()
}

func (m *Mapwalker) neighbours(r, c int) {
	dr := []int{-1, 1, 0, 0}
	dc := []int{0, 0, -1, 1}
	compare := m.m[r][c] - 1
	if string(m.m[r][c]) == m.start {
		switch m.start {
		case "S":
			compare = 97
		case "E":
			compare = 122
		}
	}
	if string(m.m[r][c]) == m.end {
		switch m.start {
		case "S":
			compare = 97
		case "E":
			compare = 122
		}
	}
	for i := 0; i < 4; i++ {
		nr := r + dr[i]
		nc := c + dc[i]
		if nr < 0 || nr >= m.rows || nc < 0 || nc >= m.columns {
			continue
		}
		if m.visited[nr][nc] {
			continue
		}
		compareTo := m.m[nr][nc]
		if string(compareTo) == m.end {
			switch m.end {
			case "S":
				compareTo = 97
			case "E":
				compareTo = 122
			}
		}
		if compareTo >= compare {
			m.rowQueue = append([]int{nr}, m.rowQueue...)
			m.columnQueue = append([]int{nc}, m.columnQueue...)
			m.visited[nr][nc] = true
			m.nextLayerNodes++
		}
	}
}

func (m Mapwalker) distance() int {
	if m.ended {
		return m.moves
	}
	return -1
}

func part1(land []string) (moves int) {
	m := NewMapWalker(land, "E")
	m.solve("S")
	return m.distance()
}

func part2(land []string) (moves int) {
	m := NewMapWalker(land, "E")
	m.solve("a")
	return m.moves
}
