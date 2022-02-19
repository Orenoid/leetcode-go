package main

import "fmt"

func coinChange(coins []int, amount int) int {
	dp := make([]int, amount+1)
	for j := 0; j <= amount; j++ {
		if j%coins[0] == 0 {
			dp[j] = j / coins[0]
		} else {
			dp[j] = -1
		}
	}
	for i := 1; i < len(coins); i++ {
		dp[0] = 0
		for j := coins[i]; j <= amount; j++ {
			if dp[j] == -1 && dp[j-coins[i]] == -1 {
				dp[j] = -1
			} else if dp[j] == -1 {
				dp[j] = dp[j-coins[i]] + 1
			} else if dp[j-coins[i]] == -1 {
				dp[j] = dp[j]
			} else {
				dp[j] = min(dp[j], dp[j-coins[i]]+1)
			}
		}
	}

	return dp[amount]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	fmt.Println(coinChange([]int{1, 2, 5}, 11))
	fmt.Println(coinChange([]int{2}, 3))
	fmt.Println(coinChange([]int{1}, 0))
	fmt.Println(coinChange([]int{2, 5, 10, 1}, 27))
}
