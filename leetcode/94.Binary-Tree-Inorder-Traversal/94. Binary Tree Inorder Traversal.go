package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {
	var res []int
	var trav func(node *TreeNode)
	trav = func(root *TreeNode) {
		if root == nil {
			return
		}
		trav(root.Left)
		res = append(res, root.Val)
		trav(root.Right)
	}
	trav(root)
	return res
}

func inorderTraversal2(root *TreeNode) []int {
	var res []int
	var stack []*TreeNode
	cur := root

	for cur != nil || len(stack) > 0 {
		if cur != nil {
			stack = append(stack, cur)
			cur = cur.Left
		} else {
			cur = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			res = append(res, cur.Val)
			cur = cur.Right
		}
	}
	return res
}
