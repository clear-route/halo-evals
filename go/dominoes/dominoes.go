package dominoes

// Domino represents a tile with two numbers.
// It is defined as an array of two integers.
type Domino [2]int

// MakeChain attempts to arrange a given set of dominoes into a valid chain.
// A valid chain is one where adjacent dominoes match (e.g., [a,b][b,c])
// and the ends of the entire chain also match (e.g., first part of first domino
// matches second part of last domino).
// Dominoes can be flipped.
// An empty input is a valid chain. A single domino [a,a] is a valid chain.
func MakeChain(input []Domino) ([]Domino, bool) {
	n := len(input)
	if n == 0 {
		return nil, true
	}

	used := make([]bool, n)
	// currentChain stores the dominoes in their placed orientation during search
	currentChain := make([]Domino, n)

	// Define the recursive helper function using a closure to capture 'input', 'used', 'currentChain', and 'n'.
	var solve func(numPlaced int, lastValue int, targetStartValue int) bool
	solve = func(numPlaced int, lastValue int, targetStartValue int) bool {
		if numPlaced == n { // All dominoes have been placed
			// Check if the chain closes by matching the end of the last domino
			// with the beginning of the first domino.
			return lastValue == targetStartValue
		}

		for i := 0; i < n; i++ { // Iterate through all original dominoes
			if !used[i] { // If this domino hasn't been used yet
				dominoToTry := input[i]
				used[i] = true // Mark as used for this path

				// Try original orientation: [a,b] -> check if a == lastValue
				if dominoToTry[0] == lastValue {
					currentChain[numPlaced] = dominoToTry
					if solve(numPlaced+1, dominoToTry[1], targetStartValue) {
						return true // Solution found
					}
				}

				// Try flipped orientation: [b,a] -> check if b == lastValue
				// Only try flipping if it's not a double (e.g., [2,2]) to avoid redundant checks,
				// and if the first attempt with original orientation didn't already match.
				if dominoToTry[0] != dominoToTry[1] { // Not a double like [x,x]
					if dominoToTry[1] == lastValue {
						currentChain[numPlaced] = Domino{dominoToTry[1], dominoToTry[0]}
						if solve(numPlaced+1, dominoToTry[0], targetStartValue) {
							return true // Solution found
						}
					}
				}
				used[i] = false // Backtrack: unmark as used
			}
		}
		return false // No solution found down this path
	}

	// Try starting the chain with each domino from the input
	for i := 0; i < n; i++ {
		used[i] = true // Mark the first domino as used

		// Attempt 1: Start with input[i] in its original orientation [a,b]
		// Chain starts with 'a', next domino must start with 'b'. Target is 'a'.
		currentChain[0] = input[i]
		if solve(1, input[i][1], input[i][0]) {
			// Solution found, create a copy of currentChain to return
			solution := make([]Domino, n)
			copy(solution, currentChain)
			return solution, true
		}

		// Attempt 2: Start with input[i] flipped [b,a] (if not a double)
		// Chain starts with 'b', next domino must start with 'a'. Target is 'b'.
		if input[i][0] != input[i][1] { // If it's not a double like [x,x]
			currentChain[0] = Domino{input[i][1], input[i][0]}
			if solve(1, input[i][0], input[i][1]) {
				solution := make([]Domino, n)
				copy(solution, currentChain)
				return solution, true
			}
		}
		used[i] = false // Backtrack: this starting domino didn't lead to a full solution
	}

	return nil, false // No solution found after trying all starting configurations
}
