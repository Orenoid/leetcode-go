package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	if head == nil {
		return nil
	}

	dummy := &ListNode{Next: head}
	before, prev, curr := dummy, dummy, head
	var after *ListNode
	totalCount := 0
	length := getLength(head)

	for {
		if length-totalCount < k {
			break
		}
		count := 0
		for count < k && curr != nil {
			next := curr.Next
			curr.Next = prev
			prev, curr = curr, next
			after = next
			count++
			totalCount++
		}
		start := before.Next
		before.Next = prev
		start.Next = after
		before = start
		//printList(dummy.Next)
		var _ int
	}
	return dummy.Next
}

func printList(head *ListNode) {
	for curr := head; curr != nil; {
		fmt.Print(curr.Val, " ")
		curr = curr.Next
	}
}

func getLength(head *ListNode) int {
	l := 0
	for curr := head; curr != nil; {
		l++
		curr = curr.Next
	}
	return l
}

func main() {
	head := &ListNode{Val: 1}
	head.Next = &ListNode{Val: 2}
	head.Next.Next = &ListNode{Val: 3}
	head.Next.Next.Next = &ListNode{Val: 4}
	head.Next.Next.Next.Next = &ListNode{Val: 5}
	head.Next.Next.Next.Next.Next = &ListNode{Val: 6}
	printList(reverseKGroup(head, 1))

}
