// Problem: https://leetcode.com/problems/remove-stones-to-minimize-the-total/
// Results: https://leetcode.com/problems/remove-stones-to-minimize-the-total/submissions/866631279/
package main

import (
	"container/heap"
)

type IntHeap []int

func (h IntHeap) Len() int {
	return len(h)
}

func (h IntHeap) Less(i, j int) bool {
	return h[i] > h[j]
}

func (h IntHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	x := old[len(old)-1]
	*h = old[0 : len(old)-1]
	return x
}

func minStoneSum(piles IntHeap, k int) int {
	heap.Init(&piles)
	for k > 0 {
		piles[0] -= piles[0] / 2
		heap.Fix(&piles, 0)
		k--
	}
	sum := 0
	for _, n := range piles {
		sum += n
	}
	return sum
}
