package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type WrappedNode struct {
	*TreeNode
	Position int
}

func widthOfBinaryTree(root *TreeNode) int {
	if root == nil {
		return 0
	}
	maxWidth := 0
	queue := []WrappedNode{{root, 0}}
	for len(queue) > 0 {
		layerNodeCount := len(queue)
		left := -1
		currWidth := 0
		for i := 0; i < layerNodeCount; i++ {
			node := queue[0]
			queue = queue[1:]
			if node.TreeNode != nil {
				if left == -1 {
					left = node.Position
				}
				currWidth = node.Position - left + 1
				queue = append(queue, WrappedNode{node.Left, node.Position * 2})
				queue = append(queue, WrappedNode{node.Right, node.Position*2 + 1})
			}
		}
		maxWidth = max(maxWidth, currWidth)
	}
	return maxWidth
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
