// Problem: https://leetcode.com/problems/maximum-score-from-performing-multiplication-operations/
// Results: https://leetcode.com/submissions/detail/800971256/
package main

func calcScore(nums, multi []int, cache [][]int, leftIdx, rightIdx, mIdx int) int {
	if mIdx == len(multi) {
		return 0
	}
	if score := cache[leftIdx][mIdx]; score != 0 {
		return score
	}
	left := (nums[leftIdx] * multi[mIdx]) + calcScore(nums, multi, cache, leftIdx+1, rightIdx, mIdx+1)
	right := (nums[rightIdx] * multi[mIdx]) + calcScore(nums, multi, cache, leftIdx, rightIdx-1, mIdx+1)
	if left > right {
		cache[leftIdx][mIdx] = left
		return left
	} else {
		cache[leftIdx][mIdx] = right
		return right
	}
}

func maximumScore(nums, multi []int) int {
	cache := make([][]int, len(multi)+1)
	for i := range cache {
		cache[i] = make([]int, len(multi)+1)
	}
	return calcScore(nums, multi, cache, 0, len(nums)-1, 0)
}
