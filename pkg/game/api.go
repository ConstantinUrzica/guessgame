package game

type DB[T any] interface {
	Save(*T) error
	Load() (*T, error)
}

type Game struct {
	Id            int
	Target_number int
	Guess_counter int
	Game_state    string
}

const (
	GameStateWon     = "WON"
	GameStateLost    = "LOST"
	GameStateOngoing = "ONGOING"
	GameStateHigher  = "HIGHER"
	GameStateLower   = "LOWER"
)
