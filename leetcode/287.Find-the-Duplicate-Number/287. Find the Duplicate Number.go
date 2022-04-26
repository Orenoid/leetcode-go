package main

import (
	"fmt"
)

func findDuplicate(nums []int) int {
	n := len(nums) - 1
	left, right := 1, n
	for left <= right {
		if left == right {
			return left
		}
		mid := (left + right) / 2
		count := 0
		for _, num := range nums {
			if num <= mid {
				count++
			}
		}
		if count <= mid {
			left = mid + 1
		} else if count > mid {
			right = mid
		}
	}
	return 0
}

func main() {
	fmt.Println(findDuplicate([]int{2, 2, 2, 2, 2}))
}
