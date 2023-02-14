package main

import "math"

func main() {

}

// 每遍历到一个格子，就把面积加一
func maxAreaOfIsland(grid [][]int) int {
	res := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == 1 {
				a := area(grid, i, j)
				res = int(math.Max(float64(res), float64(a)))
			}
		}
	}
	return res
}

func area(grid [][]int, r, c int) int {
	if !inArea(grid, r, c) {
		return 0
	}
	if grid[r][c] != 1 {
		return 0
	}
	grid[r][c] = 2
	return 1 + area(grid, r-1, c) + area(grid, r+1, c) + area(grid, r, c-1) + area(grid, r, c+1)
}

func inArea(grid [][]int, r, c int) bool {
	return 0 <= r && r < len(grid) && 0 <= c && c < len(grid[0])
}
