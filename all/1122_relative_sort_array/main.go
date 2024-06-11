// Problem: https://leetcode.com/problems/relative-sort-array/
// Results: https://leetcode.com/problems/relative-sort-array/submissions/1284510026
package main

import "sort"

// Time: O(n*log(n))
// Space: O(n); n = len(arr2)
func relativeSortArray(arr1 []int, arr2 []int) []int {
	have := map[int]int{}
	for i, n := range arr2 {
		have[n] = i
	}
	sort.Slice(arr1, func(i, j int) bool {
		_, aHave := have[arr1[i]]
		_, bHave := have[arr1[j]]
		if aHave && bHave {
			return have[arr1[i]] < have[arr1[j]]
		} else if aHave {
			return true
		} else if bHave {
			return false
		}
		return arr1[i] < arr1[j]
	})
	return arr1
}
