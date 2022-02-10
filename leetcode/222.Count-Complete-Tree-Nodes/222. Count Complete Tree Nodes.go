package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func countNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}
	leftDepth, rightDepth := 1, 1
	left := root.Left
	right := root.Right
	for left != nil {
		leftDepth++
		left = left.Left
	}
	for right != nil {
		rightDepth++
		right = right.Right
	}
	if leftDepth == rightDepth {
		return 1<<leftDepth - 1
	}
	return countNodes(root.Left) + countNodes(root.Right) + 1
}
