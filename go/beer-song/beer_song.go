package beer

import (
	"errors"
	"fmt"
	"strings"
)

// Verse returns the nth verse of the beer song.
// It returns an error if n is outside the valid range (0-99).
func Verse(n int) (string, error) {
	if n < 0 || n > 99 {
		return "", errors.New("verse number must be between 0 and 99")
	}

	switch n {
	case 0:
		return "No more bottles of beer on the wall, no more bottles of beer.\n" +
			"Go to the store and buy some more, 99 bottles of beer on the wall.\n", nil
	case 1:
		return "1 bottle of beer on the wall, 1 bottle of beer.\n" +
			"Take it down and pass it around, no more bottles of beer on the wall.\n", nil
	case 2:
		return "2 bottles of beer on the wall, 2 bottles of beer.\n" +
			"Take one down and pass it around, 1 bottle of beer on the wall.\n", nil
	default: // n > 2 and n <= 99
		return fmt.Sprintf("%d bottles of beer on the wall, %d bottles of beer.\n"+
			"Take one down and pass it around, %d bottles of beer on the wall.\n", n, n, n-1), nil
	}
}

// Verses returns a string containing all verses from start down to stop, inclusive.
// Verses are separated by a blank line. The entire result will also end with a newline if verses were generated.
// It returns an error if start < stop or if start/stop are outside the valid range (0-99).
func Verses(start, stop int) (string, error) {
	if start < stop {
		return "", errors.New("start verse must be greater than or equal to stop verse")
	}
	if start > 99 || start < 0 || stop > 99 || stop < 0 {
		return "", errors.New("verse numbers must be between 0 and 99")
	}

	var verseParts []string
	for i := start; i >= stop; i-- {
		v, err := Verse(i)
		if err != nil {
			return "", fmt.Errorf("error generating verse %d: %w", i, err)
		}
		verseParts = append(verseParts, v)
	}

	if len(verseParts) == 0 {
		return "", nil // Should not occur if validations pass and stop <= start
	}
    
    // Join verses with a single newline separator. Since each verse already ends with one newline,
    // this effectively creates a blank line between verses.
	joinedString := strings.Join(verseParts, "\n")
    
    // The tests expect the entire block of joined verses to end with an additional newline (making it \n\n at the very end).
	return joinedString + "\n", nil
}

// Song returns the entire beer song, from 99 bottles down to 0.
func Song() string {
	s, err := Verses(99, 0)
	if err != nil {
		panic(fmt.Sprintf("unexpected error generating full song: %v", err))
	}
	return s
}
