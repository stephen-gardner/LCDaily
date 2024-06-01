// Problem: https://leetcode.com/problems/top-k-frequent-words/submissions/
// Results: https://leetcode.com/submissions/detail/825660791/
package main

import (
	"container/heap"
)

type Word struct {
	value string
	count int
}

type MaxHeap []Word

func (mh MaxHeap) Len() int {
	return len(mh)
}

func (mh MaxHeap) Less(i, j int) bool {
	if mh[i].count == mh[j].count {
		return mh[i].value < mh[j].value
	}
	return mh[i].count > mh[j].count
}

func (mh MaxHeap) Swap(i, j int) {
	mh[i], mh[j] = mh[j], mh[i]
}

func (mh *MaxHeap) Push(x interface{}) {
	*mh = append(*mh, x.(Word))
}

func (mh *MaxHeap) Pop() interface{} {
	arr := *mh
	item := arr[len(arr)-1]
	*mh = arr[:len(arr)-1]
	return item
}

func topKFrequent(words []string, k int) []string {
	wordCount := map[string]int{}
	for _, w := range words {
		wordCount[w]++
	}
	mh := make(MaxHeap, 0, len(wordCount))
	for w, count := range wordCount {
		heap.Push(&mh, Word{
			value: w,
			count: count,
		})
	}
	res := make([]string, k)
	for i := 0; i < k; i++ {
		res[i] = heap.Pop(&mh).(Word).value
	}
	return res
}
