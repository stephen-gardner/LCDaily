// Problem: https://leetcode.com/problems/longest-subsequence-with-limited-sum/
// Results: https://leetcode.com/problems/longest-subsequence-with-limited-sum/submissions/865086249/
package main

import (
	"sort"
)

func answerQueries(nums []int, queries []int) []int {
	sort.Ints(nums)
	answer := make([]int, len(queries))
	for i, q := range queries {
		sum := 0
		j := 0
		for sum < q && j < len(nums) {
			if sum += nums[j]; sum <= q {
				j++
			}
		}
		answer[i] = j
	}
	return answer
}
