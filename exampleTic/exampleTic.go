package main

import (
	"fmt"
)

type TicTacToe struct {
	board [][]string
	turn  string
}

func (t *TicTacToe) initBoard() {
	for i, _ := range t.board {
		t.board[i] = make([]string, 3)
		for j, _ := range t.board[i] {
			t.board[i][j] = "-"
		}
	}
}

func (t *TicTacToe) logTurn() {
	fmt.Println("It is player " + t.turn + "'s turn.\n")
}

func (t *TicTacToe) logBoard() {
	fmt.Println("The board: \n")

	fmt.Println("   0 1 2")

	for i, _ := range t.board {
		fmt.Println(i, t.board[i])
	}
}

func (t *TicTacToe) markSpot(row int, col int) {
	// If space is claimed, don't override
	if t.board[row][col] != "-" {
		return
	}

	t.board[row][col] = t.turn

	t.logBoard()

	// If winner, log it.
	if t.isWinner() {
		fmt.Println(t.turn + " wins!")
		return
	}

	// Otherwise, continue to next turn
	if t.turn == "X" {
		t.turn = "O"
	} else {
		t.turn = "X"
	}
	t.logTurn()
}

func (t *TicTacToe) promptMove() {
	var row, col int
	fmt.Println("Which row would you like to mark?")
	fmt.Scanf("%d", &row)
	fmt.Println("Which column would you like to mark?")
	fmt.Scanf("%d", &col)

	if col >= 0 && col <= 2 && row >= 0 && row <= 2 {
		t.markSpot(row, col)
	}
}

func (t *TicTacToe) isWinner() bool {
	return (t.board[0][0] == t.turn && t.board[0][1] == t.turn && t.board[0][2] == t.turn) ||
		(t.board[1][0] == t.turn && t.board[1][1] == t.turn && t.board[1][2] == t.turn) ||
		(t.board[2][0] == t.turn && t.board[2][1] == t.turn && t.board[2][2] == t.turn) ||
		(t.board[0][0] == t.turn && t.board[1][0] == t.turn && t.board[2][0] == t.turn) ||
		(t.board[0][1] == t.turn && t.board[1][1] == t.turn && t.board[2][1] == t.turn) ||
		(t.board[0][2] == t.turn && t.board[1][2] == t.turn && t.board[2][2] == t.turn) ||
		(t.board[0][0] == t.turn && t.board[1][1] == t.turn && t.board[2][2] == t.turn) ||
		(t.board[0][2] == t.turn && t.board[1][1] == t.turn && t.board[2][0] == t.turn)
}

func main() {
	gameBoard := TicTacToe{make([][]string, 3), "X"}
	gameBoard.initBoard()
	fmt.Println("Welcome to Tic-Tac-Toe! This game uses a zero-index. 3 columns and three rows\n",
		"(0, 0)(0, 1)(0, 2)\n",
		"(1, 0)(1, 1)(1, 2)\n",
		"(2, 0)(2, 1)(2, 2)\n")
	gameBoard.logBoard()

	// There can only be 9 turns in Tic-Tac-Toe
	for i := 0; i < 9; i++ {
		if gameBoard.isWinner() {
			break
		}
		gameBoard.promptMove()
	}
}
