package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	fast := head
	dummy := &ListNode{Next: head}

	// find
	for i := 1; i <= n-1; i++ {
		fast = fast.Next
	}
	slowPrev, slow := dummy, head
	for fast.Next != nil {
		fast = fast.Next
		slow = slow.Next
		slowPrev = slowPrev.Next
	}
	// remove
	slowPrev.Next = slow.Next
	return dummy.Next
}
