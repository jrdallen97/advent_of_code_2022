package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	f, err := os.Open("day6/signal.txt")
	//f, err := os.Open("day6/simple.txt")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(f)

	// First, parse the initial state
	for scanner.Scan() {
		text := scanner.Text()
		fmt.Println("pkt start:", findStart(text, 4)+1)
		fmt.Println("msg start:", findStart(text, 14)+1)
	}
}

// Finds the first index at which the previous n characters were unique
func findStart(s string, n int) int {
	match := ""
	for current, r := range []rune(s) {
		i := strings.IndexRune(match, r)
		if i != -1 {
			// If the rune already exists, cut away the invalid parts of the match
			match = match[i+1:] + string(r)
		} else {
			match += string(r)
		}
		if len(match) == n {
			return current
		}
	}
	return -1
}
