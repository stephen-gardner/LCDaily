// Problem: https://leetcode.com/problems/design-circular-queue/
// Results: https://leetcode.com/submissions/detail/807890913/
package main

type MyCircularQueue struct {
	arr   []int
	start int
	n     int
}

func Constructor(k int) MyCircularQueue {
	return MyCircularQueue{
		arr: make([]int, k),
	}
}

func (this *MyCircularQueue) EnQueue(value int) bool {
	if this.IsFull() {
		return false
	}
	this.arr[(this.start+this.n)%len(this.arr)] = value
	this.n++
	return true
}

func (this *MyCircularQueue) DeQueue() bool {
	if this.IsEmpty() {
		return false
	}
	this.start = (this.start + 1) % len(this.arr)
	this.n--
	return true
}

func (this *MyCircularQueue) Front() int {
	if this.IsEmpty() {
		return -1
	}
	return this.arr[this.start]
}

func (this *MyCircularQueue) Rear() int {
	if this.IsEmpty() {
		return -1
	}
	return this.arr[(this.start+(this.n-1))%len(this.arr)]
}

func (this *MyCircularQueue) IsEmpty() bool {
	return this.n == 0
}

func (this *MyCircularQueue) IsFull() bool {
	return this.n == len(this.arr)
}
