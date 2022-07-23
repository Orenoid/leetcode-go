package main

import "fmt"

func wordBreak(s string, wordDict []string) bool {
	cache := make(map[string]bool, len(wordDict))
	for _, word := range wordDict {
		cache[word] = true
	}
	dp := make([]bool, len(s)+1)
	dp[0] = true
	for i := 1; i < len(dp); i++ {
		for _, word := range wordDict {
			if i >= len(word) {
				suffix := s[i-len(word) : i]
				dp[i] = dp[i] || (dp[i-len(word)] && cache[suffix])
			}
		}
	}
	return dp[len(dp)-1]
}

func wordBreakBadVersion(s string, wordDict []string) bool {
	// 实现得比较繁琐的版本，上面 dp 从 1 开始定义会简化很多
	dp := make([]bool, len(s))
	cache := make(map[string]bool, len(wordDict))
	for _, word := range wordDict {
		cache[word] = true
	}
	for i := range dp {
		for word := range cache {
			if i-len(word) < -1 {
				continue
			}
			// word 刚好为 s 头部的情况
			if i == len(word)-1 {
				if cache[s[:i+1]] {
					dp[i] = true
					continue
				}
			} else {
				if dp[i-len(word)] && cache[s[i-len(word)+1:i+1]] {
					dp[i] = true
				}
			}

		}
	}
	return dp[len(s)-1]
}

func main() {
	fmt.Println(wordBreak("applepenapplen", []string{"apple", "pen"}))
	fmt.Println(wordBreak("helloworld", []string{"hello", "world"}))
	fmt.Println(wordBreak("a", []string{"a", "world"}))
}
