// Problem: https://leetcode.com/problems/increasing-triplet-subsequence/
// Results: https://leetcode.com/submissions/detail/819837387/
package main

func increasingTriplet(nums []int) bool {
	var first, second *int
	for i := range nums {
		if first == nil || nums[i] <= *first {
			first = &nums[i]
		} else if second == nil || nums[i] <= *second {
			second = &nums[i]
		} else {
			return true
		}
	}
	return false
}
