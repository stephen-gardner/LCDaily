// Problem: https://leetcode.com/problems/ugly-number/
// Results: https://leetcode.com/submissions/detail/845513344/
package main

func isUgly(n int) bool {
	if n <= 0 {
		return false
	}
	divisor := 5
	for divisor > 1 {
		if n%divisor == 0 {
			n /= divisor
		} else {
			divisor--
		}
	}
	return n == 1
}
