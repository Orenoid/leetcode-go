package main

import "fmt"

func removeDuplicates(s string) string {
	var stack []int32
	for _, ch := range s {
		if len(stack) != 0 && stack[len(stack)-1] == ch {
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, ch)
		}
	}
	return string(stack)
}

func main() {
	fmt.Println(removeDuplicates("abbacd"))
}
