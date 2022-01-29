package main

import "fmt"

func maxSlidingWindow(nums []int, k int) []int {
	queue := make([]int, 0, k)
	// 遍历第一个窗口
	for i := 0; i < k; i++ {
		for len(queue) > 0 && queue[len(queue)-1] < nums[i] {
			queue = queue[:len(queue)-1]
		}
		queue = append(queue, nums[i])
	}
	result := []int{queue[0]}

	for slow, fast := 0, k; fast < len(nums); {
		if nums[slow] == queue[0] {
			queue = queue[1:]
		}
		for len(queue) > 0 && queue[len(queue)-1] < nums[fast] {
			queue = queue[:len(queue)-1]
		}
		queue = append(queue, nums[fast])
		result = append(result, queue[0])
		slow++
		fast++
	}
	return result
}

func main() {
	fmt.Println(maxSlidingWindow([]int{-7, -8, 7, 5, 7, 1, 6, 0}, 4))
}
