package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

var root = NewDir("/", nil)

func main() {
	f, err := os.Open("day7/terminal.txt")
	//f, err := os.Open("day7/simple.txt")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(f)
	// Skip the first line - we'll just assume it's "% cd /"
	_ = scanner.Text()

	current := root
	listing := false
	for scanner.Scan() {
		text := scanner.Text()
		fmt.Println(text)
		// Check if line is a command or output
		if text[0] == '$' {
			listing = false
			parts := strings.Split(text, " ")
			if parts[1] == "cd" {
				current = changeDirectory(current, parts[2])
			} else if parts[1] == "ls" {
				listing = true
			}
		} else if listing {
			// TODO: track seen but unlisted directories?
			if parts := strings.Split(text, " "); parts[0] != "dir" {
				current.AddFile(parts[0], parts[1])
			}
		}
	}

	root.Tree()

	fmt.Println("root size:", root.CalcSize())
	dirSizes := root.GetDirSizes()
	sort.Slice(dirSizes, func(i, j int) bool {
		return dirSizes[i].size < dirSizes[j].size
	})

	target := 30_000_000 - (70_000_000 - dirSizes[len(dirSizes)-1].size)
	fmt.Println("target:", target)

	sum := 0
	toDelete := 0
	for _, ds := range dirSizes {
		if ds.size <= 100000 {
			sum += ds.size
		}
		if ds.size >= target && toDelete == 0 {
			toDelete = ds.size
		}
	}
	fmt.Println("part 1:", sum)
	fmt.Println("part 2:", toDelete)
}

func changeDirectory(current *Dir, name string) *Dir {
	if name == "/" {
		return root
	} else if name == ".." {
		return current.parent
	} else {
		return current.AddChild(name)
	}
}

func MakeIndent(levels int) string {
	indent := ""
	for i := 0; i < levels*2; i++ {
		indent += " "
	}
	return indent
}
