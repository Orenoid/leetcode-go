package main

import (
	"fmt"
	"sort"
)

type myInterface struct {
	intervals [][]int
}

func (receiver *myInterface) Len() int {
	return len(receiver.intervals)
}

func (receiver *myInterface) Less(i, j int) bool {
	return receiver.intervals[i][0] < receiver.intervals[j][0]
}

func (receiver *myInterface) Swap(i, j int) {
	receiver.intervals[i], receiver.intervals[j] = receiver.intervals[j], receiver.intervals[i]
}

func merge(intervals [][]int) [][]int {
	sortI := &myInterface{intervals: intervals}
	sort.Sort(sortI)

	var res [][]int
	currInterval := intervals[0]

	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] <= currInterval[1] {
			if currInterval[1] < intervals[i][1] {
				currInterval[1] = intervals[i][1]
			}
		} else {
			res = append(res, []int{currInterval[0], currInterval[1]})
			currInterval = []int{intervals[i][0], intervals[i][1]}
		}
	}
	res = append(res, []int{currInterval[0], currInterval[1]})

	return res
}

func main() {
	fmt.Println(merge([][]int{{1, 3}, {3, 6}, {2, 3}, {5, 7}, {8, 12}, {9, 10}}))
}
