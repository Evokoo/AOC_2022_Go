package day22

import (
	"strconv"
	"strings"

	"github.com/Evokoo/AOC_2022_Go/tools"
)

// ========================
// POINT
// ========================
type Point struct{ x, y int }

// ========================
// MOVE
// ========================
type Moves []int

func NewMoveList(moveString string) Moves {
	var moves Moves

	for _, move := range tools.QuickMatch(moveString, `\d+|[RL]`) {
		switch move {
		case "R":
			moves = append(moves, 90)
		case "L":
			moves = append(moves, -90)
		default:
			n, _ := strconv.Atoi(move)
			moves = append(moves, n)
		}
	}

	return moves
}

// ========================
// GRID
// ========================
type Grid struct {
	wraps map[Point]Point
	walls map[Point]struct{}
	start Point
}

func NewGrid() Grid {
	return Grid{
		wraps: make(map[Point]Point),
		walls: make(map[Point]struct{}),
		start: Point{-1, -1},
	}
}
func (g *Grid) UpdateWraps(rows, cols map[int][2]int) {
	for y, x := range rows {
		// Left to Right
		(*g).wraps[Point{x[0] - 1, y}] = Point{x[1], y}
		// Right to Left
		(*g).wraps[Point{x[1] + 1, y}] = Point{x[0], y}
	}
	for x, y := range cols {
		// Top to Bottom
		(*g).wraps[Point{x, y[0] - 1}] = Point{x, y[1]}
		// Bottom to Top
		(*g).wraps[Point{x, y[1] + 1}] = Point{x, y[0]}
	}
}
func (g *Grid) AddWall(point Point) {
	(*g).walls[point] = struct{}{}
}
func (g Grid) isWrap(point Point) Point {
	if newPoint, found := g.wraps[point]; found {
		return newPoint
	} else {
		return point
	}
}
func (g Grid) isWall(point Point) bool {
	_, found := g.walls[point]
	return found
}

// ========================
// STATE
// ========================
type State struct {
	pos     Point
	bearing int
}

func NewState(start Point) State {
	return State{pos: start, bearing: 90}
}
func (s *State) UpdateState(move int, grid Grid, index int) {
	if index%2 == 0 {
		s.UpdatePosition(move, grid)
	} else {
		s.Turn(move)
	}
}
func (s *State) UpdatePosition(steps int, grid Grid) {
	for i := 0; i < steps; i++ {
		next := s.NextPosition(grid)
		if grid.isWall(next) {
			break
		} else {
			(*s).pos = next
		}
	}
}
func (s *State) NextPosition(grid Grid) Point {
	next := Point{(*s).pos.x, (*s).pos.y}

	switch (*s).bearing {
	case 0:
		next.y--
	case 90:
		next.x++
	case 180:
		next.y++
	case 270:
		next.x--
	default:
		panic("Invalid Bearing")
	}

	return grid.isWrap(next)
}
func (s *State) Turn(n int) {
	(*s).bearing = (s.bearing + n + 360) % 360
}
func (s *State) GetPassword() int {
	row := s.pos.y + 1
	col := s.pos.x + 1
	dir := 0

	switch s.bearing {
	case 0:
		dir = 3
	case 90:
		dir = 0
	case 180:
		dir = 1
	case 270:
		dir = 2
	}

	return row*1000 + col*4 + dir
}

// ========================
// PARSER
// ========================
func ParseInput(file string) (Grid, Moves) {
	data := tools.ReadFile(file)
	sections := strings.Split(data, "\n\n")

	rows := make(map[int][2]int)
	cols := make(map[int][2]int)
	grid := NewGrid()

	for y, line := range strings.Split(sections[0], "\n") {
		for x, r := range line {

			if grid.start.x == -1 && r == '.' {
				grid.start = Point{x, y}
			}

			if r == ' ' {
				continue
			}

			if r == '#' {
				grid.AddWall(Point{x, y})
			}

			if limits, ok := rows[y]; ok {
				if x < limits[0] {
					limits[0] = x
				}
				if x > limits[1] {
					limits[1] = x
				}
				rows[y] = limits
			} else {
				rows[y] = [2]int{x, x}
			}

			if limits, ok := cols[x]; ok {
				if y < limits[0] {
					limits[0] = y
				}
				if y > limits[1] {
					limits[1] = y
				}
				cols[x] = limits
			} else {
				cols[x] = [2]int{y, y}
			}
		}
	}

	grid.UpdateWraps(rows, cols)

	return grid, NewMoveList(sections[1])
}
