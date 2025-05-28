package foodchain

import (
	"strings"
)

type animalInfo struct {
	name     string
	reaction string // The unique second line for this animal's verse, if any.
}

var animals = []animalInfo{
	{name: "fly"}, // Verse 1 (reaction is handled by the common end line)
	{name: "spider", reaction: "It wriggled and jiggled and tickled inside her."},
	{name: "bird", reaction: "How absurd to swallow a bird!"},
	{name: "cat", reaction: "Imagine that, to swallow a cat!"},
	{name: "dog", reaction: "What a hog, to swallow a dog!"},
	{name: "goat", reaction: "Just opened her throat and swallowed a goat!"},
	{name: "cow", reaction: "I don't know how she swallowed a cow!"},
	{name: "horse", reaction: "She's dead, of course!"}, // This reaction replaces the usual ending.
}

const commonEndLine = "I don't know why she swallowed the fly. Perhaps she'll die."

// Verse generates the lyrics for a single verse of the song.
func Verse(v int) string {
	if v < 1 || v > len(animals) {
		return "" // Invalid verse number
	}

	currentIndex := v - 1 // 0-indexed
	currentAnimal := animals[currentIndex]

	var lines []string

	// Line 1: Introduction of the animal.
	lines = append(lines, "I know an old lady who swallowed a "+currentAnimal.name+".")

	// Special case for the horse (last verse).
	if currentAnimal.name == "horse" {
		lines = append(lines, currentAnimal.reaction)
		return strings.Join(lines, "\n")
	}

	// Line 2: Unique reaction for the current animal (if any).
	if currentAnimal.reaction != "" {
		lines = append(lines, currentAnimal.reaction)
	}

	// Cumulative part: from current animal down to fly.
	for i := currentIndex; i > 0; i-- {
		predator := animals[i]
		prey := animals[i-1]
		catchLine := "She swallowed the " + predator.name + " to catch the " + prey.name
		// Special line for catching the spider.
		if prey.name == "spider" {
			catchLine += " that wriggled and jiggled and tickled inside her"
		}
		lines = append(lines, catchLine+".")
	}

	// Common ending line for verses 1-7.
	lines = append(lines, commonEndLine)

	return strings.Join(lines, "\n")
}

// Verses generates the lyrics for a range of verses.
func Verses(start, end int) string {
	var result []string
	for i := start; i <= end; i++ {
		result = append(result, Verse(i))
	}
	return strings.Join(result, "\n\n")
}

// Song generates the full lyrics of the song.
func Song() string {
	return Verses(1, len(animals))
}
