package bowling

// NewGame creates a new, empty bowling game.
func NewGame() *Game {
	return &Game{}
}

// checkCompletionAndValidity analyzes the current rolls to determine game state.
// It returns:
// - score: The calculated score if the game is validly structured up to that point.
// - isComplete: True if the current rolls form a perfectly complete 10-frame game.
// - err: An error if rolls are invalid (ErrNegativeRoll, ErrTooManyPins) or game is definitively incomplete for scoring (ErrGameNotOver).
func (g *Game) checkCompletionAndValidity() (score int, isComplete bool, err error) {
	currentScore := 0
	rollIdx := 0 

	for frame := 0; frame < 10; frame++ {
		if rollIdx >= g.numRolls {
			return 0, false, ErrGameNotOver 
		}

		firstRollInFrame := g.rolls[rollIdx]
		if firstRollInFrame < 0 { return 0, false, ErrNegativeRoll }
		if firstRollInFrame > 10 { return 0, false, ErrTooManyPins } 

		if firstRollInFrame == 10 { // Strike
			if rollIdx+2 >= g.numRolls { return 0, false, ErrGameNotOver } 
			
			bonus1 := g.rolls[rollIdx+1]
			bonus2 := g.rolls[rollIdx+2]
			if bonus1 < 0 || bonus2 < 0 { return 0, false, ErrNegativeRoll }
			if bonus1 > 10 || bonus2 > 10 { return 0, false, ErrTooManyPins }

			// Special 10th frame bonus roll validation for non-strike followed by sum > 10
			if frame == 9 && bonus1 < 10 && (bonus1+bonus2 > 10) {
				return 0, false, ErrTooManyPins
			}
			
			currentScore += 10 + bonus1 + bonus2
			
			if frame == 9 { // 10th frame strike uses all 3 rolls for its structure
				rollIdx += 3
			} else { // Strike in frames 1-9 uses 1 roll for its structure
				rollIdx++
			}
		} else { // Not a strike on first ball
			if rollIdx+1 >= g.numRolls { return 0, false, ErrGameNotOver } 
			
			secondRollInFrame := g.rolls[rollIdx+1]
			if secondRollInFrame < 0 { return 0, false, ErrNegativeRoll }
			if secondRollInFrame > 10 { return 0, false, ErrTooManyPins }
			if firstRollInFrame+secondRollInFrame > 10 { return 0, false, ErrTooManyPins } 

			if firstRollInFrame+secondRollInFrame == 10 { // Spare
				if rollIdx+2 >= g.numRolls { return 0, false, ErrGameNotOver } 
				
				bonusSpare := g.rolls[rollIdx+2]
				if bonusSpare < 0 { return 0, false, ErrNegativeRoll }
				if bonusSpare > 10 { return 0, false, ErrTooManyPins }
				
				currentScore += 10 + bonusSpare
				
				if frame == 9 { // 10th frame spare uses all 3 rolls
					rollIdx += 3
				} else { // Spare in frames 1-9 uses 2 rolls
					rollIdx += 2
				}
			} else { // Open frame
				currentScore += firstRollInFrame + secondRollInFrame
				rollIdx += 2 // Open frame always uses 2 rolls
			}
		}
	}

	// After 10 frames are processed and rollIdx is updated:
	if rollIdx == g.numRolls {
		return currentScore, true, nil // Perfectly complete game.
	} else if rollIdx < g.numRolls {
		// Rolls are left over after 10 frames were fully defined by rollIdx.
		// This means too many rolls were provided.
		return 0, false, ErrTooManyPins 
	} else { // rollIdx > g.numRolls
		// This case implies internal logic error or not enough rolls,
		// which should have been caught by ErrGameNotOver inside the loop.
		return 0, false, ErrGameNotOver 
	}
}

