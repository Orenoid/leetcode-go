package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseBetween(head *ListNode, left int, right int) *ListNode {
	if right <= left {
		return head
	}
	dummy := &ListNode{Next: head}
	var stack []*ListNode
	var leftPrev, rightAfter *ListNode

	for i, curr := 0, dummy; i <= right; i++ {
		if i == left-1 {
			leftPrev = curr
		}
		if i == right {
			rightAfter = curr.Next
		}
		if i >= left && i <= right {
			stack = append(stack, curr)
		}
		curr = curr.Next
	}

	leftPrev.Next = stack[len(stack)-1]
	stack[0].Next = rightAfter
	for i := len(stack) - 1; i >= 1; i-- {
		stack[i].Next = stack[i-1]
	}
	return dummy.Next
}

func main() {
	head := &ListNode{Val: 1}
	head.Next = &ListNode{Val: 2}
	head.Next.Next = &ListNode{Val: 3}
	head.Next.Next.Next = &ListNode{Val: 4}
	head.Next.Next.Next.Next = &ListNode{Val: 5}
	head = reverseBetween(head, 2, 4)
	for head != nil {
		fmt.Println(head.Val)
		head = head.Next
	}
}
