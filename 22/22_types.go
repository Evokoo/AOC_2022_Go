package day22

import (
	"fmt"
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
type Position struct {
	xy      Point
	bearing int
}

type Grid struct {
	wraps map[Position]Position
	walls map[Point]struct{}
	start Point
}

func NewGrid() Grid {
	return Grid{
		wraps: make(map[Position]Position),
		walls: make(map[Point]struct{}),
		start: Point{-1, -1},
	}
}

func (g *Grid) AddWrap(from, to Point, facing, turn int) {
	(*g).wraps[Position{from, facing}] = Position{to, turn}
}

//	func (g *Grid) UpdateWraps(rows, cols map[int][2]int) {
//		for y, x := range rows {
//			// Left to Right
//			(*g).wraps[Point{x[0] - 1, y}] = Point{x[1], y}
//			// Right to Left
//			(*g).wraps[Point{x[1] + 1, y}] = Point{x[0], y}
//		}
//		for x, y := range cols {
//			// Top to Bottom
//			(*g).wraps[Point{x, y[0] - 1}] = Point{x, y[1]}
//			// Bottom to Top
//			(*g).wraps[Point{x, y[1] + 1}] = Point{x, y[0]}
//		}
//	}
func (g *Grid) AddWall(point Point) {
	(*g).walls[point] = struct{}{}
}

//	func (g Grid) isWrap(point Point) Point {
//		if newPoint, found := g.wraps[point]; found {
//			return newPoint
//		} else {
//			return point
//		}
//	}
//
//	func (g Grid) isWall(point Point) bool {
//		_, found := g.walls[point]
//		return found
//	}
func (g *Grid) CheckPoint(pos Position) (Position, bool) {
	if _, isWall := g.walls[pos.xy]; isWall {
		return pos, isWall
	} else if target, isWrap := g.wraps[pos]; isWrap {
		if _, destIsWall := g.walls[target.xy]; destIsWall {
			return pos, destIsWall
		}
		return target, false
	}

	return pos, false
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

	if s.pos.x <= -2 || s.pos.x >= 152 || s.pos.y <= -2 || s.pos.y >= 202 {
		fmt.Println(s.pos)
		panic("OUT OF BOUNDS")
	}

}
func (s *State) UpdatePosition(steps int, grid Grid) {
	for i := 0; i < steps; i++ {
		next := s.NextPosition(grid)

		// fmt.Println(next)

		if target, isWall := grid.CheckPoint(next); isWall {
			break
		} else {
			s.Turn(target.bearing)
			s.SetPosition(target.xy)
			fmt.Printf("x: %d, y: %d, bearing: %d\n", s.pos.x, s.pos.y, s.bearing)
		}
	}
}
func (s *State) NextPosition(grid Grid) Position {
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

	return Position{xy: next, bearing: (*s).bearing}
}
func (s *State) Turn(n int) {
	(*s).bearing = (s.bearing + n + 360) % 360
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
func (s *State) SetPosition(point Point) {
	(*s).pos = point
}

// ========================
// PARSER
// ========================
func ParseInput(file string, part int) (Grid, Moves) {
	data := tools.ReadFile(file)
	sections := strings.Split(data, "\n\n")
	lines := strings.Split(sections[0], "\n")

	// rows := make(map[int][2]int)
	// cols := make(map[int][2]int)

	//Size depends on the net shape!!
	// Input is 2 * 6
	// Example is 3 * 4
	size := len(lines) / 3
	grid := NewGrid()

	for y, line := range lines {
		for x, r := range line {

			if grid.start.x == -1 && r == '.' {
				grid.start = Point{x + 1, y + 1}
			}

			// if r == ' ' {
			// 	continue
			// }

			if r == '#' {
				grid.AddWall(Point{x + 1, y + 1})
			}

			// if limits, ok := rows[y]; ok {
			// 	if x < limits[0] {
			// 		limits[0] = x
			// 	}
			// 	if x > limits[1] {
			// 		limits[1] = x
			// 	}
			// 	rows[y] = limits
			// } else {
			// 	rows[y] = [2]int{x, x}
			// }

			// if limits, ok := cols[x]; ok {
			// 	if y < limits[0] {
			// 		limits[0] = y
			// 	}
			// 	if y > limits[1] {
			// 		limits[1] = y
			// 	}
			// 	cols[x] = limits
			// } else {
			// 	cols[x] = [2]int{y, y}
			// }
		}
	}

	fmt.Printf("size: %d, start: %v\n", size, grid.start)
	// fmt.Println(grid.start)

	if file == "example.txt" {
		if part == 2 {
			generateExampleCubeMap(size, grid)
		}
	}

	// if file == "input.txt" && part == 2 {
	// 	generateInputCubeMap(size, grid)
	// }

	return grid, NewMoveList(sections[1])
}
