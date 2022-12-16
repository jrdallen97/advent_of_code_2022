package main

import "fmt"

// Track the maximum valid x/y values
var (
	MaxX = 0
	MaxY = 0
)

// Coord {x, y}, e.g.:
// {0, 0} {1, 0}
// {0, 1} {1, 1}
type Coord [2]int

func (c Coord) String() string {
	return fmt.Sprintf("{%d, %d}", c.X(), c.Y())
}

func (c Coord) X() int {
	return c[0]
}

func (c Coord) Y() int {
	return c[1]
}

func (c Coord) AddX(n int) Coord {
	return Coord{c.X() + n, c.Y()}
}

func (c Coord) AddY(n int) Coord {
	return Coord{c.X(), c.Y() + n}
}

func (c Coord) IsValidCoord() bool {
	if c.X() < 0 || c.X() > MaxX || c.Y() < 0 || c.Y() > MaxY {
		return false
	}
	return true
}
