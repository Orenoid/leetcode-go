package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}
	lenA, lenB := 0, 0
	currA, currB := headA, headB
	for currA != nil {
		lenA++
		currA = currA.Next
	}
	for currB != nil {
		lenB++
		currB = currB.Next
	}

	var step int
	var slow, fast *ListNode
	if lenA > lenB {
		step = lenA - lenB
		slow, fast = headB, headA
	} else {
		step = lenB - lenA
		slow, fast = headA, headB
	}
	for i := 0; i < step; i++ {
		fast = fast.Next
	}

	for slow != fast {
		slow = slow.Next
		fast = fast.Next
	}

	return slow
}
