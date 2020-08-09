package main

import (
	"fmt"
)

func main() {
	board := make([][]string, 3)

	for i := 0; i < 3; i++ {
		board[i] = make([]string, 3)
	}

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			board[i][j] = "."
		}
	}

	count := 0
	playerOne := Player{name: "Player1", mark: "O"}
	playerTwo := Player{name: "Player2", mark: "X"}
	var currPlayer Player
	currPlayer = playerOne

	for count < 9 {
		fmt.Printf("%s's Chance\n", currPlayer.GetName())
		err := currPlayer.MakeMove(board)
		if err != nil {
			fmt.Printf("There was an issue processing the input")
		}

		printBoard(board)

		if winCheck(board, currPlayer.GetMark()) {
			fmt.Printf("%s with mark %s has won the game\n", currPlayer.GetName(), currPlayer.GetMark())
			break
		}

		if currPlayer.GetMark() == "O" {
			currPlayer = playerTwo
		} else {
			currPlayer = playerOne
		}

		count++
	}
	printBoard(board)
}

// printBoard print the current status of the board
func printBoard(board [][]string) {
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			fmt.Print(board[i][j])
		}
		fmt.Printf("\n")
	}
}

// winCheck checks to see if any of the 4 conditions to win
// have been satisfied - Diagonal(*2), horizontal & vertical
func winCheck(board [][]string, playerMark string) bool {
	if checkCol(board, playerMark) || checkRow(board, playerMark) || checkDiag(board, playerMark) {
		return true
	}
	return false
}

func checkCol(board [][]string, playerMark string) bool {
	for i := 0; i < 3; i++ {
		if board[0][i] == playerMark && board[1][i] == playerMark && board[2][i] == playerMark {
			return true
		}
	}
	return false
}

func checkRow(board [][]string, playerMark string) bool {
	for i := 0; i < 3; i++ {
		if board[i][0] == playerMark && board[i][1] == playerMark && board[i][2] == playerMark {
			return true
		}
	}
	return false
}

func checkDiag(board [][]string, playerMark string) bool {
	if (board[0][0] == playerMark && board[1][1] == playerMark && board[2][2] == playerMark) ||
		(board[0][2] == playerMark && board[1][1] == playerMark && board[2][0] == playerMark) {
		return true
	}
	return false
}
