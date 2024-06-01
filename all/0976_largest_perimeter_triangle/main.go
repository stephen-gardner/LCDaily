// Problem: https://leetcode.com/problems/largest-perimeter-triangle/
// Results: https://leetcode.com/submissions/detail/820583204/
package main

import "sort"

func largestPerimeter(nums []int) int {
	sort.Sort(sort.Reverse(sort.IntSlice(nums)))
	for i := 0; i < len(nums)-2; i++ {
		if nums[i] < nums[i+1]+nums[i+2] {
			return nums[i] + nums[i+1] + nums[i+2]
		}
	}
	return 0
}
