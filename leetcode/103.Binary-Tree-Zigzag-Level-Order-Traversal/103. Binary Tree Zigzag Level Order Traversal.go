package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func zigzagLevelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	var res [][]int
	queue := []*TreeNode{root}
	isOddLayer := true
	for len(queue) > 0 {
		qLen := len(queue)
		layer := make([]int, qLen)
		for i := 0; i < qLen; i++ {
			node := queue[0]
			queue = queue[1:]
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
			layer[i] = node.Val
		}
		if !isOddLayer {
			reverse(layer)
		}
		isOddLayer = !isOddLayer
		res = append(res, layer)
	}
	return res
}

func reverse(slice []int) {
	left, right := 0, len(slice)-1
	for left <= right {
		slice[left], slice[right] = slice[right], slice[left]
		left++
		right--
	}
}
