package main

import "fmt"

func removeDuplicates(nums []int) int {
	length := len(nums)
	if length <= 1 {
		return length
	}
	current := nums[0]
	slow := 1
	for fast := 1; fast < len(nums); fast++ {
		if nums[fast] != current {
			current = nums[fast]
			nums[slow] = nums[fast]
			slow++
		}
	}
	fmt.Println(nums)
	return slow
}

func main() {
	fmt.Println(removeDuplicates([]int{2, 3, 4, 4, 10}))
}
