package main

import (
	"fmt"
)

func generateParenthesis(n int) []string {
	var result []string
	var backTrack func(stack []byte, track []byte)
	backTrack = func(stack []byte, track []byte) {
		for _, bracket := range []byte{'(', ')'} {
			if (len(stack) > 0 && stack[len(stack)-1] == ')') || (len(stack) == 0 && bracket == ')') {
				return
			}
			track = append(track, bracket)
			var popout byte
			if len(stack) > 0 && stack[len(stack)-1] == '(' && bracket == ')' {
				popout = stack[len(stack)-1]
				stack = stack[:len(stack)-1]
			} else {
				stack = append(stack, bracket)
			}

			if len(track) == n+n {
				if len(stack) == 0 {
					result = append(result, string(track))
				}
			} else {
				backTrack(stack, track)
			}

			// 回溯
			track = track[:len(track)-1]
			if popout != 0 {
				stack = append(stack, popout)
			} else {
				stack = stack[:len(stack)-1]
			}
		}
	}
	backTrack(nil, nil)
	return result
}

func main() {
	fmt.Println(generateParenthesis(3))
}
