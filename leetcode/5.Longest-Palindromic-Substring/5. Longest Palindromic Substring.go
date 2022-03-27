package main

import "fmt"

func longestPalindrome(s string) string {
	sLen := len(s)
	dp := make([][]bool, sLen)
	for i := 0; i < sLen; i++ {
		dp[i] = make([]bool, sLen)
	}
	for i := 0; i < sLen; i++ {
		dp[i][i] = true
	}
	maxLen := 1
	ansLeft, ansRight := 0, 0
	for subLen := 2; subLen <= sLen; subLen++ {
		for left := 0; left < sLen-1; left++ {
			right := left + subLen - 1
			if right >= sLen {
				continue
			}
			if s[left] != s[right] {
				dp[left][right] = false
			} else {
				if subLen == 2 {
					dp[left][right] = true
				} else {
					dp[left][right] = dp[left+1][right-1]
				}
			}
			if dp[left][right] == true && subLen > maxLen {
				maxLen = subLen
				ansLeft, ansRight = left, right
			}
		}
	}
	return s[ansLeft : ansRight+1]
}

func main() {
	fmt.Println(longestPalindrome("c"))
}
