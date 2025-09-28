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
		valves[details[0]] = NewValve(i, details[0], rate, details[2:], false)
	}

	mapDistance(&valves)
	return valves
}
func mapDistance(valves *ValveMap) {

	for id := range *valves {
		queue := NewQueue(DistanceState{id, 0})
		dists := make(map[string]int)
		for !queue.IsEmpty() {
			current := queue.Pop()
			if _, found := dists[current.location]; found {
				continue
			} else {
				dists[current.location] = current.distance
			}
			for _, path := range (*valves)[current.location].paths {
				if _, found := dists[path]; !found {
					queue.Push(DistanceState{path, current.distance + 1})
				}

			}
		}
		(*valves)[id].UpdateDistanceMap(dists)
	}
}

type State struct {
	id       string
	time     int
	pressure int
	openMask uint64
}

func traverseValves(valves ValveMap, time int) int {
	valveList := valves.RemoveZeroFlowValves()
	queue := NewQueue(State{"AA", time, 0, 0})
	bestSeen := make(map[[2]uint64]int)
	best := 0

	for !queue.IsEmpty() {
		current := queue.Pop()

		if current.pressure > best {
			best = current.pressure
		}

		for id, valve := range valveList {
			if current.openMask&(1<<valve.index) != 0 {
				continue
			}

			newMask := current.openMask | (1 << valve.index)
			distTo := valve.distance[current.id]
			newTime := current.time - (distTo + 1)
			newPressure := current.pressure + (newTime * valve.rate)

			if newTime < 0 {
				continue
			}

			// --- pruning: skip worse duplicate states ---
			key := [2]uint64{uint64(valve.index), newMask}
			if bestSeen[key] >= newPressure {
				continue
			}
			bestSeen[key] = newPressure
			// -------------------------------------------

			queue.Push(State{id, newTime, newPressure, newMask})
		}
	}

	return best
}
