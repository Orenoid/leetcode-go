package main

import "fmt"

type MyQueue struct {
	in  []int
	out []int
}

func Constructor() MyQueue {
	return MyQueue{
		in:  []int{},
		out: []int{},
	}
}

func (this *MyQueue) Push(x int) {
	this.in = append(this.in, x)
}

func (this *MyQueue) Pop() int {
	if this.Empty() {
		return 0
	}
	if len(this.out) != 0 {
		return this.popOut()
	}
	for len(this.in) > 0 {
		val := this.in[len(this.in)-1]
		this.in = this.in[:len(this.in)-1]
		this.out = append(this.out, val)
	}
	return this.popOut()
}

func (this *MyQueue) popOut() int {
	if len(this.out) != 0 {
		val := this.out[len(this.out)-1]
		this.out = this.out[:len(this.out)-1]
		return val
	}
	return 0
}

func (this *MyQueue) Peek() int {
	if this.Empty() {
		return 0
	}
	if len(this.out) > 0 {
		return this.out[len(this.out)-1]
	}
	for len(this.in) > 0 {
		val := this.in[len(this.in)-1]
		this.in = this.in[:len(this.in)-1]
		this.out = append(this.out, val)
	}
	return this.out[len(this.out)-1]
}

func (this *MyQueue) Empty() bool {
	return len(this.in) == 0 && len(this.out) == 0
}

/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */

func main() {
	q := Constructor()
	fmt.Println(q.Empty())
	q.Push(1)
	q.Push(2)
	fmt.Println(q.Pop())
	q.Push(3)
	fmt.Println(q.Pop())
	fmt.Println(q.Peek())
}
