package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	var res [][]int
	var queue []*TreeNode
	queue = append(queue, root)

	var currLayerArr []int
	for len(queue) > 0 {
		layerLength := len(queue)
		for i := 0; i < layerLength; i++ {
			node := queue[0]
			queue = queue[1:]
			currLayerArr = append(currLayerArr, node.Val)
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		res = append(res, currLayerArr)
		currLayerArr = []int{}
	}
	return res
}
