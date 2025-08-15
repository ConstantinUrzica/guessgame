package game

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/rs/zerolog/log"
)

func NewGameHandler(dbPath string) func(http.ResponseWriter, *http.Request, httprouter.Params) {

	return func(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
		result := newGame(params.ByName("gameID"), dbPath)
		fmt.Fprint(w, result)
		log.Info().
			Str("handler", "NewGame").
			Msg(result)
	}

}

func GuessOnlineHandler(dbPath string) func(http.ResponseWriter, *http.Request, httprouter.Params) {
	return func(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
		guessString := r.URL.Query().Get("userguess")
		if guessString == "" {
			http.Error(w, "Missing <userguess> parameter from GET request", http.StatusBadRequest)
			log.Error().Msg("GuessOnline: missing <userguess> parameter from GET request")

			return
		}

		gameInstance := params.ByName("gameID")
		if gameInstance == "" {
			http.Error(w, "Missing <userID> parameter from url path", http.StatusBadRequest)
			log.Error().Msg("GuessOnline: missing <userID> parameter from url path")
			return
		}

		result, err := guessOnline(guessString, gameInstance, dbPath)
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

}
