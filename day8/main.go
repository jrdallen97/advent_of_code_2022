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
}
