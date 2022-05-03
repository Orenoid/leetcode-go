package types

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestListNodeString(t *testing.T) {
	head := &ListNode{Val: 1}
	head.Next = &ListNode{Val: 2}
	head.Next.Next = &ListNode{Val: 3}
	assert.Equal(t, "1->2->3", head.String())
}

func TestNewListFromSlice(t *testing.T) {

	assert.Empty(t, NewListFromSlice([]int{}))
	assert.Empty(t, NewListFromSlice(nil))

	node := NewListFromSlice([]int{2, 3, 4})
	var nums []int
	for node != nil {
		nums = append(nums, node.Val)
		node = node.Next
	}
	assert.Equal(t, []int{2, 3, 4}, nums)
}

func TestTreeNodeString(t *testing.T) {
	root := &TreeNode{Val: 0}
	root.Left = &TreeNode{Val: 1}
	root.Right = &TreeNode{Val: 2}
	root.Left.Left = &TreeNode{Val: 3}
	root.Left.Right = &TreeNode{Val: 4}
	assert.Equal(t, "[0, 1, 2, 3, 4, null, null]", root.String())

}
