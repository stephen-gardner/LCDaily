// Problem: https://leetcode.com/problems/find-median-from-data-stream/
// Results: https://leetcode.com/submissions/detail/841821629/
package main

import "container/heap"

type Heap struct {
	arr []int
	min bool
}

func (heap Heap) Len() int {
	return len(heap.arr)
}

func (heap Heap) Less(i, j int) bool {
	if heap.min {
		return heap.arr[i] < heap.arr[j]
	} else {
		return heap.arr[i] > heap.arr[j]
	}
}

func (heap Heap) Swap(i, j int) {
	heap.arr[i], heap.arr[j] = heap.arr[j], heap.arr[i]
}

func (heap *Heap) Push(x interface{}) {
	heap.arr = append(heap.arr, x.(int))
}

func (heap *Heap) Pop() interface{} {
	x := heap.arr[len(heap.arr)-1]
	heap.arr = heap.arr[:len(heap.arr)-1]
	return x
}

type MedianFinder struct {
	lower  *Heap
	upper  *Heap
	median float64
}

func Constructor() MedianFinder {
	return MedianFinder{
		lower: &Heap{min: false},
		upper: &Heap{min: true},
	}
}

func (this *MedianFinder) AddNum(n int) {
	if this.lower.Len() == 0 && this.upper.Len() == 0 {
		heap.Push(this.lower, n)
		this.median = float64(n)
		return
	}
	if n > int(this.median) {
		heap.Push(this.upper, n)
	} else {
		heap.Push(this.lower, n)
	}
	a, b := this.upper, this.lower
	if b.Len() > a.Len() {
		a, b = b, a
	}
	if a.Len()-b.Len() > 1 {
		heap.Push(b, heap.Pop(a))
	}
	if a.Len() == b.Len() {
		this.median = (float64(a.arr[0]) + float64(b.arr[0])) / 2.0
	} else if a.Len() > b.Len() {
		this.median = float64(a.arr[0])
	} else {
		this.median = float64(b.arr[0])
	}
}

func (this *MedianFinder) FindMedian() float64 {
	return this.median
}
