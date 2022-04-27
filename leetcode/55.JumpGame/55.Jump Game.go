package main

func canJump(nums []int) bool {
	k := nums[0]
	end := len(nums) - 1
	for i := 0; i <= k; i++ {
		k = max(k, i+nums[i])
		if k >= end {
			return true
		}
	}
	return false
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
