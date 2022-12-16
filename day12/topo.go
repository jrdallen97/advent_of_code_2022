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
		return false
	}
	val := t.Get(c)
	// Handle the start case (target == S + 1)
	if limit == 'T' {
		return val == 'a'
	}
	return val <= limit
}

func (t Topo) GetMoves(c Coord, v Visited) map[Direction]Coord {
	limit := t.Get(c) + 1
	moves := map[Direction]Coord{}
	if up := c.AddY(-1); t.isValidMove(up, limit, v) {
		moves[Up] = up
	}
	if down := c.AddY(1); t.isValidMove(down, limit, v) {
		moves[Down] = down
	}
	if left := c.AddX(-1); t.isValidMove(left, limit, v) {
		moves[Left] = left
	}
	if right := c.AddX(1); t.isValidMove(right, limit, v) {
		moves[Right] = right
	}
	return moves
}

func (t Topo) Traverse(start Coord, v Visited) {
	// if E, end?
	// else, traverse possible moves
	// only return if you reach E?

	// I'm thinking to make this recursive somehow?
	moves := t.GetMoves(start, v)
	for _, move := range moves {
		fmt.Println("possible move:", move)
	}
}
