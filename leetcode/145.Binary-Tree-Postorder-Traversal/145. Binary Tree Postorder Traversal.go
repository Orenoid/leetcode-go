package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func postorderTraversal(root *TreeNode) []int {
	var res []int
	var trav func(node *TreeNode)
	trav = func(root *TreeNode) {
		if root == nil {
			return
		}
		trav(root.Left)
		trav(root.Right)
		res = append(res, root.Val)
	}
	trav(root)
	return res
}

func postorderTraversal2(root *TreeNode) []int {
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
		stack = append(stack, node.Left)
		stack = append(stack, node.Right)
	}
	reverse(res)
	return res
}

func reverse(values []int) {
	l, r := 0, len(values)-1
	for l < r {
		values[l], values[r] = values[r], values[l]
		l++
		r--
	}
	return
}
