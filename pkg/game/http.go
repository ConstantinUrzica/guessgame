package game

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/rs/zerolog/log"
)

/*
func NewGameHandler(game Game) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		result := game.New(r.URL.Query().Get("gameID")) // game.New()
		fmt.Fprint(w, result)
	}
}
*/

func NewGame(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	result := newGame(ps.ByName("gameID"))
	fmt.Fprint(w, result)
	log.Info().
		Str("handler", "NewGame").
		Msg(result)
}

func GuessOnline(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	guessString := r.URL.Query().Get("userguess")
	if guessString == "" {
		http.Error(w, "Missing <userguess> parameter from GET request", http.StatusBadRequest)
		log.Error().Msg("GuessOnline: missing <userguess> parameter from GET request")

		return
	}

	gameInstance := ps.ByName("gameID")
	if gameInstance == "" {
		http.Error(w, "Missing <userID> parameter from url path", http.StatusBadRequest)
		log.Error().Msg("GuessOnline: missing <userID> parameter from url path")
		return
	}

	result, err := guessOnline(guessString, gameInstance)
	if err != nil {
		http.Error(w, result, http.StatusBadRequest)
		log.Error().Err(err).Str("gameID", gameInstance).Msg(result)
	}

	log.Info().
		Str("handler", "GuessOnline").
		Str("guess", guessString).
		Str("gameID", gameInstance).
		Msg(result)
	fmt.Fprint(w, result)
}
