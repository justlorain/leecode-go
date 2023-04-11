package main

func main() {
	uniquePaths2(5, 5)
}

// f(i, j) = f(i-1, j) + f(i, j-1)
func uniquePaths(m int, n int) int {
	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
		dp[i][0] = 1
	}
	for j := range dp[0] {
		dp[0][j] = 1
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			dp[i][j] = dp[i-1][j] + dp[i][j-1]
		}
	}
	return dp[m-1][n-1]
}

// 使用滚动数组优化后的动态规划
func uniquePaths2(m int, n int) int {
	dp := make([][]int, 2)
	for i := 0; i < 2; i++ {
		dp[i] = make([]int, n)
		dp[i][0] = 1
	}
	for j := 0; j < n; j++ {
		dp[0][j] = 1
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			dp[i%2][j] = dp[(i-1)%2][j] + dp[i%2][j-1]
		}
	}
	return dp[(m-1)%2][n-1]
}
