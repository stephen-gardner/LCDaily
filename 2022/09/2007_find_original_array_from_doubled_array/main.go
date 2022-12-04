// Problem: https://leetcode.com/problems/find-original-array-from-doubled-array/
// Results: https://leetcode.com/submissions/detail/800166307/
package main

import "sort"

func findOriginalArray(changed []int) []int {
	if len(changed)&1 != 0 {
		return changed[:0]
	}
	sort.Ints(changed)
	pool := map[int]int{}
	for _, n := range changed {
		pool[n]++
	}
	res := make([]int, 0, len(changed)/2)
	for _, n := range changed {
		if amount := pool[n]; amount == 0 {
			continue
		}
		pool[n]--
		doubleN := n * 2
		if amount := pool[doubleN]; amount > 0 {
			pool[doubleN]--
			res = append(res, n)
		} else {
			return res[:0]
		}
	}
	return res
}
