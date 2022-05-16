package main

import (
	"container/heap"
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

type myHeap struct {
	heap []*ListNode
}

func (receiver *myHeap) Len() int {
	return len(receiver.heap)
}

func (receiver *myHeap) Less(i, j int) bool {
	return receiver.heap[i].Val < receiver.heap[j].Val
}

func (receiver *myHeap) Swap(i, j int) {
	receiver.heap[i], receiver.heap[j] = receiver.heap[j], receiver.heap[i]
}

func (receiver *myHeap) Push(x interface{}) {
	receiver.heap = append(receiver.heap, x.(*ListNode))
}

func (receiver *myHeap) Pop() interface{} {
	res := receiver.heap[len(receiver.heap)-1]
	receiver.heap = receiver.heap[:len(receiver.heap)-1]
	return res
}

func mergeKLists(lists []*ListNode) *ListNode {
	h := &myHeap{}
	heap.Init(h)
	dummy := &ListNode{}
	curr := dummy
	for _, node := range lists {
		if node != nil {
			heap.Push(h, node)
		}
	}
	for h.Len() > 0 {
		minNode := heap.Pop(h).(*ListNode)
		next := minNode.Next
		if next != nil {
			heap.Push(h, next)
		}
		minNode.Next = nil
		curr.Next = minNode
		curr = curr.Next
	}
	return dummy.Next
}

func main() {
	fmt.Println(mergeKLists(nil))
}
