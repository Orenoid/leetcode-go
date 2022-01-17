package main

func fourSumCount(nums1 []int, nums2 []int, nums3 []int, nums4 []int) int {
	counter := map[int]int{}
	for _, num1 := range nums1 {
		for _, num2 := range nums2 {
			counter[num1+num2]++
		}
	}
	count := 0
	for _, num3 := range nums3 {
		for _, num4 := range nums4 {
			count += counter[-num3-num4]
		}
	}
	return count
}
