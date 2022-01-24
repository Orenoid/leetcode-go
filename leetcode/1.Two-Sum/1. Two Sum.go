package main

import "fmt"

func twoSum(nums []int, target int) []int {
	set := make(map[int]int, len(nums))
	for i, num := range nums {
		if subIndex, found := set[target-num]; found {
			return []int{i, subIndex}
		}
		set[num] = i
	}
	return nil
}

func main() {
	fmt.Println(twoSum([]int{2, 7, 11, 15}, 9))
}
