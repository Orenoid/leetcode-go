package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	curr, next := head, head.Next
	for next != nil {
		nextNext := next.Next
		next.Next = curr
		curr, next = next, nextNext
	}
	head.Next = nil
	return curr
}
