// Problem: https://leetcode.com/problems/ipo/
// Results: https://leetcode.com/problems/ipo/submissions/1289425515
package main

import (
	"container/heap"
)

type Project struct {
	profit  int
	capital int
}

type ProjectHeap struct {
	data []Project
	cmp  func(int, int) bool
}

func (ph ProjectHeap) Len() int           { return len(ph.data) }
func (ph ProjectHeap) Less(i, j int) bool { return ph.cmp(i, j) }
func (ph ProjectHeap) Swap(i, j int)      { ph.data[i], ph.data[j] = ph.data[j], ph.data[i] }
func (ph *ProjectHeap) Peek() Project     { return ph.data[0] }
func (ph *ProjectHeap) Push(x any)        { ph.data = append(ph.data, x.(Project)) }
func (ph *ProjectHeap) Pop() any {
	old := ph.data
	n := len(old)
	x := old[n-1]
	ph.data = old[:n-1]
	return x
}

func NewProjectHeap(capital bool) *ProjectHeap {
	ph := &ProjectHeap{data: make([]Project, 0)}
	if capital {
		ph.cmp = func(i, j int) bool {
			return ph.data[i].capital < ph.data[j].capital
		}
	} else {
		ph.cmp = func(i, j int) bool {
			return ph.data[i].profit > ph.data[j].profit
		}
	}
	return ph
}

// Time: O(n*log(n))
// Space: O(n)
func findMaximizedCapital(k int, w int, profits []int, capital []int) int {
	// Max heap of profit
	// Min heap of capital
	// For every w increase, pop off min heap the projects to add to max heap
	capitalHeap := NewProjectHeap(true)
	profitHeap := NewProjectHeap(false)
	for i := range capital {
		p := Project{
			profit:  profits[i],
			capital: capital[i],
		}
		if w >= capital[i] {
			heap.Push(profitHeap, p)
		} else {
			heap.Push(capitalHeap, p)
		}
	}
	for i := 0; i < k; i++ {
		for capitalHeap.Len() > 0 && w >= capitalHeap.Peek().capital {
			heap.Push(profitHeap, heap.Pop(capitalHeap))
		}
		if profitHeap.Len() <= 0 {
			break
		}
		w += heap.Pop(profitHeap).(Project).profit
	}
	return w
}
