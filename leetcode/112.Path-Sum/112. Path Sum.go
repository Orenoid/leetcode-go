package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func hasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}
	var travel func(root *TreeNode, parentSum int) bool
	travel = func(root *TreeNode, parentSum int) bool {
		if root == nil {
			return false
		}
		currSum := parentSum + root.Val
		if currSum == targetSum && root.Left == nil && root.Right == nil {
			return true
		}
		return travel(root.Left, currSum) || travel(root.Right, currSum)
	}
	return travel(root, 0)
}
