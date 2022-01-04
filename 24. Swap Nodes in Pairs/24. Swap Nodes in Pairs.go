package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	fakeHead := &ListNode{Next: head}
	prev, curr, next := fakeHead, head, head.Next
	for curr != nil && next != nil {
		nextNext := next.Next
		next.Next = curr
		curr.Next = nextNext
		prev.Next = next

		if nextNext == nil || nextNext.Next == nil {
			break
		}
		prev, curr, next = curr, curr.Next, nextNext.Next

	}
	return fakeHead.Next
}
