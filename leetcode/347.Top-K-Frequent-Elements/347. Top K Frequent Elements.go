package main

import (
	"container/heap"
	"fmt"
)

func topKFrequent(nums []int, k int) []int {
	counter := make(map[int]int, len(nums))
	for _, num := range nums {
		counter[num]++
	}
	h := &HeapI{}
	heap.Init(h)
	for num, count := range counter {
		heap.Push(h, [2]int{num, count})
		if h.Len() > k {
			heap.Pop(h)
		}
	}
	var result []int
	for i := 0; i < k; i++ {
		val := heap.Pop(h).([2]int)
		result = append(result, val[0])
	}
	return result
}

type HeapI [][2]int

func (h HeapI) Len() int {
	return len(h)
}

func (h HeapI) Less(i, j int) bool {
	return h[i][1] < h[j][1]
}

func (h HeapI) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *HeapI) Push(x interface{}) {
	*h = append(*h, x.([2]int))
}

func (h *HeapI) Pop() interface{} {
	val := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return val
}

func main() {
	fmt.Println(topKFrequent([]int{1, 1, 1, 2, 3, 3}, 2))
}
