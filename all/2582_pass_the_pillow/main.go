// Problem: https://leetcode.com/problems/pass-the-pillow/
// Results: https://leetcode.com/problems/pass-the-pillow/submissions/1312164811
package main

// Time: O(1)
// Space: O(1)
func passThePillow(n int, time int) int {
	loops := time / (n - 1)
	rem := time % (n - 1)
	if loops&1 == 1 {
		// To the left
		return n - rem
	}
	return 1 + rem
}
