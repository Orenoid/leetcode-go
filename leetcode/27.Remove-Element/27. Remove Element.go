package main

func removeElement(nums []int, val int) int {
	length := len(nums)
	slow := 0
	for fast := 0; fast < length; fast++ {
		if nums[fast] != val {
			nums[slow] = nums[fast]
			slow++
		}
	}
	return slow
}
