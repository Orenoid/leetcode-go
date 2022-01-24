package main

import (
	"fmt"
	"sort"
)

func threeSum(nums []int) [][]int {
	length := len(nums)
	if length < 3 {
		return nil
	}
	sort.Ints(nums)
	var result [][]int

	for i := 0; i < length-2; i++ {
		if nums[i] > 0 {
			continue
		}
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		left, right := i+1, length-1
		for left < right {
			if nums[i]+nums[left]+nums[right] == 0 {
				leftValue, rightValue := nums[left], nums[right]
				result = append(result, []int{nums[i], leftValue, rightValue})
				for left < right && nums[left] == leftValue {
					left++
				}
				for left < right && nums[right] == rightValue {
					right--
				}

			} else if nums[i]+nums[left]+nums[right] < 0 {
				left++
			} else {
				right--
			}
		}
	}
	return result
}

func main() {
	fmt.Println(threeSum([]int{-4, -2, -1, -1, -1, 0, 2, 3, 4}))
	fmt.Println(threeSum([]int{0, 0, 0, 0}))
}
