// Problem: https://leetcode.com/problems/unique-number-of-occurrences/
// Results: https://leetcode.com/submissions/detail/852063995/
package main

func uniqueOccurrences(nums []int) bool {
	counts := map[int]int{}
	for _, n := range nums {
		counts[n]++
	}
	have := map[int]bool{}
	for _, n := range counts {
		if have[n] {
			return false
		}
		have[n] = true
	}
	return true
}
