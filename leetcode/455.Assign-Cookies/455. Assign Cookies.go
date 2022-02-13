package main

import (
	"fmt"
	"sort"
)

func findContentChildren(g []int, s []int) int {
	if len(s) == 0 {
		return 0
	}
	// g 胃口，s 饼干大小
	sort.Ints(g)
	sort.Ints(s)
	gIndex, sIndex, count := 0, 0, 0
	for ; sIndex < len(s) && gIndex < len(g); sIndex++ {
		if s[sIndex] >= g[gIndex] {
			count++
			gIndex++
		}
	}
	return count
}

func main() {
	fmt.Println(findContentChildren([]int{1, 2}, []int{1, 2, 3}))
}
