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

func main() {
	fmt.Println(wordBreak("applepenapple", []string{"apple", "pen"}))
	fmt.Println(wordBreak("helloworld", []string{"hello", "world"}))
}
