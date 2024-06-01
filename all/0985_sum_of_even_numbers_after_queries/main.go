// Problem: https://leetcode.com/problems/sum-of-even-numbers-after-queries/
// Results: https://leetcode.com/submissions/detail/805074854/
package main

func sumEvenAfterQueries(nums []int, queries [][]int) []int {
	sum := 0
	for _, n := range nums {
		if n&1 == 0 {
			sum += n
		}
	}
	res := make([]int, len(queries))
	for i, query := range queries {
		idx := query[1]
		if nums[idx]&1 == 0 {
			sum -= nums[idx]
		}
		nums[idx] += query[0]
		if nums[idx]&1 == 0 {
			sum += nums[idx]
		}
		res[i] = sum
	}
	return res
}
