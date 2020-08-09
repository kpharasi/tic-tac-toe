package main

import "fmt"

// Player refers to the user going to play the game
type Player struct {
	name string
	mark string
}

// NewPlayer is the constructor for creating a new player
func NewPlayer(name string, mark string) *Player {
	return &Player{
		name: name,
		mark: mark,
	}
}

// MakeMove is a function called to make the next move by a player
func (p *Player) MakeMove(board [][]string) error {
	fmt.Println("Enter coordinates of your next move")
	var xCrd, yCrd int
	if n, err := fmt.Scanf("%d %d", &xCrd, &yCrd); err != nil || n != 2 {
		return fmt.Errorf("There is an error reading coordinates")
	}

	board[xCrd][yCrd] = p.mark
	return nil
}

// GetMark gets the mark allocated to the current player
func (p *Player) GetMark() string {
	return p.mark
}

// GetName gets the name of the current player
func (p *Player) GetName() string {
	return p.name
}
