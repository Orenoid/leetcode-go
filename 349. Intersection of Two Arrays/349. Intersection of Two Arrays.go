package main

import "fmt"

func intersection(nums1 []int, nums2 []int) []int {
	result := make([]int, 0, len(nums1))
	counter := make(map[int]int, len(nums1))

	for _, num := range nums1 {
		if _, found := counter[num]; !found {
			counter[num] = 0
		}
	}
	for _, num := range nums2 {
		if count, found := counter[num]; found && count == 0 {
			counter[num] = counter[num] + 1
			result = append(result, num)
		}
	}
	return result
}

func main() {
	fmt.Println(intersection([]int{1, 2, 2}, []int{2, 3, 1}))
}
