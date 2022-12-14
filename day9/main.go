package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Direction string

const (
	Up    Direction = "U"
	Down            = "D"
	Left            = "L"
	Right           = "R"
)

// Position tracks a position in 2d space.
// Bottom-left is 0,0 and moving right/up will increase x/y.
type Position struct {
	x, y int
}

func (p *Position) String() string {
	return fmt.Sprintf("%d,%d", p.x, p.y)
}

func (p *Position) Move(d Direction) {
	switch d {
	case Up:
		p.y++
	case Down:
		p.y--
	case Left:
		p.x--
	case Right:
		p.x++
	}
}

func (p *Position) Follow(head *Position) {
	if abs(head.x-p.x) > 1 || abs(head.y-p.y) > 1 {
		p.y = p.y + move(head.y-p.y)
		p.x = p.x + move(head.x-p.x)
	}
}

func move(i int) int {
	if i == 0 {
		return 0
	}
	if i < 0 {
		return -1
	}
	return 1
}

func abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}

func main() {
	f, err := os.Open("day9/movements.txt")
	//f, err := os.Open("day9/simple.txt")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(f)
	knots := make([]*Position, 10)
	for i := 0; i < len(knots); i++ {
		knots[i] = &Position{}
	}

	part1 := map[string]struct{}{}
	part2 := map[string]struct{}{}
	for scanner.Scan() {
		d, n := parse(scanner.Text())
		for i := 0; i < n; i++ {
			knots[0].Move(d)
			for x := 1; x < len(knots); x++ {
				knots[x].Follow(knots[x-1])
			}
			part1[knots[1].String()] = struct{}{}
			part2[knots[9].String()] = struct{}{}
			// TODO: would be cool to print the rope positions as we go
		}
	}
	fmt.Println("part 1 tail positions:", len(part1))
	fmt.Println("part 2 tail positions:", len(part2))
}

func parse(line string) (Direction, int) {
	parts := strings.Split(line, " ")
	d := Direction(parts[0])
	n, err := strconv.Atoi(parts[1])
	if err != nil {
		panic(err)
	}
	return d, n
}
