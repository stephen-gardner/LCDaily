// Problem: https://leetcode.com/problems/set-mismatch/
// Results: https://leetcode.com/submissions/detail/828249395/
package main

func findErrorNums(nums []int) []int {
	seen := make([]bool, len(nums))
	dupe := -1
	for _, n := range nums {
		if seen[n-1] {
			dupe = n
		}
		seen[n-1] = true
	}
	for i := range seen {
		if !seen[i] {
			return []int{dupe, i + 1}
		}
	}
	return nil
}
