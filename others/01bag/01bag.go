package main

import "fmt"

func bagProblem(weight, value []int, bagWeight int) int {
	dp := make([][]int, len(weight))
	for i := range dp {
		dp[i] = make([]int, bagWeight+1)
	}
	for j := 0; j <= bagWeight; j++ {
		if j >= weight[0] {
			dp[0][j] = value[0]
		}
	}
	for i := 1; i < len(weight); i++ {
		for j := 1; j <= bagWeight; j++ {
			if weight[i] > j {
				dp[i][j] = dp[i-1][j]
			} else {
				dp[i][j] = max(dp[i-1][j], value[i]+dp[i-1][j-weight[i]])
			}
		}
	}
	return dp[len(weight)-1][bagWeight]
}

func bagProblem2(weight, value []int, bagWeight int) int {
	dp := make([]int, bagWeight+1)
	for j := 0; j <= bagWeight; j++ {
		if weight[0] <= j {
			dp[j] = value[0]
		}
	}
	for i := 1; i < len(weight); i++ {
		prevDp := make([]int, bagWeight+1)
		copy(prevDp, dp)
		for j := 1; j <= bagWeight; j++ {
			if weight[i] <= j {
				dp[j] = max(prevDp[j], value[i]+prevDp[j-weight[i]])
			}
		}
	}
	return dp[bagWeight]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	//fmt.Println(bagProblem([]int{1, 3, 4}, []int{15, 20, 30}, 8))
	fmt.Println(bagProblem2([]int{1, 3, 4}, []int{15, 20, 30}, 8))
}
