package main

import "fmt"

func main() {
	arr := [][]int{
		{0, 0, 0},
		{0, 1, 0},
		{0, 0, 0},
	}
	res := uniquePathsWithObstacles2(arr)
	fmt.Println(res)
}

// 使用滚动数组优化的官方题解
func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	n, m := len(obstacleGrid), len(obstacleGrid[0])
	f := make([]int, m)
	if obstacleGrid[0][0] == 0 {
		f[0] = 1
	}
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if obstacleGrid[i][j] == 1 {
				f[j] = 0
				continue
			}
			if j-1 >= 0 && obstacleGrid[i][j-1] == 0 {
				f[j] += f[j-1]
			}
		}
	}
	return f[len(f)-1]
}

// 没有使用滚动数组优化
func uniquePathsWithObstacles2(obstacleGrid [][]int) int {
	// 行
	m := len(obstacleGrid)
	// 列
	n := len(obstacleGrid[0])
	// 初始化二维数组的时候如果长度设置为0的话，不能直接下标访问，只能先append
	// 初始化二维数组
	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
	}
	// 初始化第一列
	for i := 0; i < m; i++ {
		if obstacleGrid[i][0] == 1 {
			break
		}
		dp[i][0] = 1
	}
	// 初始化第一行
	for i := 0; i < n; i++ {
		if obstacleGrid[0][i] == 1 {
			break
		}
		dp[0][i] = 1
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			// 如果当前位置是障碍物，则不可达
			if obstacleGrid[i][j] == 1 {
				dp[i][j] = 0
			} else {
				dp[i][j] = dp[i-1][j] + dp[i][j-1]
			}
		}
	}
	return dp[m-1][n-1]
}
