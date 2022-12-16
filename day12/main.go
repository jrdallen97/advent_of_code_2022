package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/davecgh/go-spew/spew"
)

type Direction int

const (
	Up Direction = iota
	Down
	Left
	Right
)

func NewVisited(coords ...Coord) Visited {
	v := Visited{}
	for _, coord := range coords {
		v[coord] = struct{}{}
	}
	return v
}

type Visited map[Coord]struct{}

func (v Visited) Values() []Coord {
	coords := make([]Coord, len(v))
	var i int
	for coord := range v {
		coords[i] = coord
		i++
	}
	return coords
}

func (v Visited) Contains(c Coord) bool {
	_, ok := v[c]
	return ok
}

func main() {
	//f, err := os.Open("day12/map.txt")
	f, err := os.Open("day12/simple.txt")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(f)
	var topo Topo
	for scanner.Scan() {
		topo = append(topo, scanner.Text())
	}
	MaxX = len(topo[0]) - 1
	MaxY = len(topo) - 1
	spew.Dump(topo)

	start, end := topo.find('S'), topo.find('E')
	fmt.Println("start:", start)
	fmt.Println("end:", end)

	//path := []Coord{start}
	path := []Coord{
		{0, 0},
		{1, 0},
		{2, 0},
		{2, 1},
		{1, 1},
		{0, 1},
		{0, 2},
		{0, 3},
		{0, 4},
		{1, 4},
		{1, 3},
		{2, 3},
		{2, 4},
		{3, 4},
		{4, 4},
		{5, 4},
		{6, 4},
		{7, 4},
		{7, 3},
		{7, 2},
		{7, 1},
		{7, 0},
		{6, 0},
		{5, 0},
		{4, 0},
	}
	// Should I just use the Set helpers for this?
	possiblePaths := topo.Traverse(Coord{4, 0}, path)
	fmt.Println("possible paths:", len(possiblePaths))

	shortest := -1
	best := 0
	for i, path := range possiblePaths {
		//fmt.Println(path)
		if path[len(path)-1] == end && (len(path) < shortest || shortest == -1) {
			shortest = len(path)
			best = i
		}
	}
	fmt.Println(possiblePaths[best])
	fmt.Println("------")
	//fmt.Println(len(bestPath))
	//fmt.Println(bestPath)

	// Starting from S, build a tree of every possible move (without backtracks)
	// and then select the tree with the shortest path to E?
}
