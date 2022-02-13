package main

import "fmt"

func combine(n int, k int) [][]int {
	var res [][]int
	var backTrack func(n, k, start int, track []int)
	backTrack = func(n, k, start int, track []int) {
		if (n-start+1)+len(track) < k {
			// 从 start 开始，剩余的数字加上已有的 track 数量都不足 k 个数字的情况
			return
		}
		for i := start; i <= n; i++ {
			track = append(track, i)
			if len(track) == k {
				tmp := make([]int, k)
				copy(tmp, track)
				res = append(res, tmp)
			} else {
				backTrack(n, k, i+1, track)
			}
			track = track[:len(track)-1]
		}
	}
	backTrack(n, k, 1, []int{})
	return res
}

func main() {
	fmt.Println(combine(4, 1))
}
