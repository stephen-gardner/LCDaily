// Problem: https://leetcode.com/problems/xor-queries-of-a-subarray/
// Results: https://leetcode.com/problems/xor-queries-of-a-subarray/submissions/1388243389
package main

// Time: O(n + m)
// Space: O(n + m)
// n = len(arr), m = len(queries)
func xorQueries(arr []int, queries [][]int) []int {
	pre := make([]int, len(arr))
	pre[0] = arr[0]
	for i := 1; i < len(arr); i++ {
		pre[i] = pre[i-1] ^ arr[i]
	}
	res := make([]int, len(queries))
	for i, q := range queries {
		start, end := q[0], q[1]
		res[i] = pre[end]
		if start-1 >= 0 {
			res[i] ^= pre[start-1]
		}
	}
	return res
}
