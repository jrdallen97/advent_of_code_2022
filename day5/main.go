package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	f, err := os.Open("day5/crates.txt")
	//f, err := os.Open("day5/simple.txt")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(f)

	// First, parse the initial state
	crates := []string{""}
	cratesV2 := []string{""}
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			// Break after the initial state is shown
			break
		}
		if !strings.Contains(line, "[") {
			// Skip the number line
			continue
		}
		crates = parseCrateLine(line, crates)
		cratesV2 = parseCrateLine(line, cratesV2)
	}
	fmt.Println("initial state:", crates)

	// Then parse & apply the instructions
	for scanner.Scan() {
		i := parseInstruction(scanner.Text())
		crates = i.apply(crates)
		cratesV2 = i.applyV2(cratesV2)
	}
	fmt.Println("final state (v1):", crates)
	fmt.Println("final state (v2):", cratesV2)

	part1 := ""
	for _, s := range crates {
		part1 += string(s[0])
	}
	fmt.Println("part 1:", part1)

	part2 := ""
	for _, s := range cratesV2 {
		part2 += string(s[0])
	}
	fmt.Println("part 2:", part2)
}

func parseCrateLine(line string, crates []string) []string {
	n := (len(line) + 1) / 4
	for i := 0; i < n; i++ {
		part := string(line[(4*i)+1])
		if len(crates)-1 < i {
			crates = append(crates, "")
		}
		if part == " " {
			continue
		}
		crates[i] = crates[i] + part
	}
	return crates
}

type instruction struct {
	n, from, to int
}

func (i *instruction) apply(crates []string) []string {
	for x := 0; x < i.n; x++ {
		crates[i.to] = crates[i.from][:1] + crates[i.to]
		crates[i.from] = crates[i.from][1:]
	}
	return crates
}

func (i *instruction) applyV2(crates []string) []string {
	crates[i.to] = crates[i.from][:i.n] + crates[i.to]
	crates[i.from] = crates[i.from][i.n:]
	return crates
}

func aToI(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return i
}

var r = regexp.MustCompile(`\d+`)

func parseInstruction(line string) *instruction {
	matches := r.FindAllString(line, 3)
	return &instruction{
		n:    aToI(matches[0]),
		from: aToI(matches[1]) - 1,
		to:   aToI(matches[2]) - 1,
	}
}
