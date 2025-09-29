package day17

import "github.com/Evokoo/AOC_2022_Go/tools"

// ========================
// POINT
// ========================
type Point struct{ x, y int }

// ========================
// PILE
// ========================
type Pile struct {
	blocks map[Point]struct{}
	height int
}

func NewPile() Pile {
	return Pile{
		blocks: make(map[Point]struct{}),
		height: 0,
	}
}

func (p *Pile) AddBlock(b Block) {
	for _, segment := range b {
		if segment.y+1 > (*p).height {
			(*p).height = segment.y + 1
		}

		(*p).blocks[segment] = struct{}{}
	}
}

func (p Pile) Has(point Point) bool {
	_, found := p.blocks[point]
	return found
}

// ========================
// BLOCK
// ========================
type Block []Point

func NewBlock(index int, pile *Pile) Block {
	x, y := 2, pile.height+3

	switch index {
	//Horiztonal Line
	case 0:
		return Block{{x, y}, {x + 1, y}, {x + 2, y}, {x + 3, y}}
	//Cross
	case 1:
		return Block{{x, y + 1}, {x + 1, y + 1}, {x + 1, y + 2}, {x + 1, y}, {x + 2, y + 1}}
	//Backwards L
	case 2:
		return Block{{x, y}, {x + 1, y}, {x + 2, y}, {x + 2, y + 1}, {x + 2, y + 2}}
	//Vertical Line
	case 3:
		return Block{{x, y}, {x, y + 1}, {x, y + 2}, {x, y + 3}}
	//2x2
	case 4:
		return Block{{x, y}, {x + 1, y}, {x, y + 1}, {x + 1, y + 1}}
	default:
		panic("Invalid Block Type")
	}
}

func (b *Block) HorizontalShift(dx int, pile *Pile) {
	var nextPosition Block
	for _, segment := range *b {
		n := Point{segment.x + dx, segment.y}
		if n.x < 0 || n.x >= 7 || pile.Has(n) {
			return
		}
		nextPosition = append(nextPosition, n)
	}
	*b = nextPosition
}

func (b *Block) VerticalShift(pile *Pile) bool {
	var nextPosition Block
	for _, segment := range *b {
		n := Point{segment.x, segment.y - 1}
		if (*pile).Has(n) || n.y < 0 {
			(*pile).AddBlock(*b)
			return false
		}
		nextPosition = append(nextPosition, n)
	}
	*b = nextPosition
	return true
}

// ========================
// JETS
// ========================
type Jets []int

func ParseJets(file string) Jets {
	var output Jets

	for _, r := range tools.ReadFile(file) {
		switch r {
		case '<':
			output = append(output, -1)
		case '>':
			output = append(output, 1)
		}
	}

	return output
}
