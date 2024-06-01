// Problem: https://leetcode.com/problems/minimum-average-difference/
// Results: https://leetcode.com/submissions/detail/854231170/
package main

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func minimumAverageDifference(nums []int) int {
	var lSum, rSum, diff int
	minDiff, minIdx := int(1e5)+1, -1
	for _, n := range nums {
		rSum += n
	}
	for i, lSize, rSize := 0, 0, len(nums); i < len(nums); i++ {
		lSum, lSize = lSum+nums[i], lSize+1
		rSum, rSize = rSum-nums[i], rSize-1
		diff = lSum / lSize
		if rSize > 0 {
			diff = abs(diff - (rSum / rSize))
		}
		if diff < minDiff {
			minDiff, minIdx = diff, i
		}
	}
	return minIdx
}
