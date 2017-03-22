package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type GameState struct {
	Field  [][]string `json:"field"`
	Winner string     `json:"winner"`
}

var gameField = GameState{[][]string{
	[]string{"X", "X", "X"},
	[]string{"O", "O", "X"},
	[]string{"_", "_", "_"},
}, ""}

func main() {

	fmt.Printf("%s %s %s\n%s %s %s\n%s %s %s\n", gameField.Field[0][0], gameField.Field[0][1], gameField.Field[0][2],
		gameField.Field[1][0], gameField.Field[1][1], gameField.Field[1][2],
		gameField.Field[2][0], gameField.Field[2][1], gameField.Field[2][2])

	if (gameField.Field[0][0] == gameField.Field[0][1] && gameField.Field[0][0] == gameField.Field[0][2]) ||
		(gameField.Field[0][0] == gameField.Field[1][0] && gameField.Field[0][0] == gameField.Field[2][0]) {
		if gameField.Field[0][0] == "X" || gameField.Field[0][0] == "O" {
			fmt.Printf("+++Game Over+++ \nThe winner is Player %s", gameField.Field[0][0])
			gameField.Winner = "Winner is  " + gameField.Field[0][0]
		}

	} else if (gameField.Field[1][0] == gameField.Field[1][1] && gameField.Field[1][0] == gameField.Field[1][2]) ||
		(gameField.Field[0][1] == gameField.Field[1][1] && gameField.Field[0][1] == gameField.Field[2][1]) {
		if gameField.Field[1][1] == "X" || gameField.Field[1][1] == "O" {
			fmt.Printf("+++Game Over+++ \nThe winner is Player %s", gameField.Field[1][1])
			gameField.Winner = "Winner is  " + gameField.Field[1][1]
		}
	} else if (gameField.Field[2][0] == gameField.Field[2][1] && gameField.Field[2][0] == gameField.Field[2][2]) ||
		(gameField.Field[0][0] == gameField.Field[1][1] && gameField.Field[0][0] == gameField.Field[2][2]) {
		if gameField.Field[2][2] == "X" || gameField.Field[2][2] == "O" {
			fmt.Printf("+++Game Over+++ \nThe winner is Player %s", gameField.Field[2][2])
			gameField.Winner = "Winner is  " + gameField.Field[2][2]
		}

	} else if (gameField.Field[0][2] == gameField.Field[1][2] && gameField.Field[0][2] == gameField.Field[2][2]) ||
		(gameField.Field[0][2] == gameField.Field[1][1] && gameField.Field[0][2] == gameField.Field[2][0]) {
		if gameField.Field[0][2] == "X" || gameField.Field[0][2] == "O" {
			fmt.Printf("+++Game Over+++ \nThe winner is Player %s", gameField.Field[2][0])
			gameField.Winner = "Winner is  " + gameField.Field[2][0]
		}

	}

	r := mux.NewRouter()
	r.HandleFunc("/users/{state}", func(w http.ResponseWriter, r *http.Request) {
		/*

		   GET http://localhost:8080/users/Tictactoe HTTP/1.1

		*/

		if currentAsJSON, err := json.Marshal(gameField); err == nil {
			w.Header().Set("Content-Type", "application/json")
			w.Header().Set("Access-Control-Allow-Origin", "*")
			w.WriteHeader(http.StatusOK)
			w.Write(currentAsJSON)
			//fmt.Fprintf(w, "%v", string(currentAsJSON))

		} else {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "Error while creating JSON : %v", err)
		}

	}).Methods("GET")
	http.ListenAndServe(":8080", r)
}
