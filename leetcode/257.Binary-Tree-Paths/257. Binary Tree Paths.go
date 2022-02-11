package main

import "strconv"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func binaryTreePaths(root *TreeNode) []string {
	if root == nil {
		return nil
	}
	var res []string
	var travel func(root *TreeNode, prePath string)
	travel = func(root *TreeNode, prePath string) {
		if root == nil {
			return
		}
		if root.Left == nil && root.Right == nil {
			v := strconv.Itoa(root.Val)
			res = append(res, prePath+v)
			return
		}
		prePath = prePath + strconv.Itoa(root.Val) + "->"
		if root.Left != nil {
			travel(root.Left, prePath)
		}
		if root.Right != nil {
			travel(root.Right, prePath)
		}
	}
	travel(root, "")
	return res
}
