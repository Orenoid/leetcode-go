package main

import "fmt"

func maxProfit(prices []int) int {
	dp := make([]int, len(prices))
	dp[0] = prices[0]
	ans := 0
	for i := 1; i < len(dp); i++ {
		ans = max(prices[i]-dp[i-1], ans)
		dp[i] = min(dp[i-1], prices[i])
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	fmt.Println(maxProfit([]int{7, 1, 5, 3, 6, 4}))
}
