package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

type Turn struct {
	Them, You Move
}

type Move int

const (
	UnknownMove Move = 0
	Rock             = 1
	Paper            = 2
	Scissors         = 3
)

type Outcome int

const (
	UnknownOutcome Outcome = 0
	Lose                   = 1
	Draw                   = 2
	Win                    = 3
)

func main() {
	{
		f, err := os.Open("day2/guide.txt")
		if err != nil {
			panic(err)
		}
		guide := ParseRPS(f, false)

		total := 0
		for _, round := range guide {
			total += Score(round)
		}
		fmt.Println("v1 score:", total)
	}

	{
		f, err := os.Open("day2/guide.txt")
		if err != nil {
			panic(err)
		}
		guide := ParseRPS(f, true)
		total := 0
		for _, round := range guide {
			total += Score(round)
		}
		fmt.Println("v2 score:", total)
	}
}

func Score(t *Turn) int {
	score := int(t.You)
	if t.You == t.Them {
		// Draw
		return score + 3
	} else if (t.You == Rock && t.Them == Scissors) || int(t.You)-1 == int(t.Them) {
		// Win
		return score + 6
	}
	// Lose
	return score
}

func ParseRPS(r io.Reader, v2 bool) []*Turn {
	scanner := bufio.NewScanner(r)
	var guide []*Turn
	for scanner.Scan() {
		split := strings.Split(scanner.Text(), " ")
		turn := &Turn{
			Them: AtoMove(split[0]),
		}
		if v2 {
			turn.You = FindMove(turn.Them, AtoOutcome(split[1]))
		} else {
			turn.You = AtoMove(split[1])
		}
		guide = append(guide, turn)
	}
	return guide
}

func FindMove(move Move, outcome Outcome) Move {
	if outcome == Draw {
		return move
	}
	if outcome == Win {
		if move == Scissors {
			return Rock
		}
		return Move(int(move) + 1)
	}
	// Lose
	if move == Rock {
		return Scissors
	}
	return Move(int(move) - 1)
}

func AtoMove(s string) Move {
	switch s {
	case "A", "X":
		return Rock
	case "B", "Y":
		return Paper
	case "C", "Z":
		return Scissors
	default:
		return UnknownMove
	}
}

func AtoOutcome(s string) Outcome {
	switch s {
	case "X":
		return Lose
	case "Y":
		return Draw
	case "Z":
		return Win
	default:
		return UnknownOutcome
	}
}
