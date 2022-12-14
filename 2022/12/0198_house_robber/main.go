// Problem: https://leetcode.com/problems/house-robber/
// Results: https://leetcode.com/problems/house-robber/submissions/859421783/
package main

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func rob(nums []int) int {
	// Hot  = Just robbed a house
	// Cold = Just skipped a house
	cold, hot := 0, 0
	for i := 0; i < len(nums); i++ {
		hot, cold = max(hot, cold+nums[i]), hot
	}
	return hot
}
