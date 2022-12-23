package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	f, err := os.Open("day12/map.txt")
	//f, err := os.Open("day12/simple.txt")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(f)

	vals := map[Coord]int{}
	y := 0
	for scanner.Scan() {
		for x, c := range scanner.Text() {
			vals[Coord{x, y}] = getVal(c)
		}
		y++
	}

	// Part 1
	{
		t := NewTopo(vals)

		start := t.Find(0)
		fmt.Println("start:", start)

		path := t.FindPath(start, 27)
		t.Print(path)
		fmt.Println("steps:", len(path) - 1)
	}

	// Part 2
	{
		fmt.Println("\nPart 2:")
		// Invert the height map so that I don't have to change my logic
		for c := range vals {
			vals[c] = 27 - vals[c]
		}
		t := NewTopo(vals)
		start := t.Find(0)
		fmt.Println("start:", start)
		// Only looking for 26 (a) this time
		path := t.FindPath(start, 26)
		t.Print(path)
		fmt.Println("steps:", len(path) - 1)
	}
}

func getVal(c rune) int {
	switch c {
	case 'S':
		return 0
	case 'E':
		return 27
	}
	return int(c) - 96
}
