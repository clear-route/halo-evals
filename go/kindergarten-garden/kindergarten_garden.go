package kindergarten

import (
	"fmt"
	"sort"
	"strings"
)

var _ = fmt.Sprintf("") // Ensure fmt is recognized as used

var plantNames = map[rune]string{
	'V': "violets",
	'R': "radishes",
	'C': "clover",
	'G': "grass",
}

// Garden holds the mapping of children to their assigned plants.
type Garden struct {
	assignments map[string][]string
}

// NewGarden parses the diagram and children list to create a Garden.
func NewGarden(diagram string, children []string) (*Garden, error) {
	if !strings.HasPrefix(diagram, "\n") {
		return nil, fmt.Errorf("diagram format error: must start with a newline")
	}

	trimmedDiagram := strings.TrimPrefix(diagram, "\n")
	rows := strings.SplitN(trimmedDiagram, "\n", 3)

	if len(rows) != 2 {
		return nil, fmt.Errorf("diagram format error: must contain exactly two rows")
	}

	row1 := rows[0]
	row2 := rows[1]

	if len(row1) != len(row2) {
		return nil, fmt.Errorf("diagram error: rows have mismatched lengths")
	}

	if len(row1)%2 != 0 { 
		return nil, fmt.Errorf("diagram error: odd number of cups")
	}

	allPlantChars := row1 + row2
	for _, plantRune := range allPlantChars {
		if _, isValidCode := plantNames[plantRune]; !isValidCode {
			return nil, fmt.Errorf("diagram error: invalid plant code '%c'", plantRune)
		}
	}

	seenNames := make(map[string]bool)
	for _, name := range children {
		if seenNames[name] {
			return nil, fmt.Errorf("children list error: duplicate name found")
		}
		seenNames[name] = true
	}

	sortedChildren := make([]string, len(children))
	copy(sortedChildren, children)
	sort.Strings(sortedChildren)

	if len(row1) != len(sortedChildren)*2 {
		return nil, fmt.Errorf("diagram error: number of cups in row (%d) is inconsistent with number of children (%d, expecting %d cups for 2 cups each)", len(row1), len(sortedChildren), len(sortedChildren)*2)
	}

	g := &Garden{
		assignments: make(map[string][]string),
	}

	for i, childName := range sortedChildren {
		cup1Idx := i * 2
		cup2Idx := i*2 + 1
		
		plantCode1 := rune(row1[cup1Idx])
		plantCode2 := rune(row1[cup2Idx])
		plantCode3 := rune(row2[cup1Idx])
		plantCode4 := rune(row2[cup2Idx])

		g.assignments[childName] = []string{
			plantNames[plantCode1],
			plantNames[plantCode2],
			plantNames[plantCode3],
			plantNames[plantCode4],
		}
	}
	return g, nil
}

func (g *Garden) Plants(child string) ([]string, bool) {
	assignedPlants, ok := g.assignments[child]
	return assignedPlants, ok
}

