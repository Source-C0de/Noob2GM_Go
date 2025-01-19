package main

import (
	"container/heap"
	"fmt"
	"math"
)

// A PriorityQueue is a min-heap of cells in the grid
type Cell struct {
	height int
	row    int
	col    int
}

// PriorityQueue implements a min-heap for Cells
type PriorityQueue []Cell

func (pq PriorityQueue) Len() int           { return len(pq) }
func (pq PriorityQueue) Less(i, j int) bool { return pq[i].height < pq[j].height }
func (pq PriorityQueue) Swap(i, j int)      { pq[i], pq[j] = pq[j], pq[i] }

func (pq *PriorityQueue) Push(x interface{}) {
	*pq = append(*pq, x.(Cell))
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[:n-1]
	return item
}

func trapRainWater(heightMap [][]int) int {
	m, n := len(heightMap), len(heightMap[0])
	if m < 3 || n < 3 {
		return 0 // No water can be trapped in a grid smaller than 3x3
	}

	// Priority queue (min-heap)
	pq := &PriorityQueue{}
	heap.Init(pq)

	// Visited matrix
	visited := make([][]bool, m)
	for i := range visited {
		visited[i] = make([]bool, n)
	}

	// Add all boundary cells to the heap and mark them as visited
	for i := 0; i < m; i++ {
		for _, j := range []int{0, n - 1} {
			heap.Push(pq, Cell{heightMap[i][j], i, j})
			visited[i][j] = true
		}
	}
	for j := 0; j < n; j++ {
		for _, i := range []int{0, m - 1} {
			if !visited[i][j] {
				heap.Push(pq, Cell{heightMap[i][j], i, j})
				visited[i][j] = true
			}
		}
	}

	// Directions for moving to neighbors
	directions := [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}

	waterTrapped := 0
	for pq.Len() > 0 {
		cell := heap.Pop(pq).(Cell)

		// Process neighbors
		for _, dir := range directions {
			nx, ny := cell.row+dir[0], cell.col+dir[1]

			// Check bounds and if the cell has already been visited
			if nx >= 0 && nx < m && ny >= 0 && ny < n && !visited[nx][ny] {
				visited[nx][ny] = true

				// Calculate trapped water
				waterTrapped += int(math.Max(0, float64(cell.height-heightMap[nx][ny])))

				// Push the neighbor with the updated height to the heap
				heap.Push(pq, Cell{int(math.Max(float64(cell.height), float64(heightMap[nx][ny]))), nx, ny})
			}
		}
	}

	return waterTrapped
}

func main() {
	heightMap := [][]int{
		{1, 4, 3, 1, 3, 2},
		{3, 2, 1, 3, 2, 4},
		{2, 3, 3, 2, 3, 1},
	}

	fmt.Println("Water trapped:", trapRainWater(heightMap)) // Output: 4
}
