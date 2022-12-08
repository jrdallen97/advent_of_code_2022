package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func aToI(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return i
}

func newSections(s string) *sections {
	split := strings.Split(s, "-")
	return &sections{
		start: aToI(split[0]),
		end:   aToI(split[1]),
	}
}

type sections struct {
	start, end int
}

func (s1 *sections) contains(s2 *sections) bool {
	return s2.start >= s1.start && s2.end <= s1.end
}

func (s1 *sections) toMap() map[int]struct{} {
	m := map[int]struct{}{}
	for i := s1.start; i <= s1.end; i++ {
		m[i] = struct{}{}
	}
	return m
}

func (s1 *sections) overlaps(s2 *sections) bool {
	m1, m2 := s1.toMap(), s2.toMap()
	for i := range m1 {
		if _, ok := m2[i]; ok {
			return true
		}
	}
	return false
}

func main() {
	f, err := os.Open("day4/pairs.txt")
	//f, err := os.Open("day4/simple.txt")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(f)
	fullyContained := 0
	overlaps := 0
	for scanner.Scan() {
		elves := strings.Split(scanner.Text(), ",")
		e1 := newSections(elves[0])
		e2 := newSections(elves[1])
		if e1.contains(e2) || e2.contains(e1) {
			fullyContained++
		}
		if e1.overlaps(e2) {
			overlaps++
		}
	}
	fmt.Println("fully contained:", fullyContained)
	fmt.Println("overlaps:", overlaps)
}
