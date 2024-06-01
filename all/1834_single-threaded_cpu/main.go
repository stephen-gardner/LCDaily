// Problem: https://leetcode.com/problems/single-threaded-cpu/
// Results: https://leetcode.com/problems/single-threaded-cpu/submissions/867287797/
package main

import (
	"container/heap"
	"sort"
)

type Process struct {
	idx           int
	executionTime int
	readyAt       int
}

type Queue []Process

func (q Queue) Len() int {
	return len(q)
}

func (q Queue) Less(i, j int) bool {
	p1, p2 := q[i], q[j]
	if p1.executionTime == p2.executionTime {
		return p1.idx < p2.idx
	}
	return p1.executionTime < p2.executionTime
}

func (q Queue) Swap(i, j int) {
	q[i], q[j] = q[j], q[i]
}

func (q *Queue) Push(x interface{}) {
	*q = append(*q, x.(Process))
}

func (q *Queue) Pop() interface{} {
	old := *q
	n := len(old)
	x := old[n-1]
	*q = old[:n-1]
	return x
}

func getOrder(tasks [][]int) []int {
	order := make([]int, 0, len(tasks))
	procs := make([]Process, len(tasks))
	for i := range procs {
		procs[i] = Process{
			idx:           i,
			executionTime: tasks[i][1],
			readyAt:       tasks[i][0],
		}
	}
	sort.Slice(procs, func(i, j int) bool {
		return procs[i].readyAt < procs[j].readyAt
	})
	i := 0
	currTime := 0
	queue := &Queue{}
	for len(order) < len(tasks) {
		for i < len(procs) && (queue.Len() == 0 || procs[i].readyAt <= currTime) {
			p := procs[i]
			if queue.Len() == 0 && p.readyAt > currTime {
				currTime = p.readyAt
			}
			heap.Push(queue, p)
			i++
		}
		p := heap.Pop(queue).(Process)
		order = append(order, p.idx)
		currTime += p.executionTime
	}
	return order
}
