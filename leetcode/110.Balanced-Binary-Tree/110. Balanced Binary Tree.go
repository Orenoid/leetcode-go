package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}

	var getHeightAndCheckBalanced func(node *TreeNode) (int, bool)
	getHeightAndCheckBalanced = func(node *TreeNode) (int, bool) {
		if node == nil {
			return 0, true
		}
		if node.Left == nil && node.Right == nil {
			return 1, true
		}
		leftH, leftBalanced := getHeightAndCheckBalanced(node.Left)
		rightH, rightBalanced := getHeightAndCheckBalanced(node.Right)
		currH := max(leftH, rightH) + 1
		if !leftBalanced || !rightBalanced {
			return currH, false
		}
		return currH, abs(leftH-rightH) <= 1

	}
	_, result := getHeightAndCheckBalanced(root)
	return result
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}
