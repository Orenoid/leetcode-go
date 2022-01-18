package main

import "fmt"

func canConstruct(ransomNote string, magazine string) bool {
	m := map[int32]int{}
	for _, ch := range magazine {
		m[ch]++
	}
	for _, ch := range ransomNote {
		if count, found := m[ch]; found && count > 0 {
			m[ch]--
		} else {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(canConstruct("abcc", "abbcc"))
}
