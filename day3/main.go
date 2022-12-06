package main

import (
	"bufio"
	"fmt"
	"os"
)

type items map[rune]struct{}

func main() {
	// Part 1
	{
		f, err := os.Open("day3/backpacks.txt")
		//f, err := os.Open("day3/simple.txt")
		if err != nil {
			panic(err)
		}

		scanner := bufio.NewScanner(f)
		sum := 0
		for scanner.Scan() {
			runes := []rune(scanner.Text())
			c1 := items{}
			for i, r := range runes {
				if i < len(runes)/2 {
					c1[r] = struct{}{}
				} else if _, ok := c1[r]; ok {
					sum += runeToPriority(r)
					break
				}
			}
		}
		fmt.Println("sum:", sum)
	}

	// Part 2
	{
	}
}

func runeToPriority(r rune) int {
	if r <= 'Z' {
		return int(r) - 64 + 26
	}
	return int(r) - 96
}
