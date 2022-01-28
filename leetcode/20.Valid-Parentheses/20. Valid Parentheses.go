package main

import "fmt"

func isValid(s string) bool {
	var stack []int32
	mapper := map[int32]int32{
		'(': ')',
		'[': ']',
		'{': '}',
	}
	reversedMapper := map[int32]int32{
		')': '(',
		']': '[',
		'}': '{',
	}
	for _, ch := range s {
		if _, found := mapper[ch]; found {
			stack = append(stack, ch)
		} else {
			if len(stack) == 0 {
				return false
			}
			val := stack[len(stack)-1]
			if val != reversedMapper[ch] {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}
	return len(stack) == 0
}

func main() {
	fmt.Println(isValid("([)]"))
}
