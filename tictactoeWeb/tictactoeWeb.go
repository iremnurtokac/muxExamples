package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type GameState struct {
	Field         [9]string
	CurrentPlayer string
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

	playMove := map[string][9]string{

		"Tictactoe": gameState.Field,
	}

	/* winner := map[string]string{

		"Tictactoe": gameState.CurrentPlayer,
	} */

	/*	playMove := map[string][9]string{

		"Tictactoe": gameState.Field,
	} */

	/*	playMove := map[string]string{

		"Tictactoe": gameState.Field[4],
	} */

	r := mux.NewRouter()
	r.HandleFunc("/users/{state}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		state := vars["state"]
		current := playMove[state]

		fmt.Fprintf(w, "Game Result: %s", current)
	}).Methods("GET")

	http.ListenAndServe(":8080", r)
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
