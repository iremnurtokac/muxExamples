package main

import (
	"fmt"
	"strings"
)

func main() {
	// Create a tic-tac-toe board.
	board := [][]string{
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
	}

	// The players take turns.
	board[0][0] = "_"
	board[0][1] = "_"
	board[0][2] = "_"
	board[1][0] = "_"
	board[1][1] = "X"
	board[1][2] = "_"
	board[2][0] = "_"
	board[2][1] = "O"
	board[2][2] = "X"

	for i := 0; i < len(board); i++ {
		fmt.Printf("%s\n", strings.Join(board[i], " "))
	}
}
