// Problem: https://leetcode.com/problems/sort-colors/
// Results: https://leetcode.com/problems/sort-colors/submissions/1285494267
package main

func fill(nums []int, val, n int) {
	for i := 0; i < n; i++ {
		nums[i] = val
	}
}

// Time: O(n)
// Space: O(1)
func sortColors(nums []int) {
	var counts [3]int
	for _, n := range nums {
		counts[n]++
	}
	for val, n := range counts {
		fill(nums, val, n)
		nums = nums[n:]
	}
}
