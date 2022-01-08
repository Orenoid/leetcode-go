package main

import "fmt"

type Node struct {
	Next *Node
	Prev *Node
	Val  int
}

type MyLinkedList struct {
	Head   *Node
	End    *Node
	Length int
}

func Constructor() MyLinkedList {
	return MyLinkedList{Length: 0}
}

func (receiver *MyLinkedList) Get(index int) int {
	curr := receiver.Head
	for i := 0; i < index && curr != nil; i++ {
		curr = curr.Next
	}
	if curr != nil {
		return curr.Val
	} else {
		return -1
	}
}

func (receiver *MyLinkedList) AddAtHead(val int) {
	if receiver.Head == nil {
		newNode := &Node{Val: val}
		receiver.Head = newNode
		receiver.End = newNode
		receiver.Length++
		return
	}
	newHead := &Node{Val: val, Next: receiver.Head}
	receiver.Head.Prev = newHead
	receiver.Head = newHead
	receiver.Length++
}

func (receiver *MyLinkedList) AddAtTail(val int) {
	if receiver.End == nil {
		newNode := &Node{Val: val}
		receiver.Head = newNode
		receiver.End = newNode
		receiver.Length++
		return
	}
	newEnd := &Node{Val: val, Prev: receiver.End}
	receiver.End.Next = newEnd
	receiver.End = newEnd
	receiver.Length++
}

func (receiver *MyLinkedList) AddAtIndex(index int, val int) {
	// 单独处理 index 刚好在最后一个节点后面的情况
	if index == receiver.Length {
		newNode := &Node{Val: val, Prev: receiver.End}
		if receiver.End == nil {
			receiver.Head = newNode
		} else {
			receiver.End.Next = newNode
		}
		receiver.End = newNode
		receiver.Length++
		return
	}

	curr := receiver.Head
	for i := 0; i < index && curr != nil; i++ {
		curr = curr.Next
	}
	if curr != nil {
		newNode := &Node{Val: val, Next: curr, Prev: curr.Prev}
		if curr.Prev != nil {
			curr.Prev.Next = newNode
		} else {
			receiver.Head = newNode
		}
		curr.Prev = newNode
		receiver.Length++
	}
}

func (receiver *MyLinkedList) DeleteAtIndex(index int) {
	curr := receiver.Head
	for i := 0; i < index && curr != nil; i++ {
		curr = curr.Next
	}
	if curr != nil {
		if curr.Prev != nil {
			curr.Prev.Next = curr.Next
		} else {
			receiver.Head = curr.Next
		}
		if curr.Next != nil {
			curr.Next.Prev = curr.Prev
		} else {
			receiver.End = curr.Prev
		}
		receiver.Length--
	}
}

func (receiver MyLinkedList) print() {
	for curr := receiver.Head; curr != nil; {
		fmt.Print(curr.Val, " ")
		curr = curr.Next
	}
}

func main() {
	l := Constructor()
	l.AddAtHead(7)
	l.AddAtHead(2)
	l.AddAtHead(1)
	l.AddAtIndex(3, 0)
	l.DeleteAtIndex(2)
	l.AddAtHead(6)
	l.AddAtTail(4)
	l.Get(4)
	l.AddAtHead(4)
	l.AddAtIndex(5, 0)
	l.AddAtHead(6)
	l.print()
}
