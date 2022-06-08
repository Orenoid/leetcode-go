package main

import "fmt"

func subsets(nums []int) [][]int {
	var result [][]int
	var backTrack func(start int, track []int)
	backTrack = func(start int, track []int) {
		for i := start; i < len(nums); i++ {
			track = append(track, nums[i])
			tmp := make([]int, len(track))
			copy(tmp, track)
			result = append(result, tmp)
			backTrack(i+1, track)
			track = track[:len(track)-1]
		}
	}
	backTrack(0, nil)
	result = append(result, nil)
	return result
}

func main() {
	fmt.Println(subsets([]int{1, 2, 3}))
}
