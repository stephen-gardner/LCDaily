// Problem: https://leetcode.com/problems/concatenation-of-consecutive-binary-numbers/
// Results: https://leetcode.com/submissions/detail/806499749/
package main

const resMod = 1e9 + 7

func concatenatedBinary(n int) int {
	res, bitLen, next := 0, 1, 2
	for i := 1; i <= n; i++ {
		if i == next {
			bitLen++
			next <<= 1
		}
		res = ((res << bitLen) | i) % resMod
	}
	return res
}
