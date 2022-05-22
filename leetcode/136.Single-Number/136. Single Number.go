package main

import "fmt"

func singleNumber(nums []int) int {
	sum := nums[0]
	for i := 1; i < len(nums); i++ {
		sum ^= nums[i]
	}
	return sum
}

func main() {
	fmt.Println(singleNumber([]int{2, 2, 1, 1, 5}))
}
