package main

import "fmt"

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	counter := map[uint8]int{}
	for i := 0; i < len(s); i++ {
		if _, ok := counter[s[i]]; ok {
			counter[s[i]] = counter[s[i]] + 1
		} else {
			counter[s[i]] = 1
		}
	}
	for i := 0; i < len(t); i++ {
		if count, found := counter[t[i]]; found && count > 0 {
			counter[t[i]] = counter[t[i]] - 1
		} else {
			return false
		}
	}

	return true
}

func main() {
	fmt.Println(isAnagram("ab", "b"))
}
