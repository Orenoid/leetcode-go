package main

import "fmt"

func lengthOfLongestSubstring(s string) int {
	sLen := len(s)
	counter := make(map[byte]int, sLen)
	maxLen := 0
	right := -1
	for left := 0; left < sLen; left++ {
		for right < sLen-1 && counter[s[right+1]] == 0 {
			right++
			counter[s[right]]++
		}
		maxLen = max(maxLen, right-left+1)
		counter[s[left]]--
	}
	return maxLen
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func lengthOfLongestSubstring2(s string) int {
	counter := make(map[byte]int)
	left, right := 0, 0
	sLen := len(s)
	maxLen, currLen := 0, 0
	for ; right < sLen; right++ {
		counter[s[right]]++
		currLen++
		if !hasRepeatedChar(counter) {
			if currLen > maxLen {
				maxLen = currLen
			}
		} else {
			for {
				counter[s[left]]--
				left++
				currLen--
				if !hasRepeatedChar(counter) {
					break
				}
			}
		}
	}
	return maxLen
}

func hasRepeatedChar(counter map[byte]int) bool {
	for _, count := range counter {
		if count > 1 {
			return true
		}
	}
	return false
}

func main() {
	fmt.Println(lengthOfLongestSubstring("aabc"))
}
