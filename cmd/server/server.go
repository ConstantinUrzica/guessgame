package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/rs/zerolog/log"

	"guessgame/internal/logger"
	"guessgame/pkg/game"
)

//TODO: Move path variable here and pass it on to NewGame and GuessOnline

func main() {
	logger.InitLogger()
	router := httprouter.New()
	router.GET("/api/newgame", game.NewGame)
	router.GET("/api/guess/:gameID/", game.GuessOnline)

	log.Info().Msg("Server starting on port :8080")
	err := http.ListenAndServe(":8080", router)
	if err != nil {
		log.Fatal().Err(err).Msg("Server failed to start")
	}

}
