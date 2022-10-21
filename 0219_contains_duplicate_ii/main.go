// Problem: https://leetcode.com/problems/contains-duplicate-ii/
// Results: https://leetcode.com/submissions/detail/827064391/
package main

func containsNearbyDuplicate(nums []int, k int) bool {
	visited := map[int]int{}
	for i, n := range nums {
		if idx, seen := visited[n]; seen && i-idx <= k {
			return true
		}
		visited[n] = i
	}
	return false
}
