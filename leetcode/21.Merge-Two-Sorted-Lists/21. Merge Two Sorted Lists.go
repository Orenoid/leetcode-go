package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	dummy := &ListNode{}
	curr := dummy
	for list1 != nil || list2 != nil {
		var minNode *ListNode
		if list1 == nil || (list2 != nil && list1.Val > list2.Val) {
			minNode = list2
			list2 = list2.Next
		} else if list2 == nil || (list1 != nil && list1.Val <= list2.Val) {
			minNode = list1
			list1 = list1.Next
		}
		curr.Next = minNode
		minNode.Next = nil
		curr = curr.Next
	}
	return dummy.Next
}
