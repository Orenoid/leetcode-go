package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxPathSum(root *TreeNode) int {
	if root == nil {
		return 0
	}
	ans := root.Val
	var maxPathSumThatEndsWithRoot func(root *TreeNode) int

	maxPathSumThatEndsWithRoot = func(root *TreeNode) int {
		if root.Left == nil && root.Right == nil {
			ans = max(ans, root.Val)
			return root.Val
		}
		if root.Left == nil {
			maxSum := max(maxPathSumThatEndsWithRoot(root.Right)+root.Val, root.Val)
			ans = max(maxSum, ans)
			return maxSum
		}
		if root.Right == nil {
			maxSum := max(maxPathSumThatEndsWithRoot(root.Left)+root.Val, root.Val)
			ans = max(maxSum, ans)
			return maxSum
		}
		leftPathMaxSum := maxPathSumThatEndsWithRoot(root.Left)
		rightPathMaxSum := maxPathSumThatEndsWithRoot(root.Right)
		maxSumEndWithRoot := max(max(leftPathMaxSum, rightPathMaxSum)+root.Val, root.Val)
		// 需要比较 root 作为中间节点的情况，以及 root 作为断点的情况（即 maxSumEndWithRoot）
		ans = max(maxSumEndWithRoot, leftPathMaxSum+rightPathMaxSum+root.Val, ans)
		return maxSumEndWithRoot
	}
	maxPathSumThatEndsWithRoot(root)
	return ans
}

func max(nums ...int) int {
	if len(nums) == 0 {
		panic("nums must not be empty")
	}
	ans := nums[0]
	for i := range nums {
		if nums[i] > ans {
			ans = nums[i]
		}
	}
	return ans
}
