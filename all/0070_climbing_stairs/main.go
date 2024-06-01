// Problem: https://leetcode.com/problems/climbing-stairs/
// Results: https://leetcode.com/problems/climbing-stairs/submissions/858560735/
package main

// It's fibonacci sequence
func climbStairs(n int) int {
	if n == 1 {
		return 1
	}
	a, b := 1, 1
	for n--; n > 0; n-- {
		a, b = b, a+b
	}
	return b
}
