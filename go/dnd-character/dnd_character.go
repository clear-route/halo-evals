package dndcharacter

import (
	"math"
	"math/rand"
	"sort"
	"time"
)

// Character represents a D&D character with abilities and hitpoints.
type Character struct {
	Strength     int
	Dexterity    int
	Constitution int
	Intelligence int
	Wisdom       int
	Charisma     int
	Hitpoints    int
}

// Seed the random number generator
func init() {
	rand.Seed(time.Now().UnixNano())
}

// Modifier calculates the ability modifier for a given ability score.
// The modifier is calculated by subtracting 10 from the score, dividing by 2, and rounding down.
func Modifier(score int) int {
	return int(math.Floor(float64(score-10) / 2.0))
}

// rollDie simulates rolling a single d6 die.
func rollDie() int {
	return rand.Intn(6) + 1 // Generates a random integer between 0 and 5, then adds 1
}

// Ability uses randomness to generate the score for an ability.
// This is done by rolling four 6-sided dice and summing the largest three dice.
func Ability() int {
	rolls := make([]int, 4)
	for i := 0; i < 4; i++ {
		rolls[i] = rollDie()
	}
	sort.Ints(rolls) // Sorts in ascending order
	// Sum the three largest dice (rolls[1], rolls[2], rolls[3] after sorting)
	return rolls[1] + rolls[2] + rolls[3]
}

// GenerateCharacter creates a new Character with random scores for abilities and calculated hitpoints.
func GenerateCharacter() Character {
	constitution := Ability()
	return Character{
		Strength:     Ability(),
		Dexterity:    Ability(),
		Constitution: constitution,
		Intelligence: Ability(),
		Wisdom:       Ability(),
		Charisma:     Ability(),
		Hitpoints:    10 + Modifier(constitution),
	}
}
