package main

import "fmt"

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

func (c Coord) Left() Coord {
  return Coord{c.X() - 1, c.Y()}
}

func (c Coord) Right() Coord {
  return Coord{c.X() + 1, c.Y()}
}

func (c Coord) Up() Coord {
  return Coord{c.X(), c.Y() - 1}
}

func (c Coord) Down() Coord {
  return Coord{c.X(), c.Y() + 1}
}

func (c Coord) PossibleNeighbours() []Coord {
	return []Coord{c.Up(), c.Down(), c.Left(), c.Right()}
}

