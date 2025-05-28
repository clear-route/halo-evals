package connect

// Point represents a coordinate on the board
type Point struct {
	row, col int
}

// ResultOf determines the winner of the Hex game.
// Player 'O' wins by connecting the top edge to the bottom edge.
// Player 'X' wins by connecting the left edge to the right edge.
func ResultOf(lines []string) (string, error) {
	if len(lines) == 0 {
		return "", nil // No board, no winner
	}
	height := len(lines)
	width := len(lines[0])
	if width == 0 {
		return "", nil // Empty rows, no winner
	}

	board := make([][]rune, height)
	for r, rowStr := range lines {
		// The test's prepare function removes spaces, so we assume lines are already stripped.
		// Or, more robustly, we could strip them here if not guaranteed.
		// Given the problem description and typical test setups, lines are likely dense.
		board[r] = []rune(rowStr)
	}

	// Check for player O (Top to Bottom)
	if checkWin(board, 'O', height, width) {
		return "O", nil
	}

	// Check for player X (Left to Right)
	if checkWin(board, 'X', height, width) {
		return "X", nil
	}

	return "", nil // No winner
}

// checkWin performs a traversal (BFS) for a given player to see if they have won.
func checkWin(board [][]rune, player rune, height, width int) bool {
	visited := make(map[Point]bool)
	var queue []Point // For BFS

	// Initialize queue with starting positions for the player
	if player == 'O' { // Player O starts from the top row
		for c := 0; c < width; c++ {
			if board[0][c] == player {
				startNode := Point{0, c}
				queue = append(queue, startNode)
				visited[startNode] = true
			}
		}
	} else { // player == 'X', Player X starts from the leftmost column
		for r := 0; r < height; r++ {
			if board[r][0] == player {
				startNode := Point{r, 0}
				queue = append(queue, startNode)
				visited[startNode] = true
			}
		}
	}

	// BFS traversal
	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]

		// Check if current point is a winning point
		if player == 'O' && curr.row == height-1 {
			return true // Reached bottom row for player O
		}
		if player == 'X' && curr.col == width-1 {
			return true // Reached rightmost column for player X
		}

		// Explore neighbors
		for _, neighbor := range getNeighbors(curr, height, width) {
			if board[neighbor.row][neighbor.col] == player && !visited[neighbor] {
				visited[neighbor] = true
				queue = append(queue, neighbor)
			}
		}
	}
	return false // No winning path found
}

// getNeighbors returns valid neighbors for a point on the hexagonal grid.
// The hexagonal grid connections on a 2D array (after space stripping) are:
// (r, c-1)    Left
// (r, c+1)    Right
// (r-1, c)    Top-Left (visually, in the parallelogram representation)
// (r-1, c+1)  Top-Right (visually)
// (r+1, c-1)  Bottom-Left (visually)
// (r+1, c)    Bottom-Right (visually)
func getNeighbors(p Point, height, width int) []Point {
	potentialNeighbors := []Point{
		{p.row, p.col - 1},     // Left
		{p.row, p.col + 1},     // Right
		{p.row - 1, p.col},     // Top-Left
		{p.row - 1, p.col + 1}, // Top-Right
		{p.row + 1, p.col - 1}, // Bottom-Left
		{p.row + 1, p.col},     // Bottom-Right
	}

	var validNeighbors []Point
	for _, n := range potentialNeighbors {
		if n.row >= 0 && n.row < height && n.col >= 0 && n.col < width {
			validNeighbors = append(validNeighbors, n)
		}
	}
	return validNeighbors
}
