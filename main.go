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

// Create a new instance and save it to json
func newGame(w http.ResponseWriter, r *http.Request) {
	id := rand.Intn(10000)
	targetNumber := rand.Intn(100)
	currentGame := Game{id, targetNumber, 10, "ONGOING"}

	saveToFile(&currentGame)
	fmt.Fprintf(w, fmt.Sprintf("New game instance created. Use id=%d", currentGame.Id))
}

// Save game instance to file as JSON
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
	// **Validation
	guessString := r.URL.Query().Get("userguess")
	gameInstance := r.URL.Query().Get("userID")
	if guessString == "" {
		http.Error(w, "Missing <userguess> parameter from GET request", http.StatusBadRequest)
		return
	}
	if gameInstance == "" {
		http.Error(w, "Missing <userID> parameter from GET request", http.StatusBadRequest)
		return
	}
	fmt.Fprintf(w, "DEBUG: guessString ="+guessString+"\n")

	guess, err := strconv.Atoi(guessString)
	if err != nil || guess < 1 || guess > 100 {
		http.Error(w, "Invalid guess: must be a number between 1 and 100", http.StatusBadRequest)
		return
	}
	gameID, err := strconv.Atoi(gameInstance)
	//**

	fmt.Fprintf(w, "userguess="+guessString)

	game := loadFromFile(strconv.Itoa(gameID))
	result := checkUserGuess(guess, game)

	switch result {
	case "WON":
		{
			game.Guess_counter--
			game.Game_state = "WON"
			saveToFile(game)
			fmt.Fprintf(w, "Congratulations! You've WON!!!")
		}
	case "LOWER":
		{
			game.Guess_counter--
			saveToFile(game)
			fmt.Fprintf(w, "Almost there...try Higher...")
		}
	case "HIGHER":
		{
			game.Guess_counter--
			saveToFile(game)
			fmt.Fprintf(w, "Almost there...try Lower...")
		}
	case "LOST":
		{
			game.Guess_counter--
			game.Game_state = "LOST"
			saveToFile(game)
			fmt.Fprintf(w, "Sorry, you've LOST!")

		}

	}

}

func main() {
	http.HandleFunc("/newgame", newGame)
	http.HandleFunc("/guess", guessOnline)
	http.ListenAndServe(":8080", nil)
}
