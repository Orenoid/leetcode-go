package main

import "math"

func minWindow(s string, t string) string {
	if len(s) == 0 || len(t) == 0 {
		return ""
	}
	currCounter, tCounter := map[byte]int{}, map[byte]int{}
	for i := range t {
		tCounter[t[i]]++
	}
	check := func() bool {
		for ch, count := range tCounter {
			if currCounter[ch] < count {
				return false
			}
		}
		return true
	}

	ansL, ansR := -1, -1
	left, right := 0, 0
	minLength := math.MaxInt64
	for ; right < len(s); right++ {
		currCounter[s[right]]++
		for check() && left <= right {
			if right-left+1 < minLength {
				minLength = right - left + 1
				ansL, ansR = left, right
			}
			currCounter[s[left]]--
			left++
		}
	}
	if ansL == -1 {
		return ""
	}
	return s[ansL : ansR+1]

}

func main() {
	println(minWindow("ab", "b"))
}
