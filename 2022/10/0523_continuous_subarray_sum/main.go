// Problem: https://leetcode.com/problems/continuous-subarray-sum/
// Results: https://leetcode.com/submissions/detail/831038443/
package main

func checkSubarraySum(nums []int, k int) bool {
	rem := map[int]int{
		0: -1, // If first value in nums is a multiple, can't return true
	}
	r := 0
	for i, n := range nums {
		r = (r + n) % k
		if idx, exists := rem[r]; exists {
			if i-idx >= 2 {
				return true
			}
		} else {
			rem[r] = i
		}
	}
	return false
}
