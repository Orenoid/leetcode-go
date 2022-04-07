package main

func merge(nums1 []int, m int, nums2 []int, n int) {
	index1, index2 := m-1, n-1
	for i := len(nums1) - 1; i >= 0; i-- {
		if index1 == -1 {
			nums1[i] = nums2[index2]
			index2--
			continue
		}
		if index2 == -1 {
			nums1[i] = nums1[index1]
			index1--
			continue
		}
		if nums1[index1] > nums2[index2] {
			nums1[i] = nums1[index1]
			index1--
		} else {
			nums1[i] = nums2[index2]
			index2--
		}
	}
}

func main() {
	merge([]int{1, 2, 3, 0, 0, 0}, 3, []int{2, 5, 6}, 3)
	merge([]int{0, 0}, 0, []int{1, 2}, 2)
	merge([]int{0, 0}, 0, []int{1, 2}, 2)
}
