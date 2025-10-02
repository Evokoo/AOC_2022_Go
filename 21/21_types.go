package day21

import (
	"strconv"
	"strings"

	"github.com/Evokoo/AOC_2022_Go/tools"
)

// ========================
// VALUE MAP
// ========================
type Values map[string]int

func (v Values) Has(key string) bool {
	_, found := v[key]
	return found
}
func (v *Values) Add(key string, value int) {
	(*v)[key] = value
}
func (v Values) Get(key string) int {
	return v[key]
}

// ========================
// EQUATION
// ========================
type Equation [4]string

func (e Equation) Solve(values Values) (int, bool) {
	if values.Has(e[1]) && values.Has(e[3]) {
		a, b := values.Get(e[1]), values.Get(e[3])

		switch e[2] {
		case "+":
			return a + b, true
		case "-":
			return a - b, true
		case "*":
			return a * b, true
		case "/":
			return a / b, true
		}
	}

	return 0, false
}

// ========================
// TRACE LOG
// ========================
type TraceLog struct {
	target string
	log    [][3]string
}

func NewTracer(target string) TraceLog {
	return TraceLog{target: target}
}
func (t *TraceLog) Log(e Equation, v Values) {
	contains := false

	for i, part := range e {
		if part == t.target {
			e[i] = "X"
			contains = true
		} else if (i != 0) && (i != 2) {
			e[i] = strconv.Itoa(v.Get(part))
		}
	}

	if contains {
		t.target = e[0]
		t.log = append(t.log, [3]string{e[1], e[2], e[3]})
	}
}

// ========================
// PARSER
// ========================

func ParseInput(file string) (Values, Queue[Equation]) {
	data := tools.ReadFile(file)

	values := make(Values)
	equations := NewQueue[Equation]()

	for line := range strings.SplitSeq(data, "\n") {
		parts := tools.QuickMatch(line, `[0-9a-z\+\-\*\/]+`)
		value, err := strconv.Atoi(parts[1])

		if err == nil {
			values[parts[0]] = value
		} else {
			equations.Push(Equation(parts))
		}
	}

	return values, equations
}

// ========================
// QUEUE
// ========================
type Queue[T any] []T

func NewQueue[T any]() Queue[T] {
	return Queue[T]{}
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

// ========================
// SOLVE FOR HUMAN
// ========================
func FindHumanValue(log [][3]string) int {
	value := log[len(log)-1][2]

	for i := len(log) - 2; i >= 0; i-- {
		equation := ReverseEquation(log[i])
		value = SolveForHuman(value, equation)
	}

	result, _ := strconv.Atoi(value)

	return result
}
func ReverseEquation(e [3]string) [3]string {
	switch e[1] {
	case "+":
		e[1] = "-"
		if e[2] == "X" {
			e[0], e[2] = e[2], e[0]
		}
	case "-":
		if e[0] == "X" {
			e[1] = "+"
		}
	case "*":
		e[1] = "/"
		if e[2] == "X" {
			e[0], e[2] = e[2], e[0]
		}
	case "/":
		e[1] = "*"
	}

	return [3]string{e[0], e[1], e[2]}
}
func SolveForHuman(value string, equation [3]string) string {
	for i, s := range equation {
		if s == "X" {
			equation[i] = value
		}
	}
	a, _ := strconv.Atoi(equation[0])
	b, _ := strconv.Atoi(equation[2])

	switch equation[1] {
	case "*":
		return strconv.Itoa(a * b)
	case "/":
		return strconv.Itoa(a / b)
	case "+":
		return strconv.Itoa(a + b)
	case "-":
		return strconv.Itoa(a - b)
	default:
		panic("Invalid operator")
	}
}
