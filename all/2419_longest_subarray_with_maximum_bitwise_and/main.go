// Problem: https://leetcode.com/problems/longest-subarray-with-maximum-bitwise-and/
// Results: https://leetcode.com/problems/longest-subarray-with-maximum-bitwise-and/submissions/1389499681
package main

// Time: O(n)
// Space: O(1)
func longestSubarray(nums []int) int {
	highVal, highLen := nums[0], 0
	for i := 0; i < len(nums); i++ {
		if nums[i] < highVal {
			continue
		}
		if nums[i] > highVal {
			highVal, highLen = nums[i], 0
		}
		j := i + 1
		for j < len(nums) && nums[j] == nums[j-1] {
			j++
		}
		if j-i > highLen {
			highLen = j - i
		}
		i = j - 1
	}
	return highLen
}
