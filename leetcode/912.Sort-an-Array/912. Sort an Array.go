package main

import (
	"fmt"
	"math/rand"
)

func sortArray(nums []int) []int {
	if len(nums) < 2 {
		return nums
	}
	pivotIndex := rand.Int() % len(nums)
	pivot := nums[pivotIndex]
	left, right := 0, len(nums)-1
	nums[pivotIndex], nums[right] = nums[right], nums[pivotIndex]
	for i := range nums {
		if nums[i] < pivot {
			nums[i], nums[left] = nums[left], nums[i]
			left++
		}
	}
	nums[left], nums[right] = nums[right], nums[left]
	sortArray(nums[:left])
	sortArray(nums[left+1:])
	return nums
}

func main() {
	fmt.Println(sortArray([]int{1, 3, 2, 8, 7, 12, 2}))
}
