package main

import "fmt"

func reverseStr(s string, k int) string {
	sBytes := []byte(s)
	for i := 0; i < len(sBytes); i += 2 * k {
		left := i
		right := i + k - 1
		if right > len(sBytes)-1 {
			right = len(sBytes) - 1
		}
		for left < right {
			sBytes[left], sBytes[right] = sBytes[right], sBytes[left]
			left++
			right--
		}
	}
	return string(sBytes)
}

func main() {
	fmt.Println(reverseStr("abcdefgh", 3))
}
