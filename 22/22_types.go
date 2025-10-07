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
type Move struct {
	amount int
	isTurn bool
}

func NewMoveList(moveString string) []Move {
	var moves []Move
	for _, move := range tools.QuickMatch(moveString, `\d+|[RL]`) {
		switch move {
		case "R":
			moves = append(moves, Move{90, true})
		case "L":
			moves = append(moves, Move{-90, true})
		default:
			n, _ := strconv.Atoi(move)
			moves = append(moves, Move{n, false})
		}
	}
	return moves
}

// ========================
// GRID
// ========================
type Wrap struct {
	position Point
	bearing  int
}
type Grid struct {
	wraps map[Wrap]Wrap
	walls map[Point]struct{}
	start Point
}

func NewGrid() Grid {
	return Grid{
		wraps: make(map[Wrap]Wrap),
		walls: make(map[Point]struct{}),
		start: Point{-1, -1},
	}
}

func (g *Grid) AddWrap(from, to Point, bearing, turn int) {
	(*g).wraps[Wrap{from, bearing}] = Wrap{to, turn}
}
func (g *Grid) AddWall(point Point) {
	(*g).walls[point] = struct{}{}
}
func (g *Grid) IsWall(point Point) bool {
	_, isWall := g.walls[point]
	return isWall
}
func (g *Grid) IsWrap(wrapFrom Wrap) (Wrap, bool) {
	if wrapTo, isWrap := g.wraps[wrapFrom]; isWrap {
		return wrapTo, isWrap
	} else {
		return Wrap{}, isWrap
	}
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
func (s *State) UpdateState(move Move, grid Grid) {
	if move.isTurn {
		s.Turn(move.amount)
	} else {
		for i := 0; i < move.amount; i++ {
			nextPosition := s.NextPosition()

			if grid.IsWall(nextPosition) {
				break
			} else if wrapTo, isWrap := grid.IsWrap(Wrap{nextPosition, s.bearing}); isWrap {
				if grid.IsWall(wrapTo.position) {
					break
				}
				s.SetPosition(wrapTo.position)
				s.Turn(wrapTo.bearing)
			} else {
				s.SetPosition(nextPosition)
			}
		}
	}
}
func (s *State) NextPosition() Point {
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

	return next
}
func (s *State) Turn(n int) {
	(*s).bearing = (s.bearing + n + 360) % 360
}
func (s *State) SetPosition(point Point) {
	(*s).pos = point
}
func (s *State) GetPassword() int {
	row := s.pos.y
	col := s.pos.x
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
func ParseInput(file string, part int) (Grid, []Move, State) {
	data := tools.ReadFile(file)
	sections := strings.Split(data, "\n\n")

	grid := NewGrid()

	lines := strings.Split(sections[0], "\n")
	for y, line := range lines {
		for x, r := range line {
			if grid.start.x == -1 && r == '.' {
				grid.start = Point{x + 1, y + 1}
			}
			if r == '#' {
				grid.AddWall(Point{x + 1, y + 1})
			}
		}
	}

	if part == 1 {
		generate2DWrapMap(len(lines)/4, grid)
	}
	if part == 2 {
		generate3DWrapMap(len(lines)/4, grid)
	}

	return grid, NewMoveList(sections[1]), NewState(grid.start)
}
