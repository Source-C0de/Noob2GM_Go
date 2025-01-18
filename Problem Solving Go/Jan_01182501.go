func minCost(grid [][]int) int {
    m, n := len(grid), len(grid[0])
	// Directions: Right, Left, Down, Up
	directions := [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}

	// Deque implementation using container/list
	dq := list.New()
	// Cost matrix to store minimum cost for each cell
	cost := make([][]int, m)
	for i := range cost {
		cost[i] = make([]int, n)
		for j := range cost[i] {
			cost[i][j] = math.MaxInt32
		}
	}

	// Start from the top-left corner with cost 0
	dq.PushFront([2]int{0, 0})
	cost[0][0] = 0

	for dq.Len() > 0 {
		// Get the front element from the deque
		cell := dq.Remove(dq.Front()).([2]int)
		x, y := cell[0], cell[1]

		// Explore all directions
		for d, dir := range directions {
			nx, ny := x+dir[0], y+dir[1]
			newCost := cost[x][y]
			if grid[x][y] != d+1 { // If the direction doesn't match, add a cost
				newCost++
			}

			// Check bounds and if a better cost is found
			if nx >= 0 && ny >= 0 && nx < m && ny < n && newCost < cost[nx][ny] {
				cost[nx][ny] = newCost
				if grid[x][y] == d+1 {
					dq.PushFront([2]int{nx, ny}) // No modification needed
				} else {
					dq.PushBack([2]int{nx, ny}) // Modification needed
				}
			}
		}
	}

	// Return the minimum cost to reach the bottom-right corner
	return cost[m-1][n-1]
}