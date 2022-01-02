package main

import "fmt"

func minSubArrayLen(target int, nums []int) int {
	begin, end, currSum, minLength := 0, 0, 0, 0
	found := false

	for ; end < len(nums); end++ {
		currSum += nums[end]
		for currSum >= target {
			currSubLen := end - begin + 1
			if !found || currSubLen < minLength {
				minLength = currSubLen
			}
			found = true
			currSum -= nums[begin]
			begin++
		}
	}
	if !found {
		return 0
	} else {
		return minLength
	}
}

func main() {
	fmt.Println(minSubArrayLen(7, []int{2, 3, 1, 3, 4, 2, 1}))
}
