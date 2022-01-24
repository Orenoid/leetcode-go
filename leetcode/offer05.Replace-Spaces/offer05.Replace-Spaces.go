package main

import "fmt"

func replaceSpace(s string) string {
	sBytes := []byte(s)
	spaceCount := 0
	for _, ch := range sBytes {
		if ch == ' ' {
			spaceCount++
		}
	}
	extended := make([]byte, spaceCount*2)
	sBytes = append(sBytes, extended...)
	left, right := len(s)-1, len(sBytes)-1
	for left < right {
		if sBytes[left] == ' ' {
			sBytes[right-2], sBytes[right-1], sBytes[right] = '%', '2', '0'
			left--
			right -= 3
		} else {
			sBytes[right] = sBytes[left]
			left--
			right--
		}
	}
	return string(sBytes)
}

func main() {
	fmt.Println(replaceSpace(" hello"))
}
