package bowling

import "errors"

// Predefined errors for the bowling game.
var (
	ErrNegativeRoll    = errors.New("negative roll is invalid")
	ErrTooManyPins     = errors.New("cannot knock down more than 10 pins in a roll or frame")
	ErrGameNotOver     = errors.New("game is not over")
	ErrGameOver        = errors.New("game is over")
)

const maxRollsInGame = 21 // Maximum physical rolls possible in a game.

// Game represents a bowling game.
type Game struct {
	rolls    [maxRollsInGame]int
	numRolls int // Number of rolls made so far.
}
