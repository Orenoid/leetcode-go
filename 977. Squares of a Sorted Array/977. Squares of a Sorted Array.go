package main

import "fmt"

func sortedSquares(nums []int) []int {
	length := len(nums)
	res := make([]int, length)
	begin, end := 0, length-1
	current := length - 1
	for begin <= end {
		if nums[begin]*nums[begin] > nums[end]*nums[end] {
			res[current] = nums[begin] * nums[begin]
			begin++
		} else {
			res[current] = nums[end] * nums[end]
			end--
		}
		current--
	}
	return res
}

func main() {
	fmt.Println(sortedSquares([]int{-4, -1, 0, 2, 3, 5}))
}
