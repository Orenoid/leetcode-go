package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func diameterOfBinaryTree(root *TreeNode) int {
	var ans = 1
	var depth func(root *TreeNode) int
	depth = func(root *TreeNode) int {
		if root == nil {
			return 0
		}
		leftDepth := depth(root.Left)
		rightDepth := depth(root.Right)
		ans = max(ans, leftDepth+rightDepth+1)
		return max(leftDepth, rightDepth) + 1
	}
	depth(root)
	return ans - 1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
