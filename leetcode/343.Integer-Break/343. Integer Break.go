package main

import "fmt"

func integerBreak(n int) int {
	dp := make([]int, n+1)
	for i := 2; i <= n; i++ {
		curMax := 0
		for j := 1; j <= i-1; j++ {
			if j != i-1 {
				curMax = max(curMax, max(j*(i-j), j*dp[i-j]))
			} else {
				// j==i-1 时，dp[1] 没有意义
				curMax = max(curMax, j*(i-j))
			}
		}
		dp[i] = curMax
	}
	return dp[n]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	fmt.Println(integerBreak(9))
}
