package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func lastRemaining(n int, m int) int {
	ans := 0
	for i := 2; i <= n; i++ {
		ans = (ans + m) % i
	}
	return ans
}

func main() {
	fmt.Println(lastRemaining(10, 17))
}
