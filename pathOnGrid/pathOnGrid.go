package pathongrid

import "math"

func allocateDPSolution(grid [][]int, k int) [][][]int {
	dpSolution := make([][][]int, len(grid))
	for i := 0; i < len(grid); i++ {
		dpSolution[i] = make([][]int, len(grid[0]))
		for j := 0; j < len(grid[0]); j++ {
			ar := make([]int, k)
			for l := range ar {
				ar[l] = math.MaxInt32
			}
			dpSolution[i][j] = ar
		}
	}
	return dpSolution
}

func allocateVisited(grid [][]int) [][]bool {
	visited := make([][]bool, len(grid))
	for i := 0; i < len(grid); i++ {
		visited[i] = make([]bool, len(grid[0]))

		for j := range visited[i] {
			visited[i][j] = false
		}
	}
	return visited
}

func shortestPath(grid [][]int, k int) int {
	dpSolution := allocateDPSolution(grid, k)
	visited := allocateVisited(grid)

	println(dpSolution, visited)
	return 0
}

func ShortestPath(grid [][]int, k int) int {
	return shortestPath(grid, k)
}
