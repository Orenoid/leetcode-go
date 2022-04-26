package main

import "fmt"

func trap(height []int) int {
	leftMax := make([]int, len(height))
	rightMax := make([]int, len(height))
	currLeftMax := height[0]
	for i := 1; i < len(height); i++ {
		leftMax[i] = currLeftMax
		currLeftMax = max(height[i], currLeftMax)
	}
	currRightMax := height[len(height)-1]
	for i := len(height) - 2; i >= 0; i-- {
		rightMax[i] = currRightMax
		currRightMax = max(height[i], currRightMax)
	}
	var ans int
	for i := range height {
		ans += max(0, min(leftMax[i], rightMax[i])-height[i])
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
	fmt.Println(trap([]int{4, 2, 0, 3, 2, 5}))
	fmt.Println(trap([]int{0, 0, 0, 0}))
}
