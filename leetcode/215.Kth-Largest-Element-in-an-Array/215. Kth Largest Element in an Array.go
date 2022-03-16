package main

import (
	"container/heap"
	"fmt"
)

type MyHeap struct {
	collection []int
}

func (h *MyHeap) Len() int {
	return len(h.collection)
}

func (h *MyHeap) Less(i, j int) bool {
	return h.collection[i] < h.collection[j]
}

func (h *MyHeap) Swap(i, j int) {
	h.collection[i], h.collection[j] = h.collection[j], h.collection[i]
}

func (h *MyHeap) Push(x interface{}) {
	h.collection = append(h.collection, x.(int))
}

func (h *MyHeap) Pop() interface{} {
	v := h.collection[len(h.collection)-1]
	h.collection = h.collection[:len(h.collection)-1]
	return v
}

func findKthLargest(nums []int, k int) int {
	myHeap := &MyHeap{[]int{}}
	heap.Init(myHeap)
	for i := range nums {
		if myHeap.Len() == k {
			if nums[i] < myHeap.collection[0] {
				continue
			}
			heap.Pop(myHeap)
		}
		heap.Push(myHeap, nums[i])
	}

	return heap.Pop(myHeap).(int)
}

func main() {
	fmt.Println(findKthLargest([]int{
		3, 2, 3, 1, 2, 4, 5, 5, 6}, 4))
}
