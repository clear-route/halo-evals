package alphametics

import (
	"errors"
	"regexp"
	"strings"
)

// Solve attempts to find a mapping of letters to digits for an alphametics puzzle.
func Solve(puzzle string) (map[string]int, error) {
	parts := strings.Split(puzzle, " == ")
	if len(parts) != 2 {
		return nil, errors.New("invalid puzzle format: missing ' == '")
	}
	sumStr := parts[0]
	resultWord := strings.TrimSpace(parts[1])

	operandWordsStr := strings.Split(sumStr, " + ")
	operandWords := make([]string, len(operandWordsStr))
	for i, op := range operandWordsStr {
		operandWords[i] = strings.TrimSpace(op)
		if operandWords[i] == "" {
			return nil, errors.New("invalid puzzle format: empty operand")
		}
	}
    if resultWord == "" {
        return nil, errors.New("invalid puzzle format: empty result word")
    }


	allWords := append([]string{}, operandWords...)
	allWords = append(allWords, resultWord)

	uniqueLettersList := []rune{}
	letterSet := make(map[rune]struct{})
	leadingLetters := make(map[rune]bool)

	re := regexp.MustCompile("^[A-Z]+$") // Ensures words contain only uppercase letters
	for _, word := range allWords {
		if !re.MatchString(word) {
			return nil, errors.New("invalid character in puzzle words: " + word)
		}
		for i, r := range word {
			if _, exists := letterSet[r]; !exists {
				letterSet[r] = struct{}{}
				uniqueLettersList = append(uniqueLettersList, r)
			}
			if i == 0 && len(word) > 1 {
				leadingLetters[r] = true
			}
		}
	}
    
    if len(uniqueLettersList) > 10 {
        return nil, errors.New("too many unique letters (more than 10)")
    }

	currentAssignment := make(map[rune]int)
	usedDigits := [10]bool{}

	solutionMap, found := solveRecursive(
		uniqueLettersList, 
		0, 
		currentAssignment, 
		usedDigits, 
		leadingLetters, 
		operandWords, 
		resultWord,
	)

	if !found {
		return nil, errors.New("no solution found")
	}
    
    finalSolutionStringKeys := make(map[string]int)
    for r, digit := range solutionMap {
        finalSolutionStringKeys[string(r)] = digit
    }
	return finalSolutionStringKeys, nil
}

func solveRecursive(
	letters []rune,
	idx int,
	currentAssignment map[rune]int, // This map is mutated and restored during backtracking
	usedDigits [10]bool, // Pass by value creates a copy for each recursive call, handling used state correctly
	leadingLetters map[rune]bool,
	operandWords []string,
	resultWord string,
) (map[rune]int, bool) {
	if idx == len(letters) { // All letters have been assigned a digit
		// Check leading zero constraint for all words explicitly based on current full assignment
		// This check is also present and critical in wordToNum
		for _, word := range operandWords {
			if len(word) > 1 && currentAssignment[rune(word[0])] == 0 {
				return nil, false
			}
		}
		if len(resultWord) > 1 && currentAssignment[rune(resultWord[0])] == 0 {
			return nil, false
		}

		sumOperandsVal := 0
		for _, word := range operandWords {
			val, err := wordToNum(word, currentAssignment) // wordToNum also checks leading zeros
			if err != nil { 
				return nil, false
			}
			sumOperandsVal += val
		}

		valResult, err := wordToNum(resultWord, currentAssignment)
		if err != nil { 
			return nil, false
		}

		if sumOperandsVal == valResult {
            solutionCopy := make(map[rune]int)
            for k, v := range currentAssignment {
                solutionCopy[k] = v
            }
			return solutionCopy, true
		}
		return nil, false
	}

	currentLetterRune := letters[idx]
	
	for digitToTry := 0; digitToTry <= 9; digitToTry++ {
		if usedDigits[digitToTry] {
			continue
		}

		if leadingLetters[currentLetterRune] && digitToTry == 0 {
			continue
		}

		currentAssignment[currentLetterRune] = digitToTry
		// Create a new copy of usedDigits for the next recursive call because arrays are pass-by-value.
		// This correctly isolates the used state for different branches of the recursion.
		newUsedDigitsState := usedDigits
		newUsedDigitsState[digitToTry] = true

		if solution, found := solveRecursive(letters, idx+1, currentAssignment, newUsedDigitsState, leadingLetters, operandWords, resultWord); found {
			return solution, true
		}
	}
    // Backtrack: remove the assignment for currentLetterRune as this path didn't lead to a solution
    delete(currentAssignment, currentLetterRune)
	return nil, false
}

func wordToNum(word string, assignment map[rune]int) (int, error) {
	if len(word) == 0 {
		return 0, errors.New("cannot convert empty word to number")
	}
	
	firstLetterRune := rune(word[0])
	assignedDigitForFirstLetter, assigned := assignment[firstLetterRune]
	if !assigned {
		return 0, errors.New("logic error: letter " + string(firstLetterRune) + " not in assignment during eval")
	}

	if assignedDigitForFirstLetter == 0 && len(word) > 1 {
		return 0, errors.New("leading zero in multi-digit number: " + word)
	}

	num := 0
	for _, letterRune := range word {
		digit, ok := assignment[letterRune]
		if !ok {
			return 0, errors.New("logic error: letter " + string(letterRune) + " not in assignment")
		}
		num = num*10 + digit
	}
	return num, nil
}
