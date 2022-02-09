package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func preorderTraversal(root *TreeNode) []int {
	var res []int
	var trav func(root *TreeNode)
	trav = func(root *TreeNode) {
		if root == nil {
			return
		}
		res = append(res, root.Val)
		trav(root.Left)
		trav(root.Right)
	}
	trav(root)
	return res
}

func preorderTraversal2(root *TreeNode) []int {
	var res []int
	var stack []*TreeNode
	stack = append(stack, root)
	for len(stack) > 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if node == nil {
			continue
		}
		res = append(res, node.Val)
		stack = append(stack, node.Right)
		stack = append(stack, node.Left)
	}
	return res
}
