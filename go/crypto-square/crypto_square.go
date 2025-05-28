package cryptosquare

import (
	"math"
	"strings"
	"unicode"
)

// normalize cleans the input string: removes spaces/punctuation, converts to lowercase.
func normalize(pt string) string {
	var sb strings.Builder
	for _, r := range pt {
		if unicode.IsLetter(r) || unicode.IsDigit(r) {
			sb.WriteRune(unicode.ToLower(r))
		}
	}
	return sb.String()
}

// calculateDimensions determines the number of columns (c) and rows (r)
// for the crypto square.
func calculateDimensions(length int) (c int, r int) {
	if length == 0 {
		return 0, 0
	}

	c = int(math.Ceil(math.Sqrt(float64(length))))
	// Ensure conditions c >= r and c-r <= 1 are met
	// r is determined by ceil(length / c)
	for {
        if c == 0 { // Ensure c is at least 1 if length > 0
            c = 1
        }
		r = (length + c - 1) / c // Integer ceil(length/c)
		if c >= r && (c-r) <= 1 {
			break
		}
		c++ // Increment columns if conditions not met
	}
	return c, r
}

// Encode encrypts plain text using the crypto square method.
func Encode(pt string) string {
	normalized := normalize(pt)
	length := len(normalized)

	if length == 0 {
		return ""
	}

	c, r := calculateDimensions(length)

	// If, after calculation (e.g. for L=1 leading to c=1, r=1), c is 0, it implies nothing to encode.
	// However, calculateDimensions is designed to return c>=1 if length > 0.
	if c == 0 { // Should technically be covered by length == 0 check, but defensive.
	    return ""
	}

	normalizedRunes := []rune(normalized)
	segments := make([]string, c)

	for j := 0; j < c; j++ { // Iterate through columns to build segments
		var colBuilder strings.Builder
		for i := 0; i < r; i++ { // Iterate through rows for current column
			// Index in the row-major flattened conceptual rectangle
			idx := i*c + j
			if idx < length {
				colBuilder.WriteRune(normalizedRunes[idx])
			} else {
				// Pad with space if current cell is beyond normalized text length
				colBuilder.WriteRune(' ')
			}
		}
		segments[j] = colBuilder.String()
	}

	return strings.Join(segments, " ")
}
