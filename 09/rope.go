package day09

import (
	"math"
)

type Point struct{ X, Y int }
type Rope []Point

func NewRope(length int, start Point) Rope {
	rope := make(Rope, length)
	for i := range rope {
		rope[i] = start
	}
	return rope
}

var DIRECTIONS = map[string]Point{
	"U": {0, 1}, "R": {1, 0}, "D": {0, -1}, "L": {-1, 0},
}

func (p *Point) update(dir string) {
	if d, found := DIRECTIONS[dir]; found {
		p.X += d.X
		p.Y += d.Y
	}
}
func (a Point) isAdjacent(b Point) bool {
	dx := int(math.Abs(float64(a.X) - float64(b.X)))
	dy := int(math.Abs(float64(a.Y) - float64(b.Y)))
	return dx <= 1 && dy <= 1
}
func (a Point) followPoint(b Point) Point {
	dx := b.X - a.X
	dy := b.Y - a.Y
	if dx != 0 {
		dx /= int(math.Abs(float64(dx)))
	}
	if dy != 0 {
		dy /= int(math.Abs(float64(dy)))
	}
	return Point{a.X + dx, a.Y + dy}
}
func (r *Rope) Update(move Move, path map[Point]struct{}) {
	for range move.steps {
		for i := range *r {
			if i == 0 {
				(*r)[i].update(move.direction)
			} else {
				if !(*r)[i].isAdjacent((*r)[i-1]) {
					(*r)[i] = (*r)[i].followPoint((*r)[i-1])
				}
			}
		}
		path[(*r).getTail()] = struct{}{}
	}
}
func (r Rope) getTail() Point {
	l := len(r)
	return r[l-1]
}