// Roll records a new roll in the game.
// It returns an error for invalid pin counts or if the game is already over.
func (g *Game) Roll(pins int) error {
	if pins < 0 { return ErrNegativeRoll }
	if pins > 10 { return ErrTooManyPins } // Single roll value check.

	// 1. Check if game was already complete BEFORE this roll.
	// Create a temporary game state representing the game *before* this current roll.
	tempGameBeforeCurrentRoll := Game{numRolls: g.numRolls}
	for i := 0; i < g.numRolls; i++ {
		tempGameBeforeCurrentRoll.rolls[i] = g.rolls[i]
	}
	_, wasAlreadyComplete, _ := tempGameBeforeCurrentRoll.checkCompletionAndValidity()
	if wasAlreadyComplete {
		return ErrGameOver
	}

	// 2. Check against absolute maximum physical rolls.
	if g.numRolls >= maxRollsInGame {
		return ErrGameOver
	}
	
	// 3. Tentatively add the roll
	originalNumRolls := g.numRolls // Store to allow revert
	g.rolls[g.numRolls] = pins
	g.numRolls++ 

	// 4. Validate the NEW state for immediate frame integrity issues.
	// This scan checks for specific invalid patterns that Roll should catch.
	tempRollIdxScan := 0
	for frameScan := 0; frameScan < 10 && tempRollIdxScan < g.numRolls; frameScan++ {
		firstInFrameScan := g.rolls[tempRollIdxScan]
		
		if frameScan == 9 { // 10th Frame validation logic
			if firstInFrameScan == 10 { // Strike in 10th
				if tempRollIdxScan+1 < g.numRolls { 
					bonus1Scan := g.rolls[tempRollIdxScan+1]
					if bonus1Scan > 10 { g.numRolls = originalNumRolls; return ErrTooManyPins }

					if tempRollIdxScan+2 < g.numRolls { 
						bonus2Scan := g.rolls[tempRollIdxScan+2]
						if bonus2Scan > 10 { g.numRolls = originalNumRolls; return ErrTooManyPins }

						if bonus1Scan < 10 && (bonus1Scan+bonus2Scan > 10) { 
							g.numRolls = originalNumRolls; return ErrTooManyPins
						}
						// Test for: X, non-strike, then strike for bonus (e.g. Roll(10) after X,6)
						// 'pins' is bonus2Scan. originalNumRolls was before 'pins' was added (so current g.numRolls is originalNumRolls+1)
						// This condition means the current 'pins' (which is bonus2Scan) is 10, bonus1Scan was not 10, and 'pins' is the roll just added.
						if bonus1Scan < 10 && bonus2Scan == 10 && (tempRollIdxScan+2 == originalNumRolls) { 
							g.numRolls = originalNumRolls; return ErrTooManyPins
						}
					}
				}
				tempRollIdxScan += 3 // Advance parser for 10th frame strike structure
			} else { // Not a strike on first ball of 10th
				if tempRollIdxScan+1 < g.numRolls { 
					secondInFrameScan := g.rolls[tempRollIdxScan+1]
					if firstInFrameScan+secondInFrameScan > 10 { 
						g.numRolls = originalNumRolls; return ErrTooManyPins
					}
					if firstInFrameScan+secondInFrameScan == 10 { 
						tempRollIdxScan += 3 // Spare in 10th
					} else { 
						tempRollIdxScan += 2 // Open in 10th
					}
				} else { break; } // Not enough rolls yet
			}
		} else { // Frames 1-9 validation logic
			if firstInFrameScan < 10 { // Not a strike
				if tempRollIdxScan+1 < g.numRolls { 
					secondInFrameScan := g.rolls[tempRollIdxScan+1]
					if firstInFrameScan+secondInFrameScan > 10 { 
						g.numRolls = originalNumRolls; return ErrTooManyPins
					}
					tempRollIdxScan += 2
				} else { break; } // Frame incomplete, stop scan
			} else { // Strike in frames 1-9
				tempRollIdxScan++
			}
		}
	}
	
	return nil // Roll is accepted
}

// Score calculates the total score for the game.
// It's called only at the end of the game.
// It returns an error if the game is not over or if rolls are invalid.
func (g *Game) Score() (int, error) {
	score, isComplete, err := g.checkCompletionAndValidity()

	if err != nil {
		// Propagate validation errors like ErrTooManyPins, ErrNegativeRoll.
		return 0, err 
	}

	if !isComplete {
		// If checkCompletionAndValidity didn't find a structural error but
		// concluded the game isn't perfectly 10 frames with all rolls matching,
		// it means the game is not over (e.g. too few rolls for bonuses).
		return 0, ErrGameNotOver
	}

	return score, nil
}
