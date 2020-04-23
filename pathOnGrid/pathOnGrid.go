package pathongrid

import "math"

func initDPSolution(grid [][]int, k int) [][][]int {
	dpSolution := make([][][]int, len(grid))
	for i := 0; i < len(grid); i++ {
		dpSolution[i] = make([][]int, len(grid[0]))
		for j := 0; j < len(grid[0]); j++ {
			ar := make([]int, k)
			for l := range ar {
				ar[l] = -1
			}
			dpSolution[i][j] = ar
		}
	}
	return dpSolution
}

func initVisited(grid [][]int) [][]bool {
	visited := make([][]bool, len(grid))
	for i := 0; i < len(grid); i++ {
		visited[i] = make([]bool, len(grid[0]))

		// for j := range visited[i] {
		// 	visited[i][j] = false
		// }
	}
	return visited
}

var dx = []int{-1, 0, 0, 1}
var dy = []int{0, -1, 1, 0}
var dpSolution [][][]int
var visited [][]bool

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func dpShortestPath(grid [][]int, i, j, k int) int {

	if k < 0 {
		return math.MaxInt32
	}
	m := len(grid)
	n := len(grid[0])

	if j < 0 || i < 0 || j >= n || i >= m {
		return math.MaxInt32
	}

	if visited[i][j] {
		return math.MaxInt32
	}

	// has already solved there
	if dpSolution[i][j][k] != -1 {
		return dpSolution[i][j][k]
	}
	//find exit
	if i == m-1 && j == n-1 {
		return 0
	}

	visited[i][j] = true
	ans := math.MaxInt32
	for iter := 0; iter < 4; iter++ {
		ans = min(ans, 1+dpShortestPath(grid,
			i+dy[iter], j+dx[iter], k-grid[i][j]))
	}
	visited[i][j] = false

	dpSolution[i][j][k] = ans
	return ans
}

func shortestPath(grid [][]int, k int) int {
	dpSolution = initDPSolution(grid, k+1)
	visited = initVisited(grid)

	sp := dpShortestPath(grid, 0, 0, k)

	if sp == math.MaxInt32 {
		return -1
	} else {
		return sp
	}
}

func ShortestPath(grid [][]int, k int) int {
	return shortestPath(grid, k)
}
