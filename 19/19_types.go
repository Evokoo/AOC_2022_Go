package day19

import (
	"errors"
	"strconv"
	"strings"

	"github.com/Evokoo/AOC_2022_Go/tools"
)

// ========================
// CONSTANTS
// ========================
type Resource int

const (
	ORE Resource = iota
	CLAY
	OBSIDAN
	GEODE
)

// ========================
// Recipe
// ========================
type Recipe [3]int

func NewRecipe(ore, clay, obsidian int) Recipe {
	return Recipe{ore, clay, obsidian}
}

// ========================
// BLUEPRINT
// ========================
type Blueprint struct {
	id      int
	recipes [4]Recipe
}

func NewBlueprint(values []int) Blueprint {
	return Blueprint{
		id: values[0],
		recipes: [4]Recipe{
			NewRecipe(values[1], 0, 0),
			NewRecipe(values[2], 0, 0),
			NewRecipe(values[3], values[4], 0),
			NewRecipe(values[5], 0, values[6]),
		},
	}
}

func ParseBlueprints(file string) (blueprints []Blueprint) {
	data := tools.ReadFile(file)

	for line := range strings.SplitSeq(data, "\n") {
		digits := tools.QuickMatch(line, `\d+`)
		values := make([]int, len(digits))
		for i, d := range digits {
			n, _ := strconv.Atoi(d)
			values[i] = n
		}

		blueprints = append(blueprints, NewBlueprint(values))
	}
	return
}

func (b Blueprint) GetRecipe(robot Resource) Recipe {
	return b.recipes[robot]
}

func (b Blueprint) MaxResourceCap() Materials {
	var target Materials

	for _, recipe := range b.recipes {
		for i, value := range recipe {
			if target[i] < value {
				target[i] = value
			}
		}
	}
	return target
}

// ========================
// ROBOTS
// ========================
type Robots [4]int

func NewRobotMap() Robots {
	return Robots{1, 0, 0, 0}
}

func (r *Robots) Build(robot Resource, m *Materials, blueprint Blueprint) (Resource, error) {
	recipe := blueprint.GetRecipe(robot)
	if m.CanBuild(recipe) {
		m.UseMaterials(recipe)
		return robot, nil
	}
	return -1, errors.New("insufficient materials")
}

func (r *Robots) Deploy(robot Resource) {
	(*r)[robot]++
}

func (r Robots) Collect(m *Materials) {
	for i, count := range r {
		(*m)[i] += count
	}
}

// ========================
// MATERIALS
// ========================
type Materials [4]int

func NewMaterialList() Materials {
	return Materials{}
}

func (m Materials) CanBuild(recipe Recipe) bool {
	for i, count := range recipe {
		if m[i] < count {
			return false
		}
	}
	return true
}

func (m *Materials) UseMaterials(recipe Recipe) {
	for i, amount := range recipe {
		(*m)[i] -= amount
	}
}

// ========================
// SIMULATION
// ========================
type State struct {
	time      int
	robots    Robots
	materials Materials
}

func NewState(time int) State {
	return State{
		time:      time,
		robots:    NewRobotMap(),
		materials: NewMaterialList(),
	}
}

func (s State) CopyState() State {
	robots := s.robots
	materials := s.materials
	return State{s.time, robots, materials}
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
