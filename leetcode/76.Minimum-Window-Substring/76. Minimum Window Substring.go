package main

func minWindow(s string, t string) string {
	if len(s) == 0 || len(t) == 0 {
		return ""
	}
	cache := make(map[rune]int, len(t))
	for _, ch := range t {
		cache[ch]++
	}
	containsAllChar := func() bool {
		for _, count := range cache {
			if count > 0 {
				return false
			}
		}
		return true
	}

	left, right := 0, 0
	minWinLeft, minWinRight := -1, -1
	minLength := len(s) + 1

	for right < len(s) {
		if _, ok := cache[rune(s[right])]; ok {
			cache[rune(s[right])]--
		}
		if containsAllChar() {
			if right-left+1 < minLength {
				minLength = left + right + 1
				minWinLeft, minWinRight = left, right
			}
			for left < right {
				if _, found := cache[rune(s[left])]; found {
					cache[rune(s[left])]++
				}
				left++
				if containsAllChar() {
					if right-left+1 < minLength {
						minLength = right - left + 1
						minWinLeft, minWinRight = left, right
					}
				} else {
					break
				}
			}
		}
		right++
	}
	if minWinLeft == -1 {
		return ""
	}
	return s[minWinLeft : minWinRight+1]
}

func main() {
	println(minWindow("a", "aa"))
}
