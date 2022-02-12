package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func constructMaximumBinaryTree(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	maxIndex := findMaxIndex(nums)
	maxValue := nums[maxIndex]
	root := &TreeNode{Val: maxValue}
	root.Left = constructMaximumBinaryTree(nums[:maxIndex])
	if maxIndex < len(nums)-1 {
		root.Right = constructMaximumBinaryTree(nums[maxIndex+1:])
	} else {
		root.Right = constructMaximumBinaryTree(nil)
	}
	return root
}

func findMaxIndex(slice []int) int {
	max := slice[0]
	index := 0
	for i, v := range slice {
		if v > max {
			max = v
			index = i
		}
	}
	return index
}

func main() {
	fmt.Println(findMaxIndex([]int{1, 22, 3, 4}))
}
