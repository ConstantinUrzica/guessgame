package game

import "testing"

func TestCheckUserGuess(t *testing.T) {
	tests := []struct {
		name     string
		guess    int
		game     Game
		expected string
	}{
		{
			name:     "correct guess, guesses remaining",
			guess:    42,
			game:     Game{Target_number: 42, Guess_counter: 3},
			expected: GameStateWon,
		},
		{
			name:     "guess lower than target",
			guess:    30,
			game:     Game{Target_number: 42, Guess_counter: 3},
			expected: GameStateLower,
		},
		{
			name:     "guess higher than target",
			guess:    50,
			game:     Game{Target_number: 42, Guess_counter: 3},
			expected: GameStateHigher,
		},
		{
			name:     "no guesses left (loss)",
			guess:    42, // even if correct
			game:     Game{Target_number: 42, Guess_counter: 1},
			expected: GameStateLost,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := checkUserGuess(tt.guess, &tt.game)
			if result != tt.expected {
				t.Errorf("got %q, want %q", result, tt.expected)
			}
		})
	}
}
