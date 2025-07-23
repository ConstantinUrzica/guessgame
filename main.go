package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"os"
	"strconv"
)

type Game struct {
	Id            int
	Target_number int
	Guess_counter int
	Game_state    string
}

func newGame(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("gameID"))
	if err != nil {
		id = rand.Intn(10000)
	}
	targetNumber := rand.Intn(100)
	currentGame := Game{id, targetNumber, 10, "ONGOING"}

	saveToFile(&currentGame)
	fmt.Fprintf(w, fmt.Sprintf("New game instance created. Use id=%d in the link to play yor session (Ex: localhost:8080/api/{id}/guess?userguess={yourguess}", currentGame.Id))
}

func saveToFile(game *Game) {
	gameData, _ := json.Marshal(game)
	filename := strconv.Itoa(game.Id)
	err := os.WriteFile(filename, gameData, 0644)
	if err != nil {
		fmt.Println("Error saving to file:", err)
	}
	fmt.Println("Operation successfull! I have written the data into a file called ", filename)
}

func loadFromFile(filename string) *Game {
	var loadedGame Game
	filecontent, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println("Error reading file:", err)
	}

	//TODO: add a check here in case the filename is not valid??
	json.Unmarshal(filecontent, &loadedGame)
	return &loadedGame
}

func checkUserGuess(guess int, game *Game) string {
	if game.Guess_counter > 1 {
		if guess == game.Target_number {
			return "WON"
		} else if guess < game.Target_number {
			return "LOWER"
		} else {
			return "HIGHER"
		}
	} else {
		return "LOST"
	}
}

func guessOnline(w http.ResponseWriter, r *http.Request) {
	guessString := r.URL.Query().Get("userguess")
	gameInstance := r.PathValue("gameID")
	if guessString == "" {
		http.Error(w, "Missing <userguess> parameter from GET request", http.StatusBadRequest)
		return
	}
	if gameInstance == "" {
		http.Error(w, "Missing <userID> parameter from url path", http.StatusBadRequest)
		return
	}

	guess, err := strconv.Atoi(guessString)
	if err != nil || guess < 1 || guess > 100 {
		http.Error(w, "Invalid guess: must be a number between 1 and 100", http.StatusBadRequest)
		return
	}

	// TODO: add validation for this
	gameID, _ := strconv.Atoi(gameInstance)

	game := loadFromFile(strconv.Itoa(gameID))
	result := checkUserGuess(guess, game)

	switch result {
	case "WON":
		{
			game.Guess_counter = 0
			game.Game_state = "WON"
			saveToFile(game)
			fmt.Fprintf(w, "\n\nCongratulations! You've WON!!!\n\n")
		}
	case "LOWER":
		{
			game.Guess_counter--
			saveToFile(game)
			fmt.Fprintf(w, "\n\nAlmost there...try Higher...\n\n")
		}
	case "HIGHER":
		{
			game.Guess_counter--
			saveToFile(game)
			fmt.Fprintf(w, "\n\nAlmost there...try Lower...\n\n")
		}
	case "LOST":
		{
			game.Guess_counter = 0
			game.Game_state = "LOST"
			saveToFile(game)
			fmt.Fprintf(w, "\n\nSorry, you've LOST! Please create a new game and try again\n\n")
		}
	}
}

func main() {
	http.HandleFunc("/api/newgame", newGame)
	http.HandleFunc("/api/{gameID}/guess", guessOnline)
	http.ListenAndServe(":8080", nil)
}
