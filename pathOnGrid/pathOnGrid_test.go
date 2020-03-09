package pathongrid

import (
	"testing"
)

func TestShortestPath(t *testing.T) {
	grid := [][]int{}
	grid = append(grid, []int{0, 0, 0})
	grid = append(grid, []int{1, 1, 0})
	grid = append(grid, []int{0, 0, 0})
	grid = append(grid, []int{0, 1, 1})
	grid = append(grid, []int{0, 0, 0})

	k := 1

	expected := 6
	result := ShortestPath(grid, k)

	if expected != result {
		t.Error("expected shortest path", expected,
			"but result is", result)
	}
}

func BenchmarkShortestPath(b *testing.B) {
	grid := [][]int{}
	grid = append(grid, []int{0, 0, 0})
	grid = append(grid, []int{1, 1, 0})
	grid = append(grid, []int{0, 0, 0})
	grid = append(grid, []int{0, 1, 1})
	grid = append(grid, []int{0, 0, 0})

	k := 1
	for i := 0; i < b.N; i++ {
		ShortestPath(grid, k)
	}
}
