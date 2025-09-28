package day16

// ========================
// VALVE
// ========================
type ValveMap map[string]*Valve
type Valve struct {
	index    int
	id       string
	rate     int
	paths    []string
	isOpen   bool
	distance map[string]int
}

func NewValve(index int, id string, rate int, paths []string, open bool) *Valve {
	return &Valve{index, id, rate, paths, open, make(map[string]int)}
}

func (v *Valve) UpdateDistanceMap(distanceMap DistanceMap) {
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
// DISTANCE
// ========================
type DistanceMap = map[string]int
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
