package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

func main() {
	f, err := os.Open("day1/calories.txt")
	//f, err := os.Open("day1/simple.txt")
	if err != nil {
		panic(err)
	}
	elves := ParseElves(f)

	elf, carrying := TopElf(elves)
	fmt.Println("carrying most:", elf+1)
	fmt.Println("carrying:", carrying)
	fmt.Println()

	top3 := 0
	for i := 1; i <= 3; i++ {
		elf, carrying = TopElf(elves)
		top3 += carrying
		// Remove the top elf from the list
		elves = append(elves[:elf], elves[elf+1:]...)
	}
	fmt.Println("top 3 total:", top3)
}

func ParseElves(r io.Reader) []int {
	scanner := bufio.NewScanner(r)

	elves := []int{0}
	current := 0
	for scanner.Scan() {
		text := scanner.Text()
		if text == "" {
			current++
			elves = append(elves, 0)
		} else if n, err := strconv.Atoi(text); err == nil {
			elves[current] += n
		}
	}
	return elves
}

// TopElf takes all elves and returns the position and value of the elf carrying
// the most.
func TopElf(elves []int) (int, int) {
	max := 0
	for i, e := range elves {
		if e > elves[max] {
			max = i
		}
	}
	return max, elves[max]
}
