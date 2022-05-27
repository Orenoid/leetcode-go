package main

func sortColors(nums []int) {
	zero, two := 0, len(nums)-1
	for i := 0; i <= two; {
		if nums[i] == 2 {
			nums[i], nums[two] = nums[two], nums[i]
			two--
		} else if nums[i] == 0 {
			nums[i], nums[zero] = nums[zero], nums[i]
			zero++
			i++
		} else {
			i++
		}
	}
}

func main() {
	sortColors([]int{2, 0, 2, 1, 1, 0})
}
