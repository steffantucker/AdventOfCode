package day6

import (
	"fmt"

	utils "github.com/steffantucker/AdventOfCode/utils/go"
)

func Run() {
	input := utils.GetStringList(2024, 6)
	p1 := part1(input)
	fmt.Printf("Part 1: %v\n", p1)
	p2 := part2(input)
	fmt.Printf("Part 2: %v\n", p2)
}

type Guard struct {
	Direction utils.Direction
	Location  utils.Coordinate
}

var clockwise map[utils.Direction]utils.Direction = map[utils.Direction]utils.Direction{
	utils.UP:    utils.RIGHT,
	utils.RIGHT: utils.DOWN,
	utils.DOWN:  utils.LEFT,
	utils.LEFT:  utils.UP,
}

func part1(input []string) (result int) {
	grid := utils.NewGenericMapGrid(".")
	grid.FillFromStringArray(input, "", func(v string) string { return v })

	guard := Guard{
		Direction: utils.UP,
	}
	if g, ok := grid.Find("^"); !ok {
		panic("Guard not found")
	} else {
		guard.Location = g
		grid.Remove(g)
	}

	gaurdOOB := false
	for !gaurdOOB {
		nextSymbol := "."
		nextLocation := utils.Coordinate{X: 0, Y: 0}
		switch guard.Direction {
		case utils.UP:
			nextLocation = utils.Sum(guard.Location, utils.Coordinate{X: 0, Y: -1})
			nextSymbol = grid.At(nextLocation)
		case utils.RIGHT:
			nextLocation = utils.Sum(guard.Location, utils.Coordinate{X: 1, Y: 0})
			nextSymbol = grid.At(nextLocation)
		case utils.DOWN:
			nextLocation = utils.Sum(guard.Location, utils.Coordinate{X: 0, Y: 1})
			nextSymbol = grid.At(nextLocation)
		case utils.LEFT:
			nextLocation = utils.Sum(guard.Location, utils.Coordinate{X: -1, Y: 0})
			nextSymbol = grid.At(nextLocation)
		}
		if nextSymbol == "#" {
			guard.Direction = clockwise[guard.Direction]
		} else {
			grid.Set(guard.Location, "X")
			guard.Location = nextLocation
		}
		if !grid.IsInBounds(nextLocation) {
			break
		}
	}

	return len(grid.FindAll("X"))
}

// 3107 too high
func part2(input []string) (result int) {
	// place objects along path guard travels
	grid := utils.NewGenericMapGrid(".")
	grid.FillFromStringArray(input, "", func(s string) string { return s })
	obstacles := grid.FindAll("#")

	g, _ := grid.Find("^")
	guard := Guard{
		Location:  g,
		Direction: utils.UP,
	}
	gaurdOOB := false
	for !gaurdOOB {
		nextSymbol := "."
		nextLocation := utils.Coordinate{X: 0, Y: 0}
		switch guard.Direction {
		case utils.UP:
			nextLocation = utils.Sum(guard.Location, utils.Coordinate{X: 0, Y: -1})
			nextSymbol = grid.At(nextLocation)
		case utils.RIGHT:
			nextLocation = utils.Sum(guard.Location, utils.Coordinate{X: 1, Y: 0})
			nextSymbol = grid.At(nextLocation)
		case utils.DOWN:
			nextLocation = utils.Sum(guard.Location, utils.Coordinate{X: 0, Y: 1})
			nextSymbol = grid.At(nextLocation)
		case utils.LEFT:
			nextLocation = utils.Sum(guard.Location, utils.Coordinate{X: -1, Y: 0})
			nextSymbol = grid.At(nextLocation)
		}
		if nextSymbol == "#" {
			guard.Direction = clockwise[guard.Direction]
		} else {
			grid.Set(guard.Location, "X")
			guard.Location = nextLocation
		}
		if !grid.IsInBounds(nextLocation) {

			break
		}
	}
	grid.Remove(g)

	xs := grid.FindAll("X")

	// c := make(chan bool)
	// defer close(c)
	// for _, x := range xs {
	// 	go isLoopingAsync(Guard{Direction: utils.UP, Location: utils.NewCoordinate(g.X, g.Y)}, append(obstacles, x), c)
	// }
	// for range xs {
	// 	if <-c {
	// 		result++
	// 	}
	// }
	for _, x := range xs {
		isLooping(Guard{Direction: utils.UP, Location: utils.NewCoordinate(g.X, g.Y)}, append(obstacles, x))
	}

	return
}

func isLooping(guard Guard, obstacles []utils.Coordinate) bool {
	grid := utils.NewGenericMapGrid(utils.NONE)
	for _, obstacle := range obstacles {
		grid.Set(obstacle, utils.CENTER)
	}

	gaurdOOB := false
	for !gaurdOOB {
		var nextSymbol utils.Direction
		nextLocation := utils.Coordinate{X: 0, Y: 0}
		switch guard.Direction {
		case utils.UP:
			nextLocation = utils.Sum(guard.Location, utils.Coordinate{X: 0, Y: -1})
			nextSymbol = grid.At(nextLocation)
		case utils.RIGHT:
			nextLocation = utils.Sum(guard.Location, utils.Coordinate{X: 1, Y: 0})
			nextSymbol = grid.At(nextLocation)
		case utils.DOWN:
			nextLocation = utils.Sum(guard.Location, utils.Coordinate{X: 0, Y: 1})
			nextSymbol = grid.At(nextLocation)
		case utils.LEFT:
			nextLocation = utils.Sum(guard.Location, utils.Coordinate{X: -1, Y: 0})
			nextSymbol = grid.At(nextLocation)
		}
		if nextSymbol == utils.CENTER {
			guard.Direction = clockwise[guard.Direction]
		} else if nextSymbol == guard.Direction {
			printGrid(*grid)
			return true
		} else {
			grid.Set(guard.Location, guard.Direction)
			guard.Location = nextLocation
		}
		if !grid.IsInBounds(nextLocation) {
			break
		}
	}
	return false
}

func isLoopingAsync(guard Guard, obstacles []utils.Coordinate, c chan bool) {
	c <- isLooping(guard, obstacles)
}

func printGrid(grid utils.GenericMapGrid[utils.Direction]) {
	for y := grid.Bounds.TopLeft.Y; y <= grid.Bounds.BottomRight.Y; y++ {
		for x := grid.Bounds.TopLeft.X; x <= grid.Bounds.BottomRight.X; x++ {
			space := grid.At(utils.Coordinate{X: x, Y: y})
			var symbol string
			switch space {
			case utils.UP:
				symbol = "^"
			case utils.RIGHT:
				symbol = ">"
			case utils.DOWN:
				symbol = "v"
			case utils.LEFT:
				symbol = "<"
			case utils.NONE:
				symbol = "."
			case utils.CENTER:
				symbol = "#"
			}
			fmt.Print(symbol)
		}
		fmt.Println()
	}
}
