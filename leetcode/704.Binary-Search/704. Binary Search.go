package main

func search(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}
	begin, end := 0, len(nums)-1
	for begin <= end {
		mid := (begin + end) / 2
		if nums[mid] == target {
			return mid
		}
		if nums[mid] < target {
			begin = mid + 1
		}
		if nums[mid] > target {
			end = mid - 1
		}
	}
	return -1
}
