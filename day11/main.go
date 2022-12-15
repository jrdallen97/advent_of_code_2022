package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Monkey struct {
	Items           []int
	Operation       string
	Divisor         int
	IfTrue, IfFalse int
	inspectCount    int
}

func (m *Monkey) InspectV1(item int) int {
	m.inspectCount++
	return m.increaseWorry(item) / 3
}

func (m *Monkey) InspectV2(item int, modulo int) int {
	m.inspectCount++
	return m.increaseWorry(item) % modulo
}

func (m *Monkey) increaseWorry(worry int) int {
	op := strings.ReplaceAll(m.Operation, "old", fmt.Sprintf("%d", worry))
	parts := strings.Split(op, " ")
	x, _ := strconv.Atoi(parts[0])
	y, _ := strconv.Atoi(parts[2])
	if parts[1] == "*" {
		return x * y
	}
	return x + y
}

func (m *Monkey) ThrowTo(item int) int {
	if item%m.Divisor == 0 {
		return m.IfTrue
	}
	return m.IfFalse
}

func main() {
	f, err := os.Open("day11/monkeys.txt")
	//f, err := os.Open("day11/simple.txt")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(f)
	var monkeys []*Monkey
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if strings.HasPrefix(line, "Monkey") {
			monkeys = append(monkeys, &Monkey{})
			continue
		}

		monkey := monkeys[len(monkeys)-1]
		if strings.HasPrefix(line, "Starting items: ") {
			items := strings.Split(strings.TrimPrefix(line, "Starting items: "), ", ")
			monkey.Items = make([]int, len(items))
			for i, item := range items {
				monkey.Items[i], _ = strconv.Atoi(item)
			}
		} else if strings.HasPrefix(line, "Operation: ") {
			monkey.Operation = strings.TrimPrefix(line, "Operation: new = ")
		} else if strings.HasPrefix(line, "Test: ") {
			cond := strings.TrimPrefix(line, "Test: divisible by ")
			monkey.Divisor, _ = strconv.Atoi(cond)
		} else if strings.HasPrefix(line, "If true:") {
			throwTo := strings.TrimPrefix(line, "If true: throw to monkey ")
			monkey.IfTrue, _ = strconv.Atoi(throwTo)
		} else if strings.HasPrefix(line, "If false:") {
			throwTo := strings.TrimPrefix(line, "If false: throw to monkey ")
			monkey.IfFalse, _ = strconv.Atoi(throwTo)
		}
	}

	// By modulo-ing by a common denominator, we can stay under max int without
	// changing behaviour.
	commonDenominator := 1
	for _, m := range monkeys {
		commonDenominator *= m.Divisor
	}

	//rounds := 20 // Part 1
	rounds := 10000 // Part 2
	for round := 1; round <= rounds; round++ {
		for _, monkey := range monkeys {
			for _, item := range monkey.Items {
				// Adjust worry
				//item = monkey.InspectV1(item) // Part 1
				item = monkey.InspectV2(item, commonDenominator) // Part 2
				// Throw item
				throwTo := monkey.ThrowTo(item)
				monkeys[throwTo].Items = append(monkeys[throwTo].Items, item)
			}
			// Clear monkey's items; they've all been thrown
			monkey.Items = []int{}
		}
	}

	var inspects []int
	for i, m := range monkeys {
		inspects = append(inspects, m.inspectCount)
		fmt.Printf("monkey %d items: %v\n", i, m.Items)
		fmt.Printf("monkey %d inspects: %d\n", i, m.inspectCount)
	}
	sort.Ints(inspects)
	monkeyBusiness := inspects[len(inspects)-1] * inspects[len(inspects)-2]
	fmt.Println("monkey business:", monkeyBusiness)
}
