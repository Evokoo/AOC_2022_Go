package main

import (
	"fmt"
	"strings"

	"github.com/Evokoo/AOC_2022_Go/tools"
)

func main() {
	fmt.Printf("PART ONE :  %d\n", Solve("input.txt", 1))
	fmt.Printf("PART TWO :  %d\n", Solve("input.txt", 2))
}

type Game struct {
	P1, P2 byte
}

func Solve(file string, part int) int {
	score := 0
	for _, game := range parseInput(file) {
		switch part {
		case 1:
			score += playGameOne(game)
		case 2:
			score += playGameTwo(game)
		}
	}
	return score
}

func parseInput(file string) []Game {
	data := strings.Split(tools.ReadFile(file), "\n")
	games := make([]Game, len(data))

	for i, game := range data {
		games[i] = Game{P1: game[0], P2: game[2]}
	}
	return games
}
func playGameOne(game Game) int {
	score := 0
	switch game.P2 {
	case 'X':
		score += 1
	case 'Y':
		score += 2
	case 'Z':
		score += 3
	}

	switch {
	case (game.P1 == 'A' && game.P2 == 'Y') ||
		(game.P1 == 'B' && game.P2 == 'Z') ||
		(game.P1 == 'C' && game.P2 == 'X'):
		score += 6
	case (game.P1 == 'A' && game.P2 == 'X') ||
		(game.P1 == 'B' && game.P2 == 'Y') ||
		(game.P1 == 'C' && game.P2 == 'Z'):
		score += 3
	}
	return score
}
func playGameTwo(game Game) int {
	score := 0
	switch game.P2 {
	case 'X':
		// lose
		switch game.P1 {
		case 'A':
			score += 3
		case 'B':
			score += 1
		case 'C':
			score += 2
		}
	case 'Y':
		// Draw
		score += 3
		switch game.P1 {
		case 'A':
			score += 1
		case 'B':
			score += 2
		case 'C':
			score += 3
		}
	case 'Z':
		// Win
		score += 6
		switch game.P1 {
		case 'A':
			score += 2
		case 'B':
			score += 3
		case 'C':
			score += 1
		}
	}
	return score
}
