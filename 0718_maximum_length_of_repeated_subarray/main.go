// Problem: https://leetcode.com/problems/maximum-length-of-repeated-subarray/
// Results: https://leetcode.com/submissions/detail/804270056/
package main

// n = len(nums1); m = len(nums2)
// Time: O(n*m)
// Space: O(m)
func findLength(nums1 []int, nums2 []int) int {
	max := 0
	if len(nums1) < len(nums2) {
		nums1, nums2 = nums2, nums1
	}
	buff := make([]int, len(nums2)+1)
	for idxN1, n1 := range nums1 {
		// Fill buffer backwards to avoid overwriting data before we need it
		for idxN2 := len(nums2) - 1; idxN2 >= 0; idxN2-- {
			i := idxN2 + 1
			if n1 != nums2[idxN2] {
				buff[i] = 0
				continue
			}
			// If previous characters also match, use previous count + 1
			if idxN1 > 0 && idxN2 > 0 && nums1[idxN1-1] == nums2[idxN2-1] {
				buff[i] = buff[i-1] + 1
			} else {
				buff[i] = 1
			}
			if buff[i] > max {
				max = buff[i]
			}
		}
	}
	return max
}
