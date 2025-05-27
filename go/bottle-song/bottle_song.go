package bottlesong

import (
	"fmt"
	"strings" // For strings.ToLower
)

// numToWord maps count to its capitalized word representation for the song.
var numToWord = map[int]string{
	10: "Ten", 9: "Nine", 8: "Eight", 7: "Seven", 6: "Six",
	5:  "Five", 4: "Four", 3: "Three", 2: "Two", 1: "One",
}

// bottleNoun returns "bottle" if n is 1, otherwise "bottles".
func bottleNoun(n int) string {
	if n == 1 {
		return "bottle"
	}
	return "bottles"
}

// Recite generates verses of the "Ten Green Bottles" song.
func Recite(startBottles, takeDown int) []string {
	var result []string

	for i := 0; i < takeDown; i++ {
		currentNum := startBottles - i

		if currentNum < 1 { // Song ends after the 1-bottle verse implies "no more"
			break 
		}

		if len(result) > 0 { // Add separator if not the first verse being added to result
			result = append(result, "")
		}

		numWordCurrent := numToWord[currentNum]
		bottleNounCurrent := bottleNoun(currentNum)
		
		line1 := fmt.Sprintf("%s green %s hanging on the wall,", numWordCurrent, bottleNounCurrent)
		line2 := line1 
		line3 := "And if one green bottle should accidentally fall,"
		
		var line4 string
		nextNum := currentNum - 1
		if nextNum == 0 {
			line4 = "There'll be no green bottles hanging on the wall."
		} else {
			// Get the capitalized word (e.g., "Nine", "One")
			originalNumWord := numToWord[nextNum]
			// Convert it to lowercase for the fourth line (e.g., "nine", "one")
			numWordNextFormatted := strings.ToLower(originalNumWord)
			
			bottleNounNext := bottleNoun(nextNum)
			line4 = fmt.Sprintf("There'll be %s green %s hanging on the wall.", numWordNextFormatted, bottleNounNext)
		}
		
		result = append(result, line1, line2, line3, line4)
	}
	return result
}
