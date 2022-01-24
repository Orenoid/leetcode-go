package main

func buildArray(nums []int) []int {
	result := make([]int, len(nums))
	for i := range nums {
		result[i] = nums[nums[i]]
	}
	return result
}
