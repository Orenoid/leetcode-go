package main

import "fmt"

func rob(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	dp := make([]int, len(nums))
	dp[0] = nums[0]
	dp[1] = max(nums[0], nums[1])
	ans := max(dp[0], dp[1])
	for i := 2; i < len(nums); i++ {
		dp[i] = max(dp[i-2]+nums[i], dp[i-1])
		ans = max(ans, dp[i])
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	fmt.Println(rob([]int{2, 7, 9, 3, 1}))
}
