package game

import (
	"fmt"
	"net/http"
)

func NewGame(w http.ResponseWriter, r *http.Request) {
	result := newGame(r.URL.Query().Get("gameID"))
	fmt.Fprint(w, result)

}

func GuessOnline(w http.ResponseWriter, r *http.Request) {
	guessString := r.URL.Query().Get("userguess")
	if guessString == "" {
		http.Error(w, "Missing <userguess> parameter from GET request", http.StatusBadRequest)
		return
	}

	gameInstance := r.PathValue("gameID")
	if gameInstance == "" {
		http.Error(w, "Missing <userID> parameter from url path", http.StatusBadRequest)
		return
	}

	result, err := guessOnline(guessString, gameInstance)
	if err != nil {
		http.Error(w, result, http.StatusBadRequest)
	}

	fmt.Fprint(w, result)
}
