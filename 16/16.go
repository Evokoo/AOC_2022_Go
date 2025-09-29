package day16

import (
	"strconv"
	"strings"

	"github.com/Evokoo/AOC_2022_Go/tools"
)

func Solve(file string) int {
	valves := parseInput(file)

	return traverseValves(valves, 30)
}

func parseInput(file string) ValveMap {
	data := tools.ReadFile(file)
	valves := make(ValveMap)

	for i, line := range strings.Split(data, "\n") {
		details := tools.QuickMatch(line, `[A-Z]{2}|\d+`)
		rate, _ := strconv.Atoi(details[1])
		valves[details[0]] = NewValve(i, rate, details[2:])
	}

	mapDistance(&valves)
	return valves.RemoveZeroFlowValves()
}
func mapDistance(valves *ValveMap) {
	for id := range *valves {
		(*valves)[id].distance = make(map[string]int)
		queue := NewQueue(DistanceState{id, 0})

		for !queue.IsEmpty() {
			current := queue.Pop()
			if _, found := (*valves)[id].distance[current.location]; found {
				continue
			} else {
				(*valves)[id].distance[current.location] = current.distance
			}
			for _, path := range (*valves)[current.location].paths {
				if _, found := (*valves)[id].distance[path]; !found {
					queue.Push(DistanceState{path, current.distance + 1})
				}

			}
		}
	}
}

func traverseValves(valves ValveMap, time int) int {
	queue := NewQueue(State{"AA", time, 0, 0})
	bestSeen := make(map[[2]uint64]int)
	best := 0

	for !queue.IsEmpty() {
		current := queue.Pop()

		if current.pressure > best {
			best = current.pressure
		}

		for id, valve := range valves {
			if current.openMask&(1<<valve.index) != 0 {
				continue
			}

			newState := current.UpdateState(id, valve)
			key := [2]uint64{uint64(valve.index), newState.openMask}

			if newState.time < 0 || bestSeen[key] >= newState.pressure {
				continue
			}

			bestSeen[key] = newState.pressure
			queue.Push(newState)
		}
	}
	return best
}
