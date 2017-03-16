package main

import (
	"fmt"
	"strings"
)

func main() {
	//Create a tic-tac-toe board
	board := [][]string{
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
	}

	//The players take turns
	board[0][0] = " "
	board[0][1] = " "
	board[0][1] = " "
	board[1][0] = " "
	board[1][1] = " "
	board[1][2] = " "
	board[2][0] = " "
	board[2][1] = " "
	board[2][2] = " "

	for i := 0; i < len(board); i++ {
		fmt.Printf("%s/n", strings.Join(board[i], " "))
	}
}
