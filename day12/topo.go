package main

import (
	"fmt"
	"strings"
)

type Topo []string

func (t Topo) find(r rune) Coord {
	for y, s := range t {
		x := strings.IndexRune(s, r)
		if x != -1 {
			return Coord{x, y}
		}
	}
	return Coord{}
}

func (t Topo) Get(c Coord) rune {
	return rune(t[c[1]][c[0]])
}

func (t Topo) isValidMove(c Coord, limit rune, v Visited) bool {
	// Outside the topo
	if !c.IsValidCoord() {
		return false
	}
	// Backtracking
	if v.Contains(c) {
		//fmt.Println("preventing backtrack")
		return false
	}
	val := t.Get(c)
	// Handle the start case (target == S + 1)
	if limit == 'T' {
		return val == 'a'
	}
	return val <= limit
}

func (t Topo) GetMoves(c Coord, v Visited) []Coord {
	// If we're at the end, stop
	current := t.Get(c)
	if current == 'E' {
		return nil
	}

	limit := current + 1
	var moves []Coord
	if up := c.AddY(-1); t.isValidMove(up, limit, v) {
		moves = append(moves, up)
	}
	if down := c.AddY(1); t.isValidMove(down, limit, v) {
		moves = append(moves, down)
	}
	if left := c.AddX(-1); t.isValidMove(left, limit, v) {
		moves = append(moves, left)
	}
	if right := c.AddX(1); t.isValidMove(right, limit, v) {
		moves = append(moves, right)
	}
	return moves
}

func (t Topo) Traverse(start Coord, currentPath []Coord) [][]Coord {
	v := NewVisited(currentPath...)

	moves := t.GetMoves(start, v)
	// If there are no moves left, this path is dead and should just return itself
	if len(moves) == 0 {
		return [][]Coord{currentPath}
	}

	// Otherwise, play out all possible moves and return all their possible outcomes
	var paths [][]Coord
	for _, move := range moves {
		fmt.Println(currentPath)
		// Next path seems to always be valid...
		nextPath := append(currentPath, move)
		fmt.Println(nextPath)
		// But paths seems to contain garbage
		paths = append(paths, t.Traverse(move, nextPath)...)
	}
	return paths
}
