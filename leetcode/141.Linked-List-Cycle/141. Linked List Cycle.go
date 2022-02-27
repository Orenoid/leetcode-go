package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {
	if head == nil {
		return false
	}
	slow, fast := head, head
	for slow != nil && fast != nil {
		slow = slow.Next
		fast = fast.Next
		if fast == nil {
			return false
		}
		fast = fast.Next
		if fast == slow {
			return true
		}
	}
	return false
}

func hasCycle2(head *ListNode) bool {
	if head == nil {
		return false
	}
	cache := map[*ListNode]bool{}
	curr := head
	for curr != nil {
		if cache[curr] {
			return true
		}
		cache[curr] = true
		curr = curr.Next
	}
	return false
}
