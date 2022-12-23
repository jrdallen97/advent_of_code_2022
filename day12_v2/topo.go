package main

import (
	"fmt"

	"github.com/jrdallen97/advent_of_code_2022/utils"
)

func NewTopo(vals map[Coord]int) *Topo {
	nodes := make(map[Coord]*Node, len(vals))
	maxX, maxY := 0, 0
	for c, v := range vals {
		nodes[c] = &Node{Val: v}
		if c.X() > maxX {
			maxX = c.X()
		}
		if c.Y() > maxY {
			maxY = c.Y()
		}
	}

	t := &Topo{
		Nodes: nodes,
		maxX: maxX,
		maxY: maxY,
	}
	t.calcNeighbours()
	return t
}

type Topo struct {
	Nodes map[Coord]*Node
	maxX, maxY int
}

func (t *Topo) calcNeighbours() {
	for c, n := range t.Nodes {
		for _, possible := range c.PossibleNeighbours() {
			if neighbour, ok := t.Nodes[possible]; ok && neighbour.Val <= n.Val + 1 {
				n.Neighbours = append(n.Neighbours, possible)
			}
		}
	}
}

// Colours for printing
const (
	grey string = "\033[37m"
	white       = "\033[31m"
	reset       = "\033[0m"
)

func (t *Topo) Print(path []Coord) {
	coords := utils.NewSet(path...)
	for y := 0; y <= t.maxY; y++ {
		for x := 0; x <= t.maxX; x++ {
			c := Coord{x, y}
			if coords.Contains(c) {
				fmt.Printf(white)
			} else {
				fmt.Printf(grey)
			}
			fmt.Printf("%2d%s ", t.Nodes[Coord{x, y}].Val, reset)
		}
		fmt.Println()
	}
}

func (t *Topo) Find(val int) Coord {
	for c, n := range t.Nodes {
		if n.Val == val {
			return c
		}
	}
	return Coord{}
}

func (t *Topo) FindPath(from Coord, target int) []Coord {
	distance := 1
	t.Nodes[from].Path = []Coord{from}
	found := false
	var end Coord
	for !found {
		for _, n := range t.Nodes {
			if len(n.Path) != distance {
				continue
			}

			for _, c2 := range n.Neighbours {
				n2 := t.Nodes[c2]
				if len(n2.Path) == 0 {
					for _, pos := range n.Path {
						n2.Path = append(n2.Path, pos)
					}
					n2.Path = append(n2.Path, c2)
				}
				if n2.Val == target {
					found = true
					end = c2
					break
				}
			}
		}
		distance++
	}
	return t.Nodes[end].Path
}
