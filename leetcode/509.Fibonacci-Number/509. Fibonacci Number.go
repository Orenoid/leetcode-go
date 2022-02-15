package main

import "fmt"

func fib(n int) int {
	if n <= 1 {
		return n
	}
	a, b, c := 0, 1, 1
	for i := 2; i <= n; i++ {
		c = a + b
		a = b
		b = c
	}
	return c
}

// 0 1 1 2 3
func main() {
	fmt.Println(fib(5))
}
