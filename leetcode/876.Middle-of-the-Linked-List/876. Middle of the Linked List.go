package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func middleNode(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	dummy := &ListNode{Next: head}
	slow, fast := dummy, dummy
	for fast != nil {
		fast = fast.Next
		if fast != nil {
			fast = fast.Next
		}
		slow = slow.Next
	}
	return slow
}
