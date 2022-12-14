package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Direction string

const (
	Top    Direction = "top"
	Bottom           = "bottom"
	Left             = "left"
	Right            = "right"
)

func main() {
	f, err := os.Open("day8/trees.txt")
	//f, err := os.Open("day8/simple.txt")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(f)
	var lines [][]int
	row := 0
	for scanner.Scan() {
		text := scanner.Text()
		lines = append(lines, make([]int, len(text)))
		for i, r := range []rune(text) {
			val, err := strconv.Atoi(string(r))
			if err != nil {
				panic(err)
			}
			lines[row][i] = val
		}
		row++
	}

	v := NewVisibility(lines)
	v.Print(Top)
	v.Print(Bottom)
	v.Print(Left)
	v.Print(Right)

	fmt.Println("visible trees:", v.CountVisible())
	scores := map[string]int{}
	for y := range lines {
		for x := range lines[y] {
			coord := fmt.Sprintf("%d,%d", x, y)
			scores[coord] = calcScore(lines, x, y)
		}
	}

	max := 0
	for coord, val := range scores {
		if val > max {
			max = val
			fmt.Println("new max:", coord, val)
		}
	}
	fmt.Println("max:", max)
}

// not my best work lol
func calcScore(lines [][]int, x, y int) int {
	self := lines[y][x]
	up, left, right, down := 0, 0, 0, 0

	for i := x + 1; i < len(lines[y]); i++ {
		right++
		if lines[y][i] >= self {
			break
		}
	}

	for i := y + 1; i < len(lines); i++ {
		down++
		if lines[i][x] >= self {
			break
		}
	}

	for i := x - 1; i >= 0; i-- {
		left++
		if lines[y][i] >= self {
			break
		}
	}

	for i := y - 1; i >= 0; i-- {
		up++
		if lines[i][x] >= self {
			break
		}
	}

	return up * left * right * down
}
