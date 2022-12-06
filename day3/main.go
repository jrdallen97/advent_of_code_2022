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
		f, err := os.Open("day3/backpacks.txt")
		//f, err := os.Open("day3/simple.txt")
		if err != nil {
			panic(err)
		}

		scanner := bufio.NewScanner(f)
		sum := 0
		e1, e2 := items{}, items{}
		for elf := 0; scanner.Scan(); elf++ {
			runes := []rune(scanner.Text())
			for _, r := range runes {
				switch elf % 3 {
				case 0:
					e1[r] = struct{}{}
				case 1:
					e2[r] = struct{}{}
				case 2:
					_, ok1 := e1[r]
					_, ok2 := e2[r]
					if ok1 && ok2 {
						sum += runeToPriority(r)
						e1, e2 = items{}, items{}
						break
					}
				}
			}
		}
		fmt.Println("sum:", sum)
	}
}

func runeToPriority(r rune) int {
	if r <= 'Z' {
		return int(r) - 64 + 26
	}
	return int(r) - 96
}
