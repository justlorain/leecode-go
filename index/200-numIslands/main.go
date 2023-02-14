package main

func main() {

}

// DFS 网格 DFS 通用模板
func DFS(grid [][]int, r, c int) {
	inArea := func(grid [][]int, r, c int) bool {
		return 0 <= r && r < len(grid) && 0 <= c && c < len(grid[0])
	}
	// 判断 base case
	// 如果坐标 (r, c) 超出网格范围，直接返回
	if !inArea(grid, r, c) {
		return
	}
	// 如果这个格子不是岛屿，直接返回
	if grid[r][c] != 1 {
		return
	}
	// 将格子标记为遍历过
	grid[r][c] = 2
	DFS(grid, r-1, c)
	DFS(grid, r+1, c)
	DFS(grid, r, c-1)
	DFS(grid, r, c+1)
}

func numIslands(grid [][]byte) int {
	res := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == '1' {
				res++
				dfs(grid, i, j)
			}
		}
	}
	return res
}

func dfs(grid [][]byte, r, c int) {
	if r < 0 || r >= len(grid) || c < 0 || c >= len(grid[0]) {
		return
	}
	if grid[r][c] == '0' {
		return
	}
	grid[r][c] = '0'
	dfs(grid, r-1, c)
	dfs(grid, r+1, c)
	dfs(grid, r, c-1)
	dfs(grid, r, c+1)
}
