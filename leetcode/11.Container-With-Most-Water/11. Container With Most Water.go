package main

import (
	"fmt"
	"math"
)

func maxArea(height []int) int {
	res := math.MinInt
	left, right := 0, len(height)-1
	for left < right {
		if height[left] < height[right] {
			res = max(res, height[left]*(right-left))
			left++
		} else {
			res = max(res, height[right]*(right-left))
			right--
		}
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	fmt.Println(maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7}))
}
