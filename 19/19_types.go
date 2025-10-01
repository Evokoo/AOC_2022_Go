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
const (
	ORE     = "ore"
	CLAY    = "clay"
	OBSIDAN = "obsidian"
	GEODE   = "geode"
)

var RobotTypes = []string{ORE, CLAY, OBSIDAN, GEODE}

// ========================
// Recipe
// ========================
type Recipe struct {
	ore      int
	clay     int
	obsidian int
}

func NewRecipe(ore, clay, obsidian int) Recipe {
	return Recipe{ore, clay, obsidian}
}

// ========================
// BLUEPRINT
// ========================
type Blueprint struct {
	id       int
	ore      Recipe
	clay     Recipe
	obsidian Recipe
	geode    Recipe
}

func NewBlueprint(values []int) Blueprint {
	return Blueprint{
		id:       values[0],
		ore:      NewRecipe(values[1], 0, 0),
		clay:     NewRecipe(values[2], 0, 0),
		obsidian: NewRecipe(values[3], values[4], 0),
		geode:    NewRecipe(values[5], 0, values[6]),
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

func (b Blueprint) GetRecipe(robot string) Recipe {
	switch robot {
	case ORE:
		return b.ore
	case CLAY:
		return b.clay
	case OBSIDAN:
		return b.obsidian
	case GEODE:
		return b.geode
	default:
		panic("Invalid robot type")
	}
}

// ========================
// ROBOTS
// ========================
type Robots map[string]int

func NewRobotMap() Robots {
	robots := make(map[string]int, 4)
	robots[ORE]++
	return robots
}

func (r Robots) Clone() Robots {
	clone := make(Robots, len(r))
	for robot, count := range r {
		clone[robot] = count
	}
	return clone
}

func (r *Robots) Build(robot string, m *Materials, blueprint Blueprint) (string, error) {
	recipe := blueprint.GetRecipe(robot)
	if m.CanBuild(recipe) {
		m.UseMaterials(recipe)
		return robot, nil
	}
	return "", errors.New("insufficient materials")
}

func (r *Robots) Deploy(robot string) {
	(*r)[robot]++
}

func (r Robots) Collect(m *Materials) {
	for robot, count := range r {
		switch robot {
		case ORE:
			m.ore += count
		case CLAY:
			m.clay += count
		case OBSIDAN:
			m.obsidian += count
		case GEODE:
			m.geode += count
		}
	}
}

// ========================
// MATERIALS
// ========================
type Materials struct {
	ore      int
	clay     int
	obsidian int
	geode    int
}

func NewMaterialList() Materials {
	return Materials{}
}

func (m Materials) CanBuild(recipe Recipe) bool {
	if m.ore >= recipe.ore &&
		m.clay >= recipe.clay &&
		m.obsidian >= recipe.obsidian {
		return true
	}
	return false
}

func (m *Materials) UseMaterials(recipe Recipe) {
	(*m).ore -= recipe.ore
	(*m).clay -= recipe.clay
	(*m).obsidian -= recipe.obsidian
}

// ========================
// SIMULATION
// ========================
type State struct {
	time      int
	robots    Robots
	materials Materials
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
