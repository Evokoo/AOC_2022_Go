package day16

// ========================
// VALVE
// ========================
type ValveMap map[string]*Valve
type Valve struct {
	index    int
	rate     int
	paths    []string
	distance map[string]int
}

func NewValve(index int, rate int, paths []string) *Valve {
	return &Valve{index, rate, paths, make(map[string]int)}
}

func (v *Valve) UpdateDistanceMap(distanceMap map[string]int) {
	v.distance = distanceMap
}

func (v ValveMap) RemoveZeroFlowValves() ValveMap {
	output := make(ValveMap)
	for id, valve := range v {
		if valve.rate != 0 || id == "AA" {
			output[id] = valve
		}
	}
	return output
}

// ========================
// STATE
// ========================
type State struct {
	id       string
	time     int
	pressure int
	openMask uint64
}

func (s State) UpdateState(id string, valve *Valve) State {
	mask := s.openMask | (1 << (*valve).index)
	time := s.time - ((*valve).distance[s.id] + 1)
	pressure := s.pressure + (time * (*valve).rate)
	return State{id, time, pressure, mask}
}

// ========================
// DISTANCE
// ========================
type DistanceState struct {
	location string
	distance int
}

// ========================
// QUEUE
// ========================

type Queue[T any] []T

func NewQueue[T any](item T) Queue[T] {
	return Queue[T]{item}
}
func (q *Queue[T]) Pop() T {
	removed := (*q)[0]
	(*q) = (*q)[1:]
	return removed
}
func (q *Queue[T]) Push(c T) {
	*q = append(*q, c)
}
func (q Queue[T]) IsEmpty() bool {
	return len(q) == 0
}
