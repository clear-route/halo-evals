package hexadecimal

import (
	"fmt"
	"math"    // For MaxInt64 in ParseHex
	"strings" // Added for ToLower and Contains in HandleErrors
)

// parseHexChar converts a single hexadecimal character (rune) to its integer value (0-15).
// It returns an error if the character is not a valid hex digit.
func parseHexChar(r rune) (int, error) {
	switch {
	case r >= '0' && r <= '9':
		return int(r - '0'), nil
	case r >= 'a' && r <= 'f':
		return int(r-'a') + 10, nil
	case r >= 'A' && r <= 'F':
		return int(r-'A') + 10, nil
	default:
		return 0, fmt.Errorf("syntax error: invalid hexadecimal character '%c'", r)
	}
}

// ParseHex converts a hexadecimal string to its int64 decimal equivalent.
// It adheres to first principles, not using built-in conversion functions.
// Errors are returned for invalid syntax or if the value is out of int64 range.
func ParseHex(hexStr string) (int64, error) {
	if hexStr == "" {
		return 0, fmt.Errorf("syntax error: input string is empty")
	}

	var acc int64 = 0
	const maxValDiv16 = math.MaxInt64 / 16
	const maxValMod16 = math.MaxInt64 % 16

	for _, charRune := range hexStr {
		digitVal, err := parseHexChar(charRune)
		if err != nil {
			return 0, err
		}

		if acc > maxValDiv16 || (acc == maxValDiv16 && int64(digitVal) > maxValMod16) {
			return 0, fmt.Errorf("range error: hex value is out of int64 range")
		}

		acc = acc*16 + int64(digitVal)
	}
	return acc, nil
}

// HandleErrors takes a list of hexadecimal strings, calls ParseHex on each,
// and returns a list of error case strings: "none", "syntax", or "range".
func HandleErrors(inputs []string) []string {
	results := make([]string, len(inputs))
	for i, inputStr := range inputs {
		_, err := ParseHex(inputStr)
		if err == nil {
			results[i] = "none"
		} else {
			errMsg := strings.ToLower(err.Error())
			if strings.Contains(errMsg, "syntax") {
				results[i] = "syntax"
			} else if strings.Contains(errMsg, "range") {
				results[i] = "range"
			} else {
				results[i] = "unknown_error"
			}
		}
	}
	return results
}
