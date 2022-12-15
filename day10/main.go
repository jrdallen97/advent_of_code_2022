package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func NewProcess() *Process {
	return &Process{
		x:       1,
		history: []int{1},
	}
}

type Process struct {
	x       int
	history []int
}

func (p *Process) AddX(x int) {
	p.Noop()
	p.Noop()
	p.x += x
}

func (p *Process) Noop() {
	p.draw()
	p.history = append(p.history, p.x)
}

func (p *Process) draw() {
	char := (len(p.history) - 1) % 40
	if abs(char-p.x) <= 1 {
		fmt.Printf("#")
	} else {
		fmt.Printf(".")
	}
	if char == 39 {
		fmt.Println()
	}
}

func (p *Process) SignalStrength() int {
	strength := 0
	for i := 20; i <= 220; i += 40 {
		strength += i * p.history[i]
	}
	return strength
}

func abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}

func main() {
	f, err := os.Open("day10/ops.txt")
	//f, err := os.Open("day10/simple.txt")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(f)
	p := NewProcess()
	for scanner.Scan() {
		parts := strings.Split(scanner.Text(), " ")
		if parts[0] == "addx" {
			x, _ := strconv.Atoi(parts[1])
			p.AddX(x)
		} else {
			p.Noop()
		}
	}

	fmt.Println("strength:", p.SignalStrength())
}
