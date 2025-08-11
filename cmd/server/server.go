package main

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"guessgame/pkg/game"

	"github.com/julienschmidt/httprouter"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func initLogger() {
	zerolog.TimeFieldFormat = time.RFC3339
	logfile, err := os.OpenFile("server.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		panic(fmt.Sprintf("Failed to open log file: %v", err))
	}
	consoleWriter := zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: time.RFC3339}
	multi := zerolog.MultiLevelWriter(consoleWriter, logfile)

	log.Logger = zerolog.New(multi).With().Timestamp().Logger()
}

//TODO: Move path variable here and pass it on to NewGame and GuessOnline

func main() {
	initLogger()
	router := httprouter.New()
	router.GET("/api/newgame", game.NewGame)
	router.GET("/api/guess/:gameID/", game.GuessOnline)

	log.Info().Msg("Server starting on port :8080")
	err := http.ListenAndServe(":8080", router)
	if err != nil {
		log.Fatal().Err(err).Msg("Server failed to start")
	}

}
