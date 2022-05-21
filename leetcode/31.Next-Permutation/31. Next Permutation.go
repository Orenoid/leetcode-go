package main

import "fmt"

func nextPermutation(nums []int) {
	if len(nums) <= 1 {
		return
	}
	// 从后往前找到第一个降序的数字
	left := len(nums) - 2
	for left >= 0 && nums[left] >= nums[left+1] {
		left--
	}
	if left == -1 {
		reverse(nums)
		return
	}
	// 从后往前找到第一个比 nums[left] 更大的数字
	right := len(nums) - 1
	for nums[right] <= nums[left] {
		right--
	}
	nums[left], nums[right] = nums[right], nums[left]
	reverse(nums[left+1:])
}

func reverse(nums []int) {
	left, right := 0, len(nums)-1
	for left <= right {
		nums[left], nums[right] = nums[right], nums[left]
		left++
		right--
	}
}

func main() {
	nums := []int{1, 2, 3}
	nextPermutation(nums)
	fmt.Println(nums)
}
