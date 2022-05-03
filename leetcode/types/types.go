package types

import (
	"fmt"
	"strconv"
	"strings"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

// String 从当前节点开始打印链表
func (receiver *ListNode) String() string {
	if receiver == nil {
		return ""
	}
	res := fmt.Sprintf("%d", receiver.Val)
	curr := receiver.Next
	for curr != nil {
		res += fmt.Sprintf("->%d", curr.Val)
		curr = curr.Next
	}
	return res
}

// NewListFromSlice 从切片构造链表
func NewListFromSlice(nums []int) *ListNode {
	if len(nums) == 0 {
		return nil
	}
	dummy := &ListNode{}
	curr := dummy
	for i := range nums {
		curr.Next = &ListNode{Val: nums[i]}
		curr = curr.Next
	}
	return dummy.Next
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// String 返回满二叉树数组形式的字符串
func (receiver *TreeNode) String() string {
	if receiver == nil {
		return ""
	}
	var res []string
	var queue []*TreeNode
	queue = append(queue, receiver)
	for len(queue) > 0 {
		currLayerLength := len(queue)
		nextLayerAllNil := true
		for i := 0; i < currLayerLength; i++ {
			node := queue[0]
			queue = queue[1:]
			if node != nil {
				res = append(res, strconv.Itoa(node.Val))
				queue = append(queue, node.Left, node.Right)
				if node.Left != nil || node.Right != nil {
					nextLayerAllNil = false
				}
			} else {
				res = append(res, "null")
				queue = append(queue, nil, nil)
			}
		}
		if nextLayerAllNil {
			break
		}
	}
	return fmt.Sprintf("[%s]", strings.Join(res, ", "))
}
