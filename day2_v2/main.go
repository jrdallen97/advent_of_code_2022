package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

const (
	Rock int = iota
	Paper
	Scissors
)

const (
	Lose int = iota
	Draw
	Win
)

type Turn struct {
	Them, You int
	Desired   int
}

func (t *Turn) outcome() int {
	switch t.Them - t.You {
	case -1, 2:
		return Win
	case 0:
		return Draw
	default:
		return Lose
	}
}

func (t *Turn) Score() int {
	return t.You + 1 + t.outcome()*3
}

func (t *Turn) CalcMove() {
	t.You = normalise(t.Them + (t.Desired - 1))
}

func normalise(n int) int {
	// Add 3 to protect against negative numbers
	return (n + 3) % 3
}

func main() {
	f, err := os.Open("day2/guide.txt")
	//f, err := os.Open("day2/simple.txt")
	if err != nil {
		panic(err)
	}
	guide := ParseGuide(f)

	// Part 1
	{
		total := 0
		for _, round := range guide {
			total += round.Score()
		}
		fmt.Println("v1 score:", total)
	}

	// Part 2
	{
		total := 0
		for _, round := range guide {
			round.CalcMove()
			total += round.Score()
		}
		fmt.Println("v2 score:", total)
	}
}

func ParseGuide(r io.Reader) []*Turn {
	scanner := bufio.NewScanner(r)
	var guide []*Turn
	for scanner.Scan() {
		split := strings.Split(scanner.Text(), " ")
		guide = append(guide, &Turn{
			Them:    AtoMove(split[0]),
			You:     AtoMove(split[1]),
			Desired: AtoOutcome(split[1]),
		})
	}
	return guide
}

func AtoMove(s string) int {
	switch s {
	case "A", "X":
		return Rock
	case "B", "Y":
		return Paper
	default: // "C", "Z"
		return Scissors
	}
}

func AtoOutcome(s string) int {
	switch s {
	case "X":
		return Lose
	case "Y":
		return Draw
	default: // "Z"
		return Win
	}
}
