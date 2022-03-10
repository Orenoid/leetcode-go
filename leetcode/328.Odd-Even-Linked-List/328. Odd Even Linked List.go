package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func oddEvenList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	var count int
	var tail *ListNode
	for curr := head; curr != nil; {
		count++
		tail = curr
		curr = curr.Next
	}
	if count <= 2 {
		return head
	}

	prev := &ListNode{Next: head}
	for curr, i := head, 1; i <= count; i++ {
		if i%2 == 0 {
			prev.Next = curr.Next
			tail.Next = curr
			tail = tail.Next
			tail.Next = nil
			curr = prev.Next
		} else {
			prev = prev.Next
			curr = curr.Next
		}
	}
	return head
}

func main() {
	head := &ListNode{Val: 1}
	head.Next = &ListNode{Val: 2}
	head.Next.Next = &ListNode{Val: 3}
	head.Next.Next.Next = &ListNode{Val: 4}
	head.Next.Next.Next.Next = &ListNode{Val: 5}
	head.Next.Next.Next.Next.Next = &ListNode{Val: 6}
	oddEvenList(head)
	for curr := head; curr != nil; {
		fmt.Println(curr.Val)
		curr = curr.Next
	}
}
