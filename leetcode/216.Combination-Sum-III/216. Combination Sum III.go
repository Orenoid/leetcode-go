package main

import "fmt"

func combinationSum3(k int, n int) [][]int {
	var res [][]int
	var backTrack func(start int, sum int, track []int)
	backTrack = func(start int, sum int, track []int) {
		if sum >= n {
			return
		}
		for i := start; i <= 9; i++ {
			sum += i
			track = append(track, i)
			if len(track) == k && sum == n {
				tmp := make([]int, k)
				copy(tmp, track)
				res = append(res, tmp)
			} else {
				backTrack(i+1, sum, track)
			}
			track = track[:len(track)-1]
			sum -= i
		}
	}
	backTrack(1, 0, []int{})
	return res
}

func main() {
	fmt.Println(combinationSum3(4, 1))
}
