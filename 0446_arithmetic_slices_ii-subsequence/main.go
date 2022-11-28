// Problem: https://leetcode.com/problems/arithmetic-slices-ii-subsequence/
// Results: https://leetcode.com/submissions/detail/850895688/
package main

func numberOfArithmeticSlices(nums []int) int {
	res := 0
	dp := make([]map[int]int, len(nums))
	for i := range nums {
		dp[i] = make(map[int]int)
		for j := 0; j < i; j++ {
			delta := nums[i] - nums[j]
			res += dp[j][delta]
			dp[i][delta] += dp[j][delta] + 1
		}
	}
	return res
}
