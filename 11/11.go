package day11

import (
	"sort"
	"strconv"
	"strings"

	"github.com/Evokoo/AOC_2022_Go/tools"
)

type Operation [3]string

type Item int

func (i Item) get() int {
	return int(i)
}
func (i *Item) set(n int) {
	*i = Item(n)
}

type Monkey struct {
	items       []Item
	operation   Operation
	test        int
	desination  [2]int
	inspections int
	lcm         int
}

func (m *Monkey) push(item Item) {
	(*m).items = append((*m).items, item)
}
func (m *Monkey) pop() Item {
	item := m.items[0]
	m.items = m.items[1:]
	return item
}
func (m *Monkey) hasItems() bool {
	return len((*m).items) > 0
}

// Inspect and update item and determine destination
func (m *Monkey) inspect(item *Item, part int) int {
	m.inspections++

	var a, b int

	if m.operation[0] == "old" {
		a = item.get()
	} else {
		n, _ := strconv.Atoi(m.operation[0])
		a = n
	}
	if m.operation[2] == "old" {
		b = item.get()
	} else {
		n, _ := strconv.Atoi(m.operation[2])
		b = n
	}

	switch m.operation[1] {
	case "+":
		if part == 2 {
			(*item).set((a + b) % m.lcm)
		} else {
			(*item).set((a + b) / 3)
		}
	case "*":
		if part == 2 {
			(*item).set((a * b) % m.lcm)
		} else {
			(*item).set((a * b) / 3)
		}
	}

	if (*item).get()%m.test == 0 {
		return m.desination[0]
	} else {
		return m.desination[1]
	}
}

func Solve(file string, part int) int {
	monkeys := parseInput(file)
	rounds := 20

	if part == 2 {
		rounds = 10000
	}

	for range rounds {
		for i := range monkeys {
			for monkeys[i].hasItems() {
				item := monkeys[i].pop()
				index := monkeys[i].inspect(&item, part)
				monkeys[index].push(item)
			}
		}
	}

	sort.Slice(monkeys, func(i, j int) bool {
		return monkeys[i].inspections > monkeys[j].inspections
	})

	return monkeys[0].inspections * monkeys[1].inspections
}
func parseInput(file string) []Monkey {
	var output []Monkey
	data := tools.ReadFile(file)
	lcm := 1

	for section := range strings.SplitSeq(data, "\n\n") {
		var monkey Monkey

		for i, line := range strings.Split(section, "\n") {
			switch i {
			case 1:
				for _, s := range tools.QuickMatch(line, `\d+`) {
					n, _ := strconv.Atoi(s)
					monkey.push(Item(n))
				}
			case 2:
				fields := strings.Fields(line)
				monkey.operation = Operation{fields[3], fields[4], fields[5]}
			case 3:
				d := tools.QuickMatch(line, `\d+`)
				n, _ := strconv.Atoi(d[0])
				monkey.test = n
				lcm *= n
			case 4, 5:
				d := tools.QuickMatch(line, `\d+`)
				n, _ := strconv.Atoi(d[0])

				if i == 4 {
					monkey.desination[0] = n
				} else {
					monkey.desination[1] = n
				}
			}
		}
		output = append(output, monkey)
	}

	for i := range output {
		output[i].lcm = lcm
	}

	return output
}
