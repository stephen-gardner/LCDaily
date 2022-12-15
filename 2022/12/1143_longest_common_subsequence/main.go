// Problem: https://leetcode.com/problems/longest-common-subsequence/
// Results: https://leetcode.com/problems/longest-common-subsequence/submissions/859996865/
package main

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func longestCommonSubsequence(text1 string, text2 string) int {
	m, n := len(text1), len(text2)
	curr := make([]int, n+1)
	for i := 0; i < m; i++ {
		prev := curr
		curr = make([]int, n+1)
		for j := 0; j < n; j++ {
			if text1[i] == text2[j] {
				curr[j+1] = 1 + prev[j]
			} else {
				curr[j+1] = max(curr[j], prev[j+1])
			}
		}
	}
	return curr[n]
}
