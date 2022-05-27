package main

import "fmt"

func minPathSum(grid [][]int) int {
	dp := make([][]int, len(grid))
	for i := range grid {
		dp[i] = make([]int, len(grid[0]))
	}
	for sum, i := 0, 0; i < len(grid[0]); i++ {
		dp[0][i] = sum + grid[0][i]
		sum = dp[0][i]
	}
	for sum, i := 0, 0; i < len(grid); i++ {
		dp[i][0] = sum + grid[i][0]
		sum = dp[i][0]
	}
	for i := 1; i < len(grid); i++ {
		for j := 1; j < len(grid[0]); j++ {
			dp[i][j] = grid[i][j] + min(dp[i][j-1], dp[i-1][j])
		}
	}
	return dp[len(dp)-1][len(dp[0])-1]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	fmt.Println(minPathSum([][]int{{1, 3, 1}, {1, 5, 1}, {4, 2, 1}}))
}
