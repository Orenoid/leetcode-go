package main

import (
	"fmt"
)

func canPartition(nums []int) bool {
	s := sum(nums)
	if s%2 != 0 {
		return false
	}
	target := s / 2
	dp := make([]bool, target+1)
	for j := 0; j <= target; j++ {
		if nums[0] == j {
			dp[j] = true
		}
	}
	for i := 1; i < len(nums); i++ {
		for j := target; j >= nums[i]; j-- {
			dp[j] = dp[j] || dp[j-nums[i]]
		}
	}
	return dp[len(dp)-1]
}

func sum(nums []int) int {
	s := 0
	for _, num := range nums {
		s += num
	}
	return s
}

func main() {
	fmt.Println(canPartition([]int{1, 5, 11, 5}))
	fmt.Println(canPartition([]int{1, 2, 3, 5}))
}
