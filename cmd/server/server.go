package main

import (
	"net/http"

	"guessgame/pkg/game"
)

func main() {
	http.HandleFunc("/api/newgame", game.NewGame)
	http.HandleFunc("/api/guess/{gameID}/", game.GuessOnline)
	http.ListenAndServe(":8080", nil)
}
