package main

import "fmt"

func findTargetSumWays(nums []int, target int) int {
	sumOfNums := getSum(nums)
	if (target+sumOfNums)%2 != 0 {
		return 0
	}
	if (target + sumOfNums) < 0 {
		return 0
	}
	bagWeight := (target + sumOfNums) / 2
	dp := make([]int, bagWeight+1)
	for j := 0; j <= bagWeight; j++ {
		if j == nums[0] {
			dp[j]++
		}
		if j == 0 {
			dp[j]++
		}
	}
	for i := 1; i < len(nums); i++ {
		for j := bagWeight; j >= nums[i]; j-- {
			dp[j] = dp[j] + dp[j-nums[i]]
		}
	}
	return dp[len(dp)-1]
}

func getSum(nums []int) int {
	s := 0
	for _, num := range nums {
		s += num
	}
	return s
}

func main() {
	fmt.Println(findTargetSumWays([]int{0, 0, 0, 0, 0, 0, 0, 0, 1}, 1))
}
