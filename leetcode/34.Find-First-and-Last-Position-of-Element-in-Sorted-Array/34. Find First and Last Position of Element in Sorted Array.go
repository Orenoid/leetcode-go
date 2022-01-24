package main

import "fmt"

func searchRange(nums []int, target int) []int {
	if len(nums) == 0 {
		return []int{-1, -1}
	}
	begin, end := -1, -1
	// 左边界
	left, right := 0, len(nums)-1
	for left <= right {
		mid := (left + right) / 2
		if nums[mid] < target {
			left = mid + 1
		} else if nums[mid] > target {
			right = mid - 1
		} else {
			if mid == 0 || nums[mid-1] != target {
				begin = mid
				break
			} else {
				right = mid - 1
			}

		}
	}
	if begin == -1 {
		return []int{begin, end}
	}
	// 右边界
	left, right = begin, len(nums)-1
	for left <= right {
		mid := (left + right) / 2
		if nums[mid] < target {
			left = mid + 1
		} else if nums[mid] > target {
			right = mid - 1
		} else {
			if mid == len(nums)-1 || nums[mid+1] != target {
				end = mid
				break
			} else {
				left = mid + 1
			}
		}
	}
	return []int{begin, end}
}

func main() {
	fmt.Println(searchRange([]int{1, 4, 4, 4}, 1))
}
