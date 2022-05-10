package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	dummy := &ListNode{}
	curr := dummy
	var add int
	for l1 != nil || l2 != nil || add != 0 {
		var sum int
		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}
		sum += add
		add = sum / 10
		sum = sum % 10
		curr.Next = &ListNode{Val: sum}
		curr = curr.Next
	}
	return dummy.Next
}

func main() {

}
