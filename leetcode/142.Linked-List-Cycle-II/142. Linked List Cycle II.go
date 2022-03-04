package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func detectCycle(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	slow, fast := head, head
	for {
		slow = slow.Next
		fast = fast.Next
		if fast == nil || fast.Next == nil {
			return nil
		}
		fast = fast.Next
		if slow == fast {
			break
		}
	}
	fast = head
	for slow != fast {
		slow = slow.Next
		fast = fast.Next
	}
	return fast
}
