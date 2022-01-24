package main

import "fmt"

func isHappy(n int) bool {
	set := map[int]struct{}{}
	for {
		sum := 0
		// 对 n 各位数求平方和
		for n > 0 {
			digit := n % 10
			sum += digit * digit
			n = n / 10
		}
		if sum == 1 {
			return true
		}
		if _, found := set[sum]; found {
			return false
		}
		set[sum] = struct{}{}
		n = sum
	}
}

func main() {
	fmt.Println(isHappy(2))
}
