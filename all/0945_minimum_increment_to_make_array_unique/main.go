// Problem: https://leetcode.com/problems/minimum-increment-to-make-array-unique/
// Results: https://leetcode.com/problems/minimum-increment-to-make-array-unique/submissions/1288560377
package main

func insert(space map[int]int, n int) int {
	insertedAt := n
	if next, full := space[n]; full {
		insertedAt = insert(space, next+1)
	}
	space[n] = insertedAt
	return insertedAt
}

// Time: O(n)
// Space: O(n)
func minIncrementForUnique(nums []int) int {
	increments := 0
	space := map[int]int{}
	for _, n := range nums {
		increments += insert(space, n) - n
	}
	return increments
}
