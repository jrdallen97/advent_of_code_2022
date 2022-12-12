package main

import "fmt"

type VisibleFrom map[Direction]struct{}

type Visibility [][]VisibleFrom

func NewVisibility(lines [][]int) Visibility {
	v := make(Visibility, len(lines))

	// Work out visibility from the left & top, and initialise all values
	// Start top-left, read to the right, then down
	topMax := initMax(len(lines[0]))
	leftMax := initMax(len(lines))
	for i, row := range lines {
		v[i] = make([]VisibleFrom, len(row))
		for j, val := range row {
			v[i][j] = VisibleFrom{}
			if val > leftMax[i] {
				leftMax[i] = val
				v[i][j][Left] = struct{}{}
			}
			if val > topMax[j] {
				topMax[j] = val
				v[i][j][Top] = struct{}{}
			}
		}
	}

	// Work out visibility from the left & top, and initialise all values
	// Start bottom-right, read to the left, then up
	bottomMax := initMax(len(lines[0]))
	rightMax := initMax(len(lines))
	for i := len(lines) - 1; i >= 0; i-- {
		row := lines[i]
		//fmt.Println(row)
		for j := len(lines[i]) - 1; j >= 0; j-- {
			val := row[j]
			if val > rightMax[i] {
				rightMax[i] = val
				v[i][j][Right] = struct{}{}
			}
			if val > bottomMax[j] {
				bottomMax[j] = val
				v[i][j][Bottom] = struct{}{}
			}
		}
	}

	return v
}

func (v Visibility) Print(from Direction) {
	fmt.Println(from, "visibility:")
	for _, row := range v {
		for _, col := range row {
			if _, ok := col[from]; ok {
				fmt.Printf("1")
			} else {
				fmt.Printf("0")
			}
		}
		fmt.Println()
	}
	fmt.Println()
}

func (v Visibility) CountVisible() int {
	count := 0
	for _, row := range v {
		for _, col := range row {
			if len(col) > 0 {
				count++
			}
		}
	}
	return count
}

// initMax fills returns a slice of length n filled with -1
func initMax(n int) []int {
	s := make([]int, n)
	for i := range s {
		s[i] = -1
	}
	return s
}
