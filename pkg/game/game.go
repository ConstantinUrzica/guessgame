package game

import (
	"fmt"
	"math/rand"
	"strconv"

	"guessgame/pkg/db"
)

func newGame(idStr string, dbPath string) string {
	id, err := strconv.Atoi(idStr)
	if err != nil {
		id = rand.Intn(10000)
	}

	targetNumber := rand.Intn(100)
	currentGame := Game{id, targetNumber, 10, GameStateOngoing}
	db := db.NewFileDB[Game](strconv.Itoa(id), dbPath)
	db.Save(&currentGame)

	return fmt.Sprintf("New game instance created. Use id=%d in the link to play yor session (Ex: localhost:8080/api/{id}/guess?userguess={yourguess}", currentGame.Id)
}

func guessOnline(userGuess string, gameInstance string, dbPath string) (string, error) {
	guess, err := strconv.Atoi(userGuess)
	if err != nil || guess < 1 || guess > 100 {
		return "Invalid guess: must be a number between 1 and 100", fmt.Errorf("user error")
	}

	gameID, _ := strconv.Atoi(gameInstance)
	db := db.NewFileDB[Game](strconv.Itoa(gameID), dbPath)

	game, err := db.Load()
	if err != nil {
		panic(err)
	}

	result := checkUserGuess(guess, game)
	switch result {
	case GameStateWon:
		{
			game.Guess_counter = 0
			game.Game_state = GameStateWon
			db.Save(game)
			return "\nCongratulations! You've WON!!!\n", nil
		}
	case GameStateLower:
		{
			game.Guess_counter--
			db.Save(game)
			return "\nAlmost there...try Higher...\n", nil
		}
	case GameStateHigher:
		{
			game.Guess_counter--
			db.Save(game)
			return "\nAlmost there...try Lower...\n", nil
		}
	case GameStateLost:
		{
			game.Guess_counter = 0
			game.Game_state = GameStateLost
			db.Save(game)
			return "\nPlease create a new game and try again\n", nil
		}
	}

	return "", nil
}

// TODO : test unit tests here
func checkUserGuess(guess int, game *Game) string {
	if game.Guess_counter > 1 {
		if guess == game.Target_number {
			return GameStateWon
		} else if guess < game.Target_number {
			return GameStateLower
		} else {
			return GameStateHigher
		}
	} else {
		return GameStateLost
	}
}
