// Problem: https://leetcode.com/problems/maximum-performance-of-a-team/
// Results: https://leetcode.com/submissions/detail/797539487/
package main

import (
	"container/heap"
	"sort"
)

type Team []int

func (t Team) Len() int {
	return len(t)
}

func (t Team) Less(i, j int) bool {
	return t[i] < t[j]
}

func (t Team) Swap(i, j int) {
	t[i], t[j] = t[j], t[i]
}

func (t *Team) Push(x interface{}) {
	*t = append(*t, x.(int))
}

func (t *Team) Pop() interface{} {
	x := (*t)[len(*t)-1]
	*t = (*t)[:len(*t)-1]
	return x
}

func maxPerformance(n int, speed []int, efficiency []int, k int) int {
	engineers := make([]int, n)
	for i := 0; i < n; i++ {
		engineers[i] = i
	}
	sort.Slice(engineers, func(i, j int) bool {
		return efficiency[engineers[i]] > efficiency[engineers[j]]
	})
	team := &Team{}
	bestScore := int64(0)
	currSum := int64(0)
	for _, i := range engineers {
		if team.Len() == k {
			currSum -= int64(heap.Pop(team).(int))
		}
		heap.Push(team, speed[i])
		currSum += int64(speed[i])
		if score := currSum * int64(efficiency[i]); score > bestScore {
			bestScore = score
		}
	}
	return int(bestScore % 1000000007)
}
