package main

import "fmt"

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	totalLength := len(nums1) + len(nums2)
	if totalLength&1 == 1 {
		return float64(findKth(nums1, nums2, totalLength/2+1))
	} else {
		return float64(findKth(nums1, nums2, totalLength/2)+findKth(nums1, nums2, totalLength/2+1)) / 2
	}
}

func findKth(nums1 []int, nums2 []int, k int) int {
	if len(nums1) == 0 {
		return nums2[k-1]
	}
	if len(nums2) == 0 {
		return nums1[k-1]
	}
	if k == 1 {
		return min(nums1[0], nums2[0])
	}
	halfK := min(k/2, len(nums1), len(nums2))
	if nums1[halfK-1] > nums2[halfK-1] {
		return findKth(nums1, nums2[halfK:], k-halfK)
	} else {
		return findKth(nums1[halfK:], nums2, k-halfK)
	}
}

func min(nums ...int) int {
	res := nums[0]
	for _, num := range nums {
		if num < res {
			res = num
		}
	}
	return res
}

func main() {
	fmt.Println(findMedianSortedArrays([]int{1, 2, 3, 5, 6, 7}, []int{1, 8, 9}))
	fmt.Println(findMedianSortedArrays([]int{2, 3}, []int{4, 5}))
}
