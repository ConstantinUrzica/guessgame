package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/rs/zerolog/log"

	"guessgame/observability/logger"
	"guessgame/pkg/game"
)

func main() {
	logger.InitLogger()
	cfgPath := "./cmd/server/config.json"

	cfg, err := Load(cfgPath)
	if err != nil {
		log.Panic().Err(err).Msg("Cannot load config file")
	}
	router := httprouter.New()
	router.GET("/api/newgame", game.NewGameHandler(cfg.DBPath))
	router.GET("/api/guess/:gameID/", game.GuessOnlineHandler(cfg.DBPath))

	log.Info().Msg("Server starting on port :8080")
	err = http.ListenAndServe(":8080", router)
	if err != nil {
		log.Fatal().Err(err).Msg("Server failed to start")
	}

}
