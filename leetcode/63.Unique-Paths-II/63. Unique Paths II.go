package main

import "fmt"

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	rows := len(obstacleGrid)
	columns := len(obstacleGrid[0])
	dp := make([][]int, rows)
	for i := range dp {
		dp[i] = make([]int, columns)
	}
	for i := 0; i < columns; i++ {
		if obstacleGrid[0][i] == 1 {
			break
		}
		dp[0][i] = 1
	}
	for i := 0; i < rows; i++ {
		if obstacleGrid[i][0] == 1 {
			break
		}
		dp[i][0] = 1
	}

	for i := 1; i < rows; i++ {
		for j := 1; j < columns; j++ {
			if obstacleGrid[i][j] != 1 {
				dp[i][j] = dp[i-1][j] + dp[i][j-1]
			}
		}
	}
	return dp[rows-1][columns-1]
}

func main() {
	fmt.Println(uniquePathsWithObstacles([][]int{{0, 0, 1}, {0, 0, 0}}))
}
