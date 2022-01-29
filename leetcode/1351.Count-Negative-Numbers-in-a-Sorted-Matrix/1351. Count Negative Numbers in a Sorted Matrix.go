package main

import "fmt"

func countNegatives(grid [][]int) int {
	count := 0
	for _, row := range grid {
		if len(row) == 0 {
			continue
		}
		left, right := 0, len(row)-1
		for left <= right {
			mid := (left + right) / 2
			if row[mid] < 0 && (mid == 0 || row[mid-1] >= 0) {
				count += len(row) - mid
				break
			} else if row[mid] >= 0 {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
	}
	return count
}

func main() {
	fmt.Println(countNegatives([][]int{{4, 2, 0, -1, -2}, {-2, -3}}))
}
