package main

import "fmt"

func change(amount int, coins []int) int {
	dp := make([]int, amount+1)
	for j := 0; j <= amount; j++ {
		if j == 0 || j%coins[0] == 0 {
			dp[j] = 1
		}
	}
	for i := 1; i < len(coins); i++ {
		for j := coins[i]; j <= amount; j++ {
			dp[j] = dp[j] + dp[j-coins[i]]
		}
	}
	return dp[amount]
}

func main() {
	fmt.Println(change(3, []int{1, 2}))
}
