package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(inorder []int, postorder []int) *TreeNode {
	if len(inorder) == 0 || len(postorder) == 0 || len(inorder) != len(postorder) {
		return nil
	}
	rootVal := postorder[len(postorder)-1]
	inorderRootIndex := index(inorder, rootVal)
	if inorderRootIndex == -1 {
		// 理论上不会出现
		return nil
	}
	root := &TreeNode{Val: rootVal}
	root.Left = buildTree(inorder[:inorderRootIndex], postorder[:inorderRootIndex])
	root.Right = buildTree(inorder[inorderRootIndex+1:], postorder[inorderRootIndex:len(postorder)-1])
	return root
}

func index(slice []int, target int) int {
	for i, v := range slice {
		if v == target {
			return i
		}
	}
	return -1
}

func main() {
	fmt.Println(index([]int{1, 2, 3}, 1))
}
