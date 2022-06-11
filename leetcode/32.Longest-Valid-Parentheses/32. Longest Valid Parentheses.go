package main

import "fmt"

func longestValidParentheses(s string) int {
	if len(s) <= 1 {
		return 0
	}
	// dp[i] 表示以 s[i] 结尾的最长括号子串长度
	dp := make([]int, len(s))
	if s[1] == ')' && s[0] == '(' {
		dp[1] = 2
	}
	ans := dp[1]
	for i := 2; i < len(s); i++ {
		if s[i] == '(' {
			continue
		}
		if s[i-1] == '(' {
			dp[i] = 2 + dp[i-2]
		} else if dp[i-1] > 0 && i-dp[i-1]-1 >= 0 && s[i-dp[i-1]-1] == '(' {
			// 如果 s[i-1] 是 ')' 且存在以 s[i-1] 为结尾匹配子串，则检查这个子串更前面的字符
			dp[i] = dp[i-1] + 2
			if i-dp[i-1]-2 >= 0 {
				// 如果再往前还有另一个子串，则把它的长度也加上
				dp[i] += dp[i-dp[i-1]-2]
			}
		}
		ans = max(ans, dp[i])
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	fmt.Println(longestValidParentheses(")"))
}
