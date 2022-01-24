package main

func moveZeroes(nums []int) {
	length := len(nums)
	if length == 0 {
		return
	}
	slow := 0
	for fast := 0; fast < length; fast++ {
		if nums[fast] != 0 {
			nums[slow] = nums[fast]
			slow++
		}
	}
	for i := slow; i < length; i++ {
		nums[i] = 0
	}
	return
}

func main() {
	moveZeroes([]int{0, 1, 0, 2, 0, 8, 99})
	moveZeroes([]int{1})
}
