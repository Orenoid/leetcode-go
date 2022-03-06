package main

import (
	"fmt"
	"sort"
)

func permuteUnique(nums []int) [][]int {
	sort.Ints(nums)
	var res [][]int
	var track []int
	usedFlags := make([]bool, len(nums))
	var backtrack func()

	backtrack = func() {
		for i, num := range nums {
			if (i > 0 && nums[i] == nums[i-1] && !usedFlags[i-1]) || usedFlags[i] {
				continue
			}
			track = append(track, num)
			usedFlags[i] = true
			if len(track) == len(nums) {
				res = append(res, append([]int{}, track...))
				//tmp := make([]int, len(track))
				//copy(tmp, track)
				//res = append(res, tmp)
			} else {
				backtrack()
			}
			track = track[:len(track)-1]
			usedFlags[i] = false
		}
	}
	backtrack()
	return res
}

func main() {
	fmt.Println(permuteUnique([]int{1, 1, 2}))
}
