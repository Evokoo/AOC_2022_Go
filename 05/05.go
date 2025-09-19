package day05

import (
	"strconv"
	"strings"

	"github.com/Evokoo/AOC_2022_Go/tools"
	"github.com/Evokoo/AOC_2022_Go/tools/ds"
)

type Move struct {
	Amount int
	From   int
	To     int
}

func moveItems(from *ds.Stack[string], amount int, to *ds.Stack[string], preserveOrder bool) {
	var items []string

	for range amount {
		item, err := (*from).Pop()
		if err == nil {
			items = append(items, item)

		}
	}
	if preserveOrder {
		for i, j := 0, len(items)-1; i < j; i, j = i+1, j-1 {
			items[i], items[j] = items[j], items[i]
		}
	}
	for _, item := range items {
		to.Push(item)
	}
}

func Solve(file string, part int) string {
	data := tools.ReadFile(file)
	stacks, moves := parseInput(data)
	preserveOrder := false

	if part == 2 {
		preserveOrder = true
	}

	for _, move := range moves {
		moveItems(&stacks[move.From], move.Amount, &stacks[move.To], preserveOrder)
	}

	var result strings.Builder
	for _, stack := range stacks {
		if item, err := stack.Peek(); err == nil {
			result.WriteString(item)
		}
	}
	return result.String()
}

func parseInput(data string) ([]ds.Stack[string], []Move) {
	var stacks []ds.Stack[string]
	var moves []Move

	for i, section := range strings.Split(data, "\n\n") {
		lines := strings.Split(section, "\n")

		if i == 0 {
			for x := 1; x < len(lines[0]); x = x + 4 {
				var stack ds.Stack[string]
				for y := 0; y < len(lines)-1; y++ {
					if lines[y][x] != ' ' {
						stack = append(stack, string(lines[y][x]))
					}
				}
				stacks = append(stacks, stack)
			}
		} else {
			for _, line := range lines {
				var move Move
				for i, s := range tools.QuickMatch(line, `\d+`) {
					digit, _ := strconv.Atoi(s)
					switch i {
					case 0:
						move.Amount = digit
					case 1:
						move.From = digit - 1
					case 2:
						move.To = digit - 1
					}
				}
				moves = append(moves, move)
			}

		}
	}
	return stacks, moves
}
