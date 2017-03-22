package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type GameState struct {
	Field         [9]string `json:"field"`
	CurrentPlayer string    `json:"currentPlayer"`
}

var gameState = GameState{[9]string{"", "", "", "", "", "", "", "", ""}, "O"}

func main() {
	i := 0
	for gameOngoing() {

		if gameState.CurrentPlayer == "X" {
			gameState.CurrentPlayer = "O"
		} else {
			gameState.CurrentPlayer = "X"
		}

		gameState.playerPut(i)

		i++
		fmt.Println(gameState.Field)

	}

	http.HandleFunc("/decode", func(w http.ResponseWriter, r *http.Request) {
		var game GameState
		json.NewDecoder(r.Body).Decode(&game)

		fmt.Fprintf(w, "Game Result: %s Winner is %s", game.Field, game.CurrentPlayer)
	})

	http.HandleFunc("/encode", func(w http.ResponseWriter, r *http.Request) {
		result := GameState{
			Field:         gameState.Field,
			CurrentPlayer: gameState.CurrentPlayer,
		}

		json.NewEncoder(w).Encode(&result)
	})

	http.ListenAndServe(":8080", nil)
}
func gameOngoing() bool {

	if !playerWon() {
		return !draw()
	}
	return false
}

func (gs *GameState) playerPut(field int) {

	gs.Field[field] = gs.CurrentPlayer

}

func compare(tocompare ...string) bool {
	res := true
	current := ""
	if len(tocompare) > 0 {
		current = tocompare[0]
	} else {
		return false
	}
	for _, v := range tocompare {
		if len(v) == 0 {
			res = false
			break
		}

		if v != current {
			res = false
			break
		}

	}
	return res

}

func playerWon() bool {

	res := false

	for i := 0; i < 3; i++ {

		res = compare(gameState.Field[i*3], gameState.Field[i*3+1], gameState.Field[i*3+2])

		res = compare(gameState.Field[i], gameState.Field[i+3], gameState.Field[i+6])

	}

	res = compare(gameState.Field[0], gameState.Field[4], gameState.Field[8])

	res = compare(gameState.Field[2], gameState.Field[4], gameState.Field[6])

	if res {
		fmt.Printf("Player %s has won !", gameState.CurrentPlayer)
	}

	return res

}

func draw() bool {
	i := 0
	for _, v := range gameState.Field {
		if v == "X" || v == "Y" {
			i++
		}
	}
	if i == 9 {

		fmt.Println("Draw!")

		return true
	}
	return false

}
