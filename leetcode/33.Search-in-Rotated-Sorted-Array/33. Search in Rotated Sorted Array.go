package main

import "fmt"

func search(nums []int, target int) int {
	// 找出旋转点
	rStart := 0
	if nums[0] > nums[len(nums)-1] {
		left, right := 0, len(nums)-1
		for left <= right {
			mid := (left + right) / 2
			if mid > 0 && nums[mid] < nums[mid-1] {
				rStart = mid
				break
			} else if mid == 0 {
				rStart = 1
				break
			} else if nums[mid] > nums[0] {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
	}
	// 找出目标数字
	res := mySearch(nums, rStart, len(nums)-1, target)
	if res != -1 {
		return res
	}
	if rStart > 0 {
		return mySearch(nums, 0, rStart-1, target)
	} else {
		return -1
	}
}

func mySearch(nums []int, begin, end int, target int) int {
	if len(nums) == 0 {
		return -1
	}
	left, right := begin, end
	for left <= right {
		mid := (left + right) / 2
		if nums[mid] == target {
			return mid
		} else if nums[mid] > target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return -1
}

func main() {
	fmt.Println(search([]int{7, 8, 9, 1, 2, 3, 5}, 8))
}
