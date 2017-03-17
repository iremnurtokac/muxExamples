package main

import "fmt"

func main() {
	gameField := [3][3]string{
		[3]string{"_", "_", "_"},
		[3]string{"_", "_", "_"},
		[3]string{"_", "_", "_"},
	}

	fmt.Printf("%s %s %s\n%s %s %s\n%s %s %s\n", gameField[0][0], gameField[0][1], gameField[0][2],
		gameField[1][0], gameField[1][1], gameField[1][2],
		gameField[2][0], gameField[2][1], gameField[2][2])

	if (gameField[0][0] == gameField[0][1] && gameField[0][0] == gameField[0][2]) ||
		(gameField[0][0] == gameField[1][0] && gameField[0][0] == gameField[2][0]) {
		if gameField[0][0] == "X" || gameField[0][0] == "O" {
			fmt.Printf("+++Game Over+++ \nThe winner is Player %s", gameField[0][0])
		}

	} else if (gameField[1][0] == gameField[1][1] && gameField[1][0] == gameField[1][2]) ||
		(gameField[0][1] == gameField[1][1] && gameField[0][1] == gameField[2][1]) {
		if gameField[1][1] == "X" || gameField[1][1] == "O" {
			fmt.Printf("+++Game Over+++ \nThe winner is Player %s", gameField[1][1])
		}
	} else if (gameField[2][0] == gameField[2][1] && gameField[2][0] == gameField[2][2]) ||
		(gameField[0][0] == gameField[1][1] && gameField[0][0] == gameField[2][2]) {
		if gameField[2][2] == "X" || gameField[2][2] == "O" {
			fmt.Printf("+++Game Over+++ \nThe winner is Player %s", gameField[2][2])
		}

	} else if (gameField[0][2] == gameField[1][2] && gameField[0][2] == gameField[2][2]) ||
		(gameField[0][2] == gameField[1][1] && gameField[0][2] == gameField[2][0]) {
		if gameField[0][2] == "X" || gameField[0][2] == "O" {
			fmt.Printf("+++Game Over+++ \nThe winner is Player %s", gameField[2][0])
		}
	}
}
