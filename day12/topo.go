package main

import (
	"fmt"
	"strings"

	"golang.org/x/text/language"
	"golang.org/x/text/message"
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

func (t Topo) isValidMove(c Coord, current rune, v Visited) bool {
	// Outside the topo
	if !c.IsValidCoord() {
		return false
	}
	// Backtracking
	if v.Contains(c) {
		//fmt.Println("preventing backtrack")
		return false
	}
	next := t.Get(c)
	// Handle the start & end cases differently
	if current == 'S' {
		return next == 'a'
	}
	if next == 'E' {
		return current == 'z'
	}
	return next <= current+1
}

var movesTried int
var p = message.NewPrinter(language.English)

func (t Topo) GetMoves(c Coord, v Visited) []Coord {
	movesTried++
	if movesTried%100_000 == 0 {
		p.Printf("moves tested: %d\n", movesTried)
	}

	// If we're at the end, stop
	current := t.Get(c)
	//if current == 'E' {
	//	return nil
	//}

	var moves []Coord
	if right := c.AddX(1); t.isValidMove(right, current, v) {
		moves = append(moves, right)
	}
	if up := c.AddY(-1); t.isValidMove(up, current, v) {
		moves = append(moves, up)
	}
	if down := c.AddY(1); t.isValidMove(down, current, v) {
		moves = append(moves, down)
	}
	if left := c.AddX(-1); t.isValidMove(left, current, v) {
		moves = append(moves, left)
	}
	return moves
}

// Initialise best to something quite high to prevent pointless spinning
var best = 100

func (t Topo) Traverse(start Coord, path []Coord) []Coord {
	fmt.Println("depth", len(path))
	// If it's already longer than the best path, give up
	if len(path) > best {
		return nil
	}

	// If this path has reached the end, return it
	if t.Get(start) == 'E' {
		if len(path) < best {
			best = len(path)
		}
		fmt.Println("found a path:", len(path))
		return path
	}

	// Otherwise, get all possible moves from this coordinate
	v := NewVisited(path...)
	moves := t.GetMoves(start, v)
	// If there are no moves left, this path is dead and should just return nil
	if len(moves) == 0 {
		return nil
	}

	// Otherwise, play out all possible moves and return the best winning outcome
	var bestPath []Coord
	for _, move := range moves {
		// Update: manually copying the array seems to fix it... Wtf?
		nextPath := make([]Coord, len(path)+1)
		copy(nextPath, append(path, move))

		if possiblePath := t.Traverse(move, nextPath); possiblePath != nil {
			if len(bestPath) == 0 || len(possiblePath) < len(bestPath) {
				bestPath = possiblePath
			}
		}
	}
	return bestPath
}
