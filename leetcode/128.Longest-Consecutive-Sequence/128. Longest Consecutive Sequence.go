package main

import (
	"fmt"
)

func longestConsecutive(nums []int) int {
	hash := map[int]bool{}
	for _, num := range nums {
		hash[num] = true
	}
	ans := 0
	for _, num := range nums {
		if hash[num-1] {
			continue
		}
		currLen := 0
		for hash[num] {
			currLen++
			num = num + 1
		}
		if currLen > ans {
			ans = currLen
		}
	}
	return ans
}

func main() {
	fmt.Println(longestConsecutive([]int{100, 4, 200, 1, 201, 202, 2, 5}))
}
