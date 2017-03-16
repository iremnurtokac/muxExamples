package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	userAges := map[string]string{

		"Alice":     "25",
		"Bob":       "30",
		"Claire":    "xx",
		"Tictactoe": "xxx",
	}

	r := mux.NewRouter()
	r.HandleFunc("/users/{name}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		name := vars["name"]
		age := userAges[name]

		fmt.Fprintf(w, "%s", age)
	}).Methods("GET")

	http.ListenAndServe(":8080", r)
}
