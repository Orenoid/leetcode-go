package main

import "fmt"

func permute(nums []int) [][]int {
	var res [][]int
	var backTrack func(track []int)
	usedFlags := make([]bool, len(nums))
	backTrack = func(track []int) {
		for i := range nums {
			if usedFlags[i] {
				continue
			}
			usedFlags[i] = true
			track = append(track, nums[i])
			if len(nums) == len(track) {
				tmp := make([]int, len(track))
				copy(tmp, track)
				res = append(res, tmp)
			} else {
				backTrack(track)
			}
			usedFlags[i] = false
			track = track[:len(track)-1]
		}
	}
	backTrack(nil)
	return res
}

func main() {
	fmt.Println(permute([]int{1, 2, 3}))
}
