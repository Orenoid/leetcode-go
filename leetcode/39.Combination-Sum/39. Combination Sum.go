package main

import "fmt"

func combinationSum(candidates []int, target int) [][]int {
	var res [][]int
	var backTrack func(start int, sum int, track []int)
	backTrack = func(start int, sum int, track []int) {
		if sum > target {
			return
		}
		for i := start; i < len(candidates); i++ {
			sum += candidates[i]
			track = append(track, candidates[i])
			if sum == target {
				tmp := make([]int, len(track))
				copy(tmp, track)
				res = append(res, tmp)
			} else {
				backTrack(i, sum, track)
			}
			track = track[:len(track)-1]
			sum -= candidates[i]
		}
	}
	backTrack(0, 0, []int{})
	return res
}

func main() {
	fmt.Println(combinationSum([]int{2, 3, 5}, 8))
	fmt.Println(combinationSum([]int{2}, 1))
}
