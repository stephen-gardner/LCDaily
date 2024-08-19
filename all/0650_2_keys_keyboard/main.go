// Problem: https://leetcode.com/problems/2-keys-keyboard/
// Results: https://leetcode.com/problems/2-keys-keyboard/submissions/1361616764
package main

// Time: O(2^n)
// Space: O(n)
func sim(ops, delta, count, target int) int {
	if count == target {
		return ops
	}
	if count+count <= target {
		if res := sim(ops+2, count, count+count, target); res > 0 {
			return res
		}
	}
	if count+delta <= target {
		return sim(ops+1, delta, count+delta, target)
	}
	return 0
}

func minSteps(n int) int {
	return sim(2, 1, 2, n)
}
