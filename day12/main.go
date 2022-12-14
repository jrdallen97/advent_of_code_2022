package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/davecgh/go-spew/spew"
)

//type Direction int
//
//const (
//	Up Direction = iota
//	Down
//	Left
//	Right
//)

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
	f, err := os.Open("day12/map.txt")
	//f, err := os.Open("day12/simple.txt")
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

	path := []Coord{start}
	bestPath := topo.Traverse(start, path)

	fmt.Println(bestPath)
	fmt.Println("steps:", len(bestPath)-1)
	fmt.Println("------")
	//fmt.Println(len(bestPath))
	//fmt.Println(bestPath)

	// Starting from S, build a tree of every possible move (without backtracks)
	// and then select the tree with the shortest path to E?
}
