package day19

import (
	"strconv"
	"strings"

	"github.com/Evokoo/AOC_2022_Go/tools"
)

// ========================
// CONSTANTS
// ========================

// ========================
// SUPPLY LIST
// ========================
type SupplyList struct {
	ore      int
	clay     int
	obsidian int
}

func NewSupplyList(ore, clay, obsidian int) SupplyList {
	return SupplyList{ore, clay, obsidian}
}

// ========================
// BLUEPRINT
// ========================
type Blueprint struct {
	id       int
	ore      SupplyList
	clay     SupplyList
	obsidian SupplyList
	geode    SupplyList
}

func NewBlueprint(values []int) Blueprint {
	return Blueprint{
		id:       values[0],
		ore:      NewSupplyList(values[1], 0, 0),
		clay:     NewSupplyList(values[2], 0, 0),
		obsidian: NewSupplyList(values[3], values[4], 0),
		geode:    NewSupplyList(values[5], 0, values[6]),
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

// ========================
// ROBOTS
// ========================
type Robots map[string]int

// ========================
// MATERIALS
// ========================
type Materials struct {
	ore      int
	clay     int
	obsidian int
	geode    int
}

func (m *Materials) UpdateMaterials(robots Robots) {
	for robot, count := range robots {
		switch robot {
		case "ore":
			m.ore += count
		case "clay":
			m.clay += count
		case "obsidian":
			m.obsidian += count
		case "geode":
			m.geode += count
		}
	}
}
